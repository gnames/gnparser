// Package preprocess performs preparsing filtering and modification of a
// scientific-name.
package preprocess

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"unicode"

	"github.com/gnames/gnparser/ent/internal/preparser"
)

var VirusException = map[string]string{
	"Aspilota":      "vector",
	"Bembidion":     "satellites",
	"Bolivina":      "prion",
	"Ceylonesmus":   "vector",
	"Cryptops":      "vector",
	"Culex":         "vector",
	"Dasyproctus":   "cevirus",
	"Desmoxytes":    "vector",
	"Dicathais":     "vector",
	"Erateina":      "satellites",
	"Euragallia":    "prion",
	"Exochus":       "virus",
	"Hilara":        "vector",
	"Ithomeis":      "satellites",
	"Microgoneplax": "prion",
	"Neoaemula":     "vector",
	"Nephodia":      "satellites",
	"Ophion":        "virus",
	"Psenulus":      "trevirus",
	"Tidabius":      "vector",
}

var AmbiguousException = map[string][]string{
	"Agnetina":       {"den"},
	"Antaplaga":      {"dela"},
	"Baeolidia":      {"dela"},
	"Bolitoglossa":   {"la"},
	"Campylosphaera": {"dela"},
	"Desmoxytes":     {"des"},
	"Dicentria":      {"dela"},
	"Eulaira":        {"dela"},
	"Gnathopleustes": {"den"},
	"Gobiosoma":      {"spec"},
	"Helophorus":     {"ser"},
	"Lampona":        {"spec"},
	"Leptonetela":    {"la"},
	"Malamatidia":    {"zu"},
	"Meteorus":       {"dos"},
	"Nocaracris":     {"van"},
	"Paralvinella":   {"dela"},
	"Ruteloryctes":   {"bis"},
	"Scoparia":       {"dela"},
	"Selenops":       {"ab"},
	"Semiothisa":     {"da"},
	"Serina":         {"ser", "subser"},
	"Stenoecia":      {"dos"},
	"Sympycnus":      {"du"},
	"Tortolena":      {"dela"},
	"Zodarion":       {"van"},
}

var NoParseException = map[string]string{
	"Navicula": "bacterium",
}

var cultivarRankRe = regexp.MustCompile(
	`\s+(cultivar\.?[\W_]|cv\.?[\W_]|['"‘’“”]).*$`,
)

var ofWordRe = regexp.MustCompile(
	`\s+(of[\W_]).*$`,
)

var dagger = []byte("†")

// Preprocessor structure keeps state of the preprocessor results.
type Preprocessor struct {
	Virus       bool
	Underscore  bool
	NoParse     bool
	DaggerChar  bool
	Approximate bool
	Annotation  bool
	Body        []byte
	Tail        []byte
	Ambiguous   ambiguous
}

type ambiguous struct {
	Orig  string
	Subst string
}

// Preprocess runs a series of regular expressions over the input to determine
// features of the input before parsing.
func Preprocess(ppr *preparser.PreParser, bs []byte) *Preprocessor {
	pr := &Preprocessor{}

	// check for empty string
	if len(bs) == 0 || strings.TrimSpace(string(bs)) == "" {
		pr.NoParse = true
		return pr
	}
	i := len(bs)
	words := strings.Fields(string(bs))

	// check for viruses, plasmids, RNA, DNA etc.
	if !isException(words, VirusException) {
		pr.Virus = IsVirus(bs[0:i])
	}
	if pr.Virus {
		pr.NoParse = true
		return pr
	}

	// check for unparseable names
	pr.NoParse = NoParse(bs[0:i])
	if isException(words, NoParseException) {
		pr.NoParse = false
	}
	if pr.NoParse {
		return pr
	}

	pr.DaggerChar = hasDagger(bs[0:i])

	if len(words) > 1 {
		pr.ambiguous(words[0], bs)
	}

	j := procAnnot(ppr, bs[0:i])
	if j < i {
		pr.Annotation = true
		i = j
	}

	// ignoring error, as it should never happen
	changed, _ := UnderscoreToSpace(bs[0:i])
	if changed {
		pr.Underscore = true
	}

	pr.Body = bs[0:i]
	pr.Tail = bs[i:]
	return pr
}

func hasDagger(bs []byte) bool {
	idx := bytes.Index(bs, dagger)
	if idx == -1 {
		return false
	}

	sp := byte(' ')
	bs[idx], bs[idx+1], bs[idx+2] = sp, sp, sp
	return true
}

func isException(words []string, names map[string]string) bool {
	if len(words) < 2 {
		return false
	}
	if epithet, ok := names[words[0]]; ok {
		for _, w := range words[1:] {
			if w == epithet {
				return true
			}
		}
	}
	return false
}

func (p *Preprocessor) ambiguous(firstWord string, bs []byte) {
	if epithets, ok := AmbiguousException[firstWord]; ok {
		var sub byte = 'k'
		for _, epithet := range epithets {
			idx := bytes.Index(bs, []byte(" "+epithet))
			if idx == -1 {
				continue
			}
			p.Ambiguous.Orig = epithet
			p.Ambiguous.Subst = string(sub) + epithet[1:]
			bs[idx+1] = sub
		}
	}
}

// procAnnot returns index where unparsed part starts. In case if
// the full string can be parsed, returns returns the index of the end of the
// input.
func procAnnot(ppr *preparser.PreParser, bs []byte) int {
	i := len(bs)
	if idx := ppr.TailIndex(string(bs)); idx >= 0 {
		i = idx
	}

	// If ` of ` is in the string, before the start of the already-calculated
	// unparsed part, but there is no cultivar rank marker before it, consider it
	// unparseable. `Anthurium 'Ace of Spades'` should parse fully;
	// `Anthurium Trustees of the British Museum` should not.
	cultivarRankLoc := cultivarRankRe.FindIndex(bs[0:i])
	ofLoc := ofWordRe.FindIndex(bs[0:i])
	if len(ofLoc) > 0 && ofLoc[0] < i &&
		(len(cultivarRankLoc) == 0 || cultivarRankLoc[0] > ofLoc[0]) {
		i = ofLoc[0]
	}

	return i
}

// UnderscoreToSpace takes a slice of bytes. If it finds that the string
// contains underscores, but not spaces, it substitutes underscores to spaces
// in the slice. In case if any spaces are present, the slice is returned
// unmodified.
func UnderscoreToSpace(bs []byte) (bool, error) {
	reader := bytes.NewReader(bs)
	var hasUnderscore bool
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return false, err
		}
		if unicode.IsSpace(r) {
			return false, nil
		}
		if r == '_' {
			hasUnderscore = true
		}
	}
	if !hasUnderscore {
		return false, nil
	}

	for i, v := range bs {
		if v == '_' {
			bs[i] = ' '
		}
	}
	return true, nil
}
