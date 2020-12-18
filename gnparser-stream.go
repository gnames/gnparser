package gnparser

import (
	"sync"

	"github.com/gnames/gnparser/pb"
)

// ParseResult structure contains parsing output and/or error generated
// by the parser.
type ParseResult struct {
	Input  string
	Output string
	Error  error
}

// ParseStream function takes input/output channels to do concurrent
// parsing jobs. Output is pushed as ParseResult objects.
func ParseStream(jobs int, in <-chan string, out chan<- *ParseResult,
	opts ...Option) {
	var wg sync.WaitGroup
	wg.Add(jobs)
	for i := 0; i < jobs; i++ {
		go parserWorker(i, in, out, &wg, opts...)
	}
	wg.Wait()
	close(out)
}

func parserWorker(i int, in <-chan string, out chan<- *ParseResult,
	wg *sync.WaitGroup, opts ...Option) {
	gnp := NewGNparser(opts...)
	defer wg.Done()
	for s := range in {
		res, err := gnp.ParseAndFormat(s)
		if err != nil {
			out <- &ParseResult{Input: s, Output: "", Error: err}
		}
		out <- &ParseResult{Input: s, Output: res, Error: nil}
	}
}

// ParseStreamToObjects function takes input/output channels to do concurrent
// parsing to object jobs. Output is pushed as ParseObjectResult objects.
func ParseStreamToObjects(jobs int, in <-chan string,
	out chan<- *pb.Parsed, opts ...Option) {
	var wg sync.WaitGroup
	wg.Add(jobs)
	for i := 0; i < jobs; i++ {
		go parserObjectWorker(i, in, out, &wg, opts...)
	}
	wg.Wait()
	close(out)
}

func parserObjectWorker(i int, in <-chan string,
	out chan<- *pb.Parsed, wg *sync.WaitGroup, opts ...Option) {
	gnp := NewGNparser(opts...)
	defer wg.Done()
	for s := range in {
		out <- gnp.ParseToObject(s)
	}
}
