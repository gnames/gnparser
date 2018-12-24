package str

import (
	"fmt"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// ToASCII converts a UTF-8 diacritics to corresponding ASCII chars.
func ToASCII(s string) (string, error) {
	var t = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	res, _, err := transform.String(t, s)
	return res, err
}

// JoinStrings contatenates two strings with a separator. If either of the
// strings is empty, then the value of another string is returned instead
// of concatenation.
func JoinStrings(s1 string, s2 string, sep string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	return fmt.Sprintf("%s%s%s", s1, sep, s2)
}
