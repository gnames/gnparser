package gnparser

import (
	"context"
	"sync"

	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/config"

	"github.com/gnames/gnparser/entity/input"
	output "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/parser"
)

// GNparser is responsible for parsing operations.
type gnparser struct {
	// cfg keeps gnparser settings.
	cfg config.Config

	// nameString keeps parsed string
	nameString string

	// parser keeps parsing engine
	parser parser.Parser
}

// NewGNparser constructor function takes options and returns
// configured GNparser.
func NewGNParser(cfg config.Config) GNParser {
	gnp := gnparser{cfg: cfg}
	gnp.parser = parser.NewParser()
	return gnp
}

// Parse function parses input string according to configuraions.
// It takes a string and returns an output.Parsed object.
func (gnp gnparser) ParseName(s string) output.Parsed {
	ver := Version
	if gnp.cfg.IsTest {
		ver = "test_version"
	}
	sciNameNode := gnp.parser.PreprocessAndParse(s, ver, gnp.cfg.IgnoreHTMLTags)
	res := sciNameNode.ToOutput(gnp.cfg.WithDetails)
	return res
}

// ParseNames function takes input names and returns parsed results.
func (gnp gnparser) ParseNames(names []string) []output.Parsed {
	res := make([]output.Parsed, len(names))
	jobsNum := gnp.cfg.JobsNum
	chOut := make(chan output.ParseResult)
	var wgIn, wgOut sync.WaitGroup
	wgIn.Add(jobsNum)
	wgOut.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chIn := loadNames(ctx, names)

	for i := jobsNum; i > 0; i-- {
		go gnp.parseWorker(ctx, chIn, chOut, &wgIn)
	}

	go func() {
		defer wgOut.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-chOut:
				if !ok {
					return
				}
				res[v.Idx] = v.Parsed
			}
		}
	}()

	wgIn.Wait()
	close(chOut)
	wgOut.Wait()
	return res
}

func (gnp gnparser) Format() format.Format {
	return gnp.cfg.Format
}

func (gnp gnparser) ChangeConfig(opts ...config.Option) GNParser {
	for i := range opts {
		opts[i](&gnp.cfg)
	}
	return gnp
}

// Version function returns version number of `gnparser`.
func (gnp gnparser) GetVersion() gn.Version {
	res := gn.Version{
		Version: Version,
		Build:   Build,
	}
	if gnp.cfg.IsTest {
		res.Version = "test_version"
	}
	return res
}

func (gnp gnparser) parseWorker(
	ctx context.Context,
	chIn <-chan input.Name,
	chOut chan<- output.ParseResult,
	wgIn *sync.WaitGroup,
) {
	defer wgIn.Done()
	gnp.parser = parser.NewParser()

	for v := range chIn {
		parsed := gnp.ParseName(v.NameString)
		select {
		case <-ctx.Done():
			return
		case chOut <- output.ParseResult{Idx: v.Index, Parsed: parsed}:
		}
	}
}

func loadNames(ctx context.Context, names []string) <-chan input.Name {
	chIn := make(chan input.Name)
	go func() {
		defer close(chIn)
		for i := range names {
			select {
			case <-ctx.Done():
				return
			case chIn <- input.Name{Index: i, NameString: names[i]}:
			}
		}
	}()
	return chIn
}
