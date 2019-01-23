package preprocess

import (
	"bytes"
	"io"
	"sync"

	"golang.org/x/net/html"
)

var tags = map[string]struct{}{
	"i":     struct{}{},
	"small": struct{}{},
	"br":    struct{}{},
	"em":    struct{}{},
	"b":     struct{}{},
}

// CleanupStream takes input and output string channels, and feeds output with
// pipe delimited strings with original name on the left and cleaned up name
// on the right from the pipe.
func CleanupStream(in <-chan string, out chan<- string, wn int) {
	var wg sync.WaitGroup
	wg.Add(wn)
	for i := 0; i < wn; i++ {
		go cleanupWorker(in, out, &wg)
	}
	wg.Wait()
	close(out)
}

func cleanupWorker(in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range in {
		res := StripTags(s)
		out <- s + "|" + res
	}
}

// StripTags takes a slice of bytes and returns a string with common
// tags removed and html entities escaped. It does keep all uncommon tags
// intact to let parser deal with them.
func StripTags(s string) string {
	var buff bytes.Buffer
	r := bytes.NewReader([]byte(s))

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
