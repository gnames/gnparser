package preprocess

import (
	"regexp"
)

var hybridCharRe1 = regexp.MustCompile(`(^)[Xx](\p{Lu})`)
var hybridCharRe2 = regexp.MustCompile(`(\s|^)[Xx](\s|$)`)

// var approxRe = regexp.MustCompile(`\s(monst\.|\?|((spp|nr|sp|aff|species)(\.\s?[a-z]|\s+[a-z]|$)))`)
var virusRe = regexp.MustCompile(
	`(?i)(\b|\d)` +
		`(ictv|[a-z]*virus(es)?|` +
		`particles?|vectors?|` +
		`(bacterio|viro)?phages?|` +
		`viroids?|prions?|[a-z]*npv|` +
		`(alpha|beta)?satellites?)\b`,
)
var noParseRe = regexp.MustCompile(
	`(^(Not|None|Unidentified)[\W_].*|.*[Ii]ncertae\s+[Ss]edis.*|[Ii]nc\.\s*[Ss]ed\.|phytoplasma\b|plasmids?\b|[^A-Z]RNA[^A-Z]*)`)
var notesRe = regexp.MustCompile(`(?i)\s+(species\s+group|species\s+complex|group|author)\b.*$`)
var taxonConceptsRe1 = regexp.MustCompile(`(?i)\s+(sensu|auct|sec|near|str)\.?\b.*$`)
var taxonConceptsRe2 = regexp.MustCompile(`(,\s*|\s+)(\(?s\.\s?s\.|\(?s\.\s?l\.|\(?s\.\s?str\.|\(?s\.\s?lat\.).*$`)
var taxonConceptsRe3 = regexp.MustCompile(`(?i)(,\s*|\s+)(pro parte|p\.\s?p\.)\s*$`)
var nomenConceptsRe = regexp.MustCompile(`(?i)(,\s*|\s+)(\(?(nomen|nom\.|comb\.)(\s.*)?)$`)
var lastWordJunkRe = regexp.MustCompile(`(?i)(,\s*|\s+)(var\.?|von|van|ined\.?|sensu|new|non|nec|nudum|ssp\.?|subsp|subgen|hybrid)\??\s*$`)

type Preprocessor struct {
	Virus       bool
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
	j := Annotation(bs[0:i])
	if j < i {
		pr.Annotation = true
		i = j
	}
	// j = Approximation(bs[0:i])
	// if j < i {
	// 	pr.Approximate = true
	// 	i = j
	// }
	pr.Virus = IsVirus(bs[0:i])
	if pr.Virus {
		pr.NoParse = true
		return pr
	}
	pr.NoParse = NoParse(bs[0:i])
	if pr.NoParse {
		return pr
	}
	pr.Body = NormalizeHybridChar(bs[0:i])
	pr.Tail = bs[i:]
	return pr
}

// NormalizeHybridChar substitutes hybrid chars 'X' or 'x' with
// the multiplication sign char.
func NormalizeHybridChar(bs []byte) []byte {
	hybridChar := []byte("$1×$2")
	res := hybridCharRe1.ReplaceAll(bs, hybridChar)
	res = hybridCharRe2.ReplaceAll(res, hybridChar)
	return res
}

// Approximation returns an index where an approximation tail starts.
// func Approximation(bs []byte) int {
// 	i := len(bs)
// 	loc := approxRe.FindIndex(bs)
// 	if len(loc) > 0 {
// 		i = loc[0]
// 	}
// 	return i
// }

// IsVirus returns if a string is a virus name.
func IsVirus(bs []byte) bool {
	return virusRe.Match(bs)
}

// NoParse retuns if a string need to be parsed.
func NoParse(bs []byte) bool {
	return noParseRe.Match(bs)
}

// Annotation returns index where unparsed part starts. In case if
// the full string can be parsed, returns returns the index of the end of the
// input.
func Annotation(bs []byte) int {
	i := len(bs)
	regexps := []*regexp.Regexp{
		notesRe, taxonConceptsRe1, taxonConceptsRe2, taxonConceptsRe3,
		nomenConceptsRe, lastWordJunkRe,
	}
	for _, r := range regexps {
		loc := r.FindIndex(bs[0:i])
		if len(loc) > 0 {
			i = loc[0]
		}
	}
	return i
}
