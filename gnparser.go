// Package gnparser implements the main use-case of the project -- parsing
// scientific names. There are methods to parse one name at a time,
// a slice of names, or a stream of names. All methods return results in the
// same order as input. It is achieved by restoring the order after concurrent
// execution of the parsing process.
package gnparser

import (
	"context"
	"sync"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/gnvers"
	"github.com/gnames/gnparser/ent/nameidx"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/parser"
)

// gnparser is an implementation of GNparser interface.
// It is responsible for main parsing operations.
type gnparser struct {
	// cfg keeps gnparser settings.
	cfg Config

	// parser keeps parsing engine
	parser parser.Parser
}

// New constructor function takes options organized into a
// configuration struct and returns an object that implements GNparser
// interface.
func New(cfg Config) GNparser {
	gnp := gnparser{cfg: cfg}
	gnp.parser = parser.New()
	return gnp
}

// Debug returns byte representation of complete and 'output' syntax trees.
func (gnp gnparser) Debug(s string) []byte {
	return gnp.parser.Debug(s)
}

// Parse function parses input string according to configurations.
// It takes a string and returns an parsed.Parsed object.
func (gnp gnparser) ParseName(s string) parsed.Parsed {
	ver := Version
	if gnp.cfg.IsTest {
		ver = "test_version"
	}
	sciNameNode := gnp.parser.PreprocessAndParse(
		s, ver, gnp.cfg.IgnoreHTMLTags, gnp.cfg.WithCapitalization,
	)
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
		var count int
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-chOut:
				if !ok {
					return
				}
				if gnp.cfg.WithNoOrder {
					res[count] = v.Parsed
					count++
				} else {
					res[v.Idx] = v.Parsed
				}
			}
		}
	}()

	wgIn.Wait()
	close(chOut)
	wgOut.Wait()
	return res
}

// Format returns the configured output format value.
func (gnp gnparser) Format() gnfmt.Format {
	return gnp.cfg.Format
}

// ChangeConfig allows change configuration of already created
// GNparser object.
func (gnp gnparser) ChangeConfig(opts ...Option) GNparser {
	for i := range opts {
		opts[i](&gnp.cfg)
	}
	return gnp
}

// Version function returns version number of `gnparser` and the timestamp
// of its build.
func (gnp gnparser) GetVersion() gnvers.Version {
	version := Version
	build := Build
	if gnp.cfg.IsTest {
		version = "test_version"
	}
	return gnvers.Version{Version: version, Build: build}
}

func (gnp gnparser) parseWorker(
	ctx context.Context,
	chIn <-chan nameidx.NameIdx,
	chOut chan<- parsed.ParsedWithIdx,
	wgIn *sync.WaitGroup,
) {
	defer wgIn.Done()
	gnp.parser = parser.New()

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
