// Package gnparser implements the main use-case of the project -- parsing
// scientific names. There are methods to parse one name at a time,
// a slice of names, or a stream of names. All methods return results in the
// same order as input. It is achieved by restoring the order after concurrent
// execution of the parsing process.
package gnparser

import (
	"context"
	"sync"

	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/config"

	"github.com/gnames/gnparser/entity/nameidx"
	"github.com/gnames/gnparser/entity/parsed"
	"github.com/gnames/gnparser/entity/parser"
)

// gnparser is an implementation of GNparser interface.
// It is responsible for main parsing operations.
type gnparser struct {
	// cfg keeps gnparser settings.
	cfg config.Config

	// parser keeps parsing engine
	parser parser.Parser
}

// New constructor function takes options organized into a
// configuration struct and returns an object that implements GNparser
// interface.
func New(cfg config.Config) GNparser {
	gnp := gnparser{cfg: cfg}
	gnp.parser = parser.NewParser()
	return gnp
}

// Parse function parses input string according to configurations.
// It takes a string and returns an parsed.Parsed object.
func (gnp gnparser) ParseName(s string) parsed.Parsed {
	ver := Version
	if gnp.cfg.IsTest {
		ver = "test_version"
	}
	sciNameNode := gnp.parser.PreprocessAndParse(s, ver, gnp.cfg.IgnoreHTMLTags)
	res := sciNameNode.ToOutput(gnp.cfg.WithDetails)
	return res
}

// ParseNames function takes input names and returns parsed results.
func (gnp gnparser) ParseNames(names []string) []parsed.Parsed {
	res := make([]parsed.Parsed, len(names))
	jobsNum := gnp.cfg.JobsNum
	chOut := make(chan parsed.ParsedWithIdx)
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

// Format returns the configured output format value.
func (gnp gnparser) Format() format.Format {
	return gnp.cfg.Format
}

// ChangeConfig allows change configuration of already created
// GNparser object.
func (gnp gnparser) ChangeConfig(opts ...config.Option) GNparser {
	for i := range opts {
		opts[i](&gnp.cfg)
	}
	return gnp
}

// Version function returns version number of `gnparser` and the timestamp
// of its build.
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
	chIn <-chan nameidx.NameIdx,
	chOut chan<- parsed.ParsedWithIdx,
	wgIn *sync.WaitGroup,
) {
	defer wgIn.Done()
	gnp.parser = parser.NewParser()

	for v := range chIn {
		parseRes := gnp.ParseName(v.NameString)
		select {
		case <-ctx.Done():
			return
		case chOut <- parsed.ParsedWithIdx{Idx: v.Index, Parsed: parseRes}:
		}
	}
}

func loadNames(ctx context.Context, names []string) <-chan nameidx.NameIdx {
	chIn := make(chan nameidx.NameIdx)
	go func() {
		defer close(chIn)
		for i := range names {
			select {
			case <-ctx.Done():
				return
			case chIn <- nameidx.NameIdx{Index: i, NameString: names[i]}:
			}
		}
	}()
	return chIn
}
