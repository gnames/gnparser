package gnparser

import (
	"sync"

	"gitlab.com/gogna/gnparser/grammar"
)

// ParseResult structure contains parsing output and/or error generated
// by the parser.
type ParseResult struct {
	Output string
	Error  error
}

// ParseStream function takes input/output channels to do concurrent
// parsing jobs. Output is pushed as ParseResult objects.
func (gnp *GNparser) ParseStream(in <-chan string, out chan<- *ParseResult,
	opts ...Option) {
	var wg sync.WaitGroup
	wg.Add(gnp.workersNum)
	for i := 0; i < gnp.workersNum; i++ {
		go gnp.parserWorker(i, in, out, &wg, opts...)
	}
	wg.Wait()
	close(out)
}

func (gnp *GNparser) parserWorker(i int, in <-chan string, out chan<- *ParseResult,
	wg *sync.WaitGroup, opts ...Option) {
	e := &grammar.Engine{}
	e.Init()
	defer wg.Done()
	for s := range in {
		e.Buffer = s
		e.Reset()
		err := e.Parse()
		if err != nil {
			out <- &ParseResult{Output: "", Error: err}
		}
		out <- &ParseResult{Output: e.ParsedName(), Error: nil}
	}
}
