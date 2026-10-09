package stemmer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// AbbrKey returns a lookup key for a name with an abbreviated genus,
// such as "C. elegans" or "P. t. altaica". The key is the lowercase
// first letter of the genus, a dash, and the stem of the epithet:
//
//	AbbrKey("C.", "elegans")             -> "c-elegans"
//	AbbrKey("Carex", "alba")             -> "c-alb"
//	AbbrKey("P.", "altaica")             -> "p-altaic"
//
// The genus can be full or abbreviated, the epithet can be a
// specific or an infraspecific one. The epithet is stemmed exactly as
// in StemCanonical, so keys built from text match keys built from
// canonical forms, and gender endings (alba, albus, album) give the
// same key. It returns an empty string if genus, epithet, or the
// stem of the epithet is empty (for example "que" stems to nothing).
//
// Databases that store these keys (for example gndb) and clients that
// build them from text (for example gnames) must both use this
// function with the same version of gnparser.
func AbbrKey(genus, epithet string) string {
	genus = strings.TrimSpace(genus)
	epithet = strings.ToLower(strings.TrimSpace(epithet))
	if genus == "" || epithet == "" {
		return ""
	}

	first, _ := utf8.DecodeRuneInString(genus)
	letter := string(unicode.ToLower(first))

	// StemCanonical keeps the first word as is and stems the rest.
	// A placeholder genus makes it stem the epithet exactly as it
	// would inside a canonical form.
	words := strings.Fields(StemCanonical("X " + epithet))
	if len(words) < 2 {
		return ""
	}
	stem := words[len(words)-1]

	return letter + "-" + stem
}
