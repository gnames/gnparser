package preprocess

import (
	"bytes"
	"io"
	"unicode"

	"golang.org/x/net/html"
)

var tags = map[string]struct{}{
	"i":     struct{}{},
	"small": struct{}{},
	"br":    struct{}{},
	"em":    struct{}{},
	"b":     struct{}{},
}

// UnderscoreToSpace takes a slice of bytes. If it finds that the string
// contains underscores, but not spaces, it substitutes underscores to spaces
// in the slice. In case if any spaces are present, the slice is returned
// unmodified.
func UnderscoreToSpace(bs []byte) ([]byte, error) {
	reader := bytes.NewReader(bs)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return bs, err
		}
		if unicode.IsSpace(r) {
			return bs, nil
		}
	}

	for i, v := range bs {
		if v == '_' {
			bs[i] = ' '
		}
	}
	return bs, nil
}

// StripTags takes a slice of bytes and returns a string with common
// tags removed and html entities escaped. It does keep all uncommon tags
// intact to let parser deal with them.
func StripTags(bs []byte) string {
	var buff bytes.Buffer
	r := bytes.NewReader(bs)

	tokenizer := html.NewTokenizer(r)
	for {
		if tokenizer.Next() == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return html.UnescapeString(buff.String())
			}
			return ""
		}
		tokenVal := string(tokenizer.Raw())

		token := tokenizer.Token()
		switch token.Type {
		case html.DoctypeToken:
		case html.CommentToken:
		case html.StartTagToken:
			if _, ok := tags[token.Data]; ok {
				break
			}
			buff.WriteString(tokenVal)

		case html.EndTagToken:
			if _, ok := tags[token.Data]; ok {
				break
			}
			buff.WriteString(tokenVal)

		case html.TextToken:
			buff.WriteString(tokenVal)

		default:
			return ""
		}
	}
}
