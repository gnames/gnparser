package gnparser

import (
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
	parser *parser.Engine
}

// NewGNparser constructor function takes options and returns
// configured GNparser.
func NewGNParser(cfg config.Config) GNParser {
	gnp := gnparser{cfg: cfg}
	e := &parser.Engine{Buffer: ""}
	e.Init()
	gnp.parser = e
	return gnp
}

// Parse function parses input string according to configuraions.
// It takes a string and returns an output.Parsed object.
func (gnp gnparser) ParseName(s string) output.Parsed {
	sciNameNode := gnp.parser.PreprocessAndParse(s, Version, gnp.cfg.KeepHTMLTags)
	res := sciNameNode.ToOutput(gnp.cfg.WithDetails)
	return res
}

// ParseNames function takes input names and returns parsed results.
func (gnp gnparser) ParseNames(names []string) []output.Parsed {
	res := make([]output.Parsed, len(names))
	jobsNum := gnp.cfg.JobsNum
	chIn := make(chan input.Name)
	chOut := make(chan output.ParseResult)
	var wgIn, wgOut sync.WaitGroup
	wgIn.Add(jobsNum)
	wgOut.Add(1)

	go func() {
		for i := range names {
			chIn <- input.Name{Index: i, NameString: names[i]}
		}
		close(chIn)
	}()

	for i := jobsNum; i > 0; i-- {
		go gnp.parseWorker(chIn, chOut, &wgIn)
	}

	go func() {
		defer wgOut.Done()
		for v := range chOut {
			res[v.Index] = v.Parsed
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
	chIn <-chan input.Name,
	chOut chan<- output.ParseResult,
	wgIn *sync.WaitGroup,
) {
	defer wgIn.Done()
	e := &parser.Engine{Buffer: ""}
	e.Init()
	gnp.parser = e

	for v := range chIn {
		parsed := gnp.ParseName(v.NameString)
		chOut <- output.ParseResult{Index: v.Index, Parsed: parsed}
	}
}
