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

var notesRe = regexp.MustCompile(
	`(?i)\s+(species\s+group|species\s+complex|group|author)\b.*$`,
)
var taxonConceptsRe1 = regexp.MustCompile(
	`(?i)\s+(sensu|auct|sec|near|str)\.?\b.*$`,
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
	`\s+(of[\W_]|\(?ht\.?\W|\(?hort\.?\W|spec\.|nov\s+spec|cv\.?\W).*$`,
)

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
	if !VirusLikeName(name) {
		pr.Virus = IsVirus(bs[0:i])
	}
	if pr.Virus {
		pr.NoParse = true
		return pr
	}
	if name == "Candidatus" || strings.HasPrefix("Candidatus ", name) {
		return pr
	}
	pr.NoParse = NoParse(bs[0:i])
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

// LikeVirus takes a string and checks it against known species that can
// easily be misparsed as viruses. If the string belongs to one of such species
// returns true.
// The following names are covered:
//    Aspilota vector Belokobylskij, 2007
//    Ceylonesmus vector Chamberlin, 1941
//    Cryptops (Cryptops) vector Chamberlin, 1939
//    Culex vector Dyar & Knab, 1906
//    Dasyproctus cevirus Leclercq, 1963
//    Desmoxytes vector (Chamberlin, 1941)
//    Dicathais vector Thornley, 1952
//    Euragallia prion Kramer, 1976
//    Exochus virus Gauld & Sithole, 2002
//    Hilara vector Miller, 1923
//    Microgoneplax prion Castro, 2007
//    Neoaemula vector Mackinnon, Hiller, Long & Marshall, 2008
//    Ophion virus Gauld & Mitchell, 1981
//    Psenulus trevirus Leclercq, 1961
//    Tidabius vector Chamberlin, 1931

func VirusLikeName(name string) bool {
	names := map[string]string{
		"Aspilota":      "vector",
		"Ceylonesmus":   "vector",
		"Cryptops":      "vector",
		"Culex":         "vector",
		"Dasyproctus":   "cevirus",
		"Desmoxytes":    "vector",
		"Dicathais":     "vector",
		"Euragallia":    "prion",
		"Exochus":       "virus",
		"Hilara":        "vector",
		"Microgoneplax": "prion",
		"Neoaemula":     "vector",
		"Ophion":        "virus",
		"Psenulus":      "trevirus",
		"Tidabius":      "vector",
	}
	words := strings.Fields(name)
	if len(words) < 2 {
		return false
	}
	if epithet, ok := names[words[0]]; ok {
		for _, w := range words[1:] {
			if strings.HasPrefix(w, epithet) {
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
