// Package preprocess performs preparsing filtering and modification of a
// scientific-name.
package preprocess

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"unicode"
)

var hybridCharRe1 = regexp.MustCompile(`(^)[Xx](\p{Lu})`)
var hybridCharRe2 = regexp.MustCompile(`(\s|^)[Xx](\s|$)`)

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

var AnnotationException = map[string]string{
	"Acrostichum":      "nudum",
	"Adiantum":         "nudum",
	"Africanthion":     "nudum",
	"Agathidium":       "nudum",
	"Aphaniosoma":      "nudum",
	"Aspidium":         "nudum",
	"Athyrium":         "nudum",
	"Blechnum":         "nudum",
	"Bottaria":         "nudum",
	"Gnathopleustes":   "den",
	"Lycopodium":       "nudum",
	"Nephrodium":       "nudum",
	"Paralvinella":     "dela",
	"Polypodium":       "nudum",
	"Polystichum":      "nudum",
	"Psilotum":         "nudum",
	"Ruteloryctes":     "bis",
	"Selenops":         "ab",
	"Tortolena":        "dela",
	"Trachyphloeosoma": "nudum",
	"Zodarion":         "van",
}

var NoParseException = map[string]string{
	"Navicula": "bacterium",
}

var notesRe = regexp.MustCompile(
	`(?i)\s+(environmental|samples|species\s+group|species\s+complex|clade|group|author|non|nec|vide|fide)\b.*$`,
)
var taxonConceptsRe1 = regexp.MustCompile(
	`(?i)\s+(sero(var|type)|sensu|auct|sec|near|str)\.?\b.*$`,
)
var taxonConceptsRe2 = regexp.MustCompile(
	`(,\s*|\s+)(\(?s\.\s?s\.|\(?s\.\s?l\.|\(?s\.\s?str\.|\(?s\.\s?lat\.).*$`,
)
var taxonConceptsRe3 = regexp.MustCompile(
	`(?i)(,\s*|\s+)(pro parte|p\.\s?p\.)\s*$`,
)
var nomenConceptsRe = regexp.MustCompile(
	`(?i)(,\s*|\s+)(\(?(nomen|nom\.|comb\.)(\s.*)?)$`,
)
var lastWordJunkRe = regexp.MustCompile(
	`(?i)(,\s*|\s+)` +
		`(var\.?|von|van|ined\.?` +
		`|sensu|new|non|nec|nudum|ssp\.?` +
		`|subsp|subgen|hybrid)\??\s*$`,
)

var stopWordsRe = regexp.MustCompile(
	`\s+(\(?ht\.?\W|\(?hort\.?\W|spec\.|nov\s+spec).*$`,
)

var cultivarRankRe = regexp.MustCompile(
	`\s+(cultivar\.?[\W_]|cv\.?[\W_]|['"‘’“”]).*$`,
)

var ofWordRe = regexp.MustCompile(
	`\s+(of[\W_]).*$`,
)

// Preprocessor structure keeps state of the preprocessor results.
type Preprocessor struct {
	Virus       bool
	Underscore  bool
	NoParse     bool
	Approximate bool
	Annotation  bool
	Body        []byte
	Tail        []byte
}

// Preprocess runs a series of regular expressions over the input to determine
// features of the input before parsing.
func Preprocess(bs []byte) *Preprocessor {
	pr := &Preprocessor{}
	if len(bs) == 0 {
		pr.NoParse = true
		return pr
	}
	i := len(bs)
	name := string(bs)
	if !IsException(name, VirusException) {
		pr.Virus = IsVirus(bs[0:i])
	}
	if pr.Virus {
		pr.NoParse = true
		return pr
	}
	pr.NoParse = NoParse(bs[0:i])
	if IsException(name, NoParseException) {
		pr.NoParse = false
	}
	if pr.NoParse {
		return pr
	}
	j := Annotation(bs[0:i])
	if j < i {
		pr.Annotation = true
		i = j
	}

	// ignoring error, as it should never happen
	changed, _ := UnderscoreToSpace(bs[0:i])
	if changed {
		pr.Underscore = true
	}

	pr.Body = NormalizeHybridChar(bs[0:i])
	pr.Tail = bs[i:]
	return pr
}

func IsException(name string, names map[string]string) bool {
	words := strings.Fields(name)
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

// NormalizeHybridChar substitutes hybrid chars 'X' or 'x' with
// the multiplication sign char.
func NormalizeHybridChar(bs []byte) []byte {
	hybridChar := []byte("$1×$2")
	res := hybridCharRe1.ReplaceAll(bs, hybridChar)
	res = hybridCharRe2.ReplaceAll(res, hybridChar)
	return res
}

// Annotation returns index where unparsed part starts. In case if
// the full string can be parsed, returns returns the index of the end of the
// input.
func Annotation(bs []byte) int {
	i := len(bs)
	if IsException(string(bs), AnnotationException) {
		return i
	}
	regexps := []*regexp.Regexp{
		notesRe, taxonConceptsRe1, taxonConceptsRe2, taxonConceptsRe3,
		nomenConceptsRe, lastWordJunkRe, stopWordsRe,
	}
	for _, r := range regexps {
		loc := r.FindIndex(bs[0:i])
		if len(loc) > 0 {
			i = loc[0]
		}
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
