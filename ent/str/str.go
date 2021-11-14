// Package str provides functions for manipulating scientific name-strings.
package str

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CapitalizeName function capitalizes the first character of a name-string.
// It can be a useful option if the data is known to contain 'real' names, for
// example canonical forms, but they are provided with all letters in lower
// case.
func CapitalizeName(name string) string {
	runes := []rune(name)
	if len(runes) < 2 {
		return name
	}

	one := runes[0]
	two := runes[1]
	if unicode.IsUpper(one) || !unicode.IsLetter(one) {
		return name
	}
	if one == 'x' && (two == ' ' || unicode.IsUpper(two)) {
		return name
	}
	runes[0] = unicode.ToUpper(one)
	return string(runes)
}

// Normalize takes a string and returns normalized version of it.
// Normalize function should be indempotent.
func Normalize(s string) string {
	return ToASCII(s, Transliterations)
}

// ToASCII converts a UTF-8 diacritics to corresponding ASCII chars.
func ToASCII(s string, m map[rune]string) string {
	if s == "" {
		return s
	}
	b := []byte(s)
	tlBuf := bytes.NewBuffer(make([]byte, 0, len(b)*125/100))
	for i, w := 0, 0; i < len(b); i += w {
		r, width := utf8.DecodeRune(b[i:])
		if s, ok := m[r]; ok {
			tlBuf.WriteString(s)
		} else {
			tlBuf.WriteRune(r)
		}
		w = width
	}
	return tlBuf.String()
}

func IsBoldSurrogate(s string) bool {
	if len(s) < 5 {
		return false
	}
	s = strings.ToLower(s)
	return strings.Contains(s, "bold:")
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

// FixAllCaps converts all-caps authors names to capitalized version.
func FixAllCaps(s string) string {
	rs := []rune(s)
	res := make([]rune, len(rs))
	var prev rune
	for i, v := range rs {
		if i == 0 || prev == '-' {
			res[i] = v
			prev = v
			continue
		}
		res[i] = unicode.ToLower(v)
		prev = v
	}
	return string(res)
}

// NumToString converts numbers in old-style species names to their
// word equivalents.
func NumToStr(num string) string {
	if v, ok := nameNums[num]; ok {
		return v
	}
	return num
}

// Transliteration table is used to convert diacritical characters to their
// latin letter equivalents.
var Transliterations = map[rune]string{

	'À': "A", 'Â': "A", 'Ã': "A", 'Á': "A", 'Ç': "C", 'Č': "C", 'Ð': "D",
	'Ë': "E", 'É': "E", 'È': "E", 'Í': "I", 'Ì': "I", 'Ï': "I", 'Ł': "L",
	'Ň': "N", 'Ñ': "N", 'Ó': "O", 'Ò': "O", 'Ô': "O", 'Õ': "O", 'Ú': "U",
	'Ù': "U", 'Ŕ': "R", 'Ř': "R", 'Ŗ': "R", 'Š': "S", 'Ş': "S", 'Ž': "Z",
	'à': "a", 'â': "a", 'ã': "a", 'á': "a", 'ç': "c", 'č': "c", 'ë': "e",
	'é': "e", 'è': "e", 'ð': "d", 'í': "i", 'ì': "i", 'ï': "i", 'ł': "l",
	'ň': "n", 'ñ': "n", 'ó': "o", 'ò': "o", 'ô': "o", 'õ': "o", 'ú': "u",
	'ù': "u", 'û': "u", 'ŕ': "r", 'ř': "r", 'ŗ': "r", 'ſ': "s", 'š': "s",
	'ş': "s", 'ž': "z", '\'': "", '‘': "", '’': "", '.': "",
	'Æ': "Ae", 'Å': "Ao", 'Ä': "Ae", 'Ø': "Oe", 'Ö': "Oe", 'Þ': "Th",
	'Ü': "Ue", 'ß': "ss", 'æ': "ae", 'å': "ao", 'ä': "ae", 'ø': "oe",
	'ö': "oe", 'þ': "th", 'Œ': "Oe", 'œ': "oe", 'ü': "ue",
}

// GlobalTransliterations are applied not only to scientific names, but
// to the whole name-string.
var GlobalTransliterations = map[rune]string{
	'‘': "'", '’': "'", '`': "'", '´': "'",
}

var nameNums = map[string]string{
	"1":  "uni",
	"2":  "bi",
	"3":  "tri",
	"4":  "quadri",
	"5":  "quinque",
	"6":  "sex",
	"7":  "septem",
	"8":  "octo",
	"9":  "novem",
	"10": "decem",
	"11": "undecim",
	"12": "duodecim",
	"13": "tredecim",
	"14": "quatuordecim",
	"15": "quindecim",
	"16": "sedecim",
	"17": "septendecim",
	"18": "octodecim",
	"19": "novemdecim",
	"20": "viginti",
	"21": "vigintiuno",
	"22": "vigintiduo",
	"23": "vigintitre",
	"24": "vigintiquatuor",
	"25": "vigintiquinque",
	"26": "vigintisex",
	"27": "vigintiseptem",
	"28": "vigintiocto",
	"30": "triginta",
	"31": "trigintauno",
	"32": "trigintaduo",
	"38": "trigintaocto",
	"40": "quadraginta",
}
