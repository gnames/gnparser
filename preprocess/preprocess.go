package preprocess

import "regexp"

var hybridCharRe1 = regexp.MustCompile(`(^)[Xx](\p{Lu})`)
var hybridCharRe2 = regexp.MustCompile(`(\s|^)[Xx](\s|$)`)

func NormalizeHybridChar(s string) string {
	hybridChar := "$1×$2"
	res := hybridCharRe1.ReplaceAllString(s, hybridChar)
	res = hybridCharRe2.ReplaceAllString(res, hybridChar)
	return res
}
