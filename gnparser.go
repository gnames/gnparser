package gnparser

import (
	"bytes"
	"fmt"
	"runtime"

	"gitlab.com/gogna/gnparser/pb"
	"gitlab.com/gogna/gnparser/preprocess"

	"gitlab.com/gogna/gnparser/grammar"
	"gitlab.com/gogna/gnparser/output"
)

// GNparser is responsible for parsing operations.
type GNparser struct {
	// Format defines the output format of the parser.
	Format
	// workersNum defines the number of goroutines running parser in parallel.
	workersNum int
	// removeHTML indicates that HTML tags have to be removed.
	removeHTML bool
	// nameString keeps parsed string
	nameString string
	// isTest indicates that parsing is done for test purposes, so instead of
	// real version of the paraser output will contain "test_version" phrase.
	isTest bool
	// parser keeps parsing engine
	parser *grammar.Engine
}

// Option is a function that creates a new option for GNparser.
type Option func(*GNparser)

// OptWorkersNum Option sets the quantity of workers to run parsing jobs.
func OptWorkersNum(wn int) Option {
	return func(gnp *GNparser) {
		gnp.workersNum = wn
	}
}

// OptFormat Option sets the output format to return/display parsing results.
func OptFormat(f string) Option {
	return func(gnp *GNparser) {
		fo := newFormat(f)
		gnp.Format = fo
	}
}

// OptIsTest Option to substitute real version of the parser with 'test_version'
// string.
func OptIsTest() Option {
	return func(gnp *GNparser) {
		gnp.isTest = true
	}
}

// OptRemoveHTML Option is true of false. When true, the preprocess removes
// HTML tags from name-strings.
func OptRemoveHTML(r bool) Option {
	return func(gnp *GNparser) {
		gnp.removeHTML = r
	}
}

// NewGNparser constructor function takes options and returns
// configured GNparser.
func NewGNparser(opts ...Option) GNparser {
	gnp := GNparser{workersNum: runtime.NumCPU(), Format: Compact, removeHTML: true}
	for _, opt := range opts {
		opt(&gnp)
	}
	e := &grammar.Engine{Buffer: ""}
	e.Init()
	gnp.parser = e
	return gnp
}

// WorkersNum returns the number of workers for concurrent parsing.
func (gnp *GNparser) WorkersNum() int {
	return gnp.workersNum
}

// Parse function parses input using GNparser's supplied options.
// The abstract syntax tree formed by the parser is stored in an
// `gnp.parser.SN` field.
func (gnp GNparser) Parse(s string) {
	gnp.nameString = s
	tagsOrEntities := false
	if gnp.removeHTML {
		orig := gnp.nameString
		gnp.nameString = preprocess.StripTags(gnp.nameString)
		if orig != gnp.nameString {
			tagsOrEntities = true
		}
	}
	preproc := preprocess.Preprocess([]byte(gnp.nameString))
	if preproc.NoParse {
		gnp.parser.NewNotParsedScientificNameNode(preproc)
	}
	gnp.parser.Buffer = string(preproc.Body)
	gnp.parser.FullReset()
	if tagsOrEntities {
		gnp.parser.AddWarn(grammar.HTMLTagsEntitiesWarn)
	}
	if len(preproc.Tail) > 0 {
		gnp.parser.AddWarn(grammar.TailWarn)
	}
	if preproc.Underscore {
		gnp.parser.AddWarn(grammar.SpaceNonStandardWarn)
	}
	err := gnp.parser.Parse()
	if err != nil {
		gnp.parser.Error = err
		gnp.parser.NewNotParsedScientificNameNode(preproc)
	} else {
		gnp.parser.OutputAST()
		gnp.parser.NewScientificNameNode()
		if len(preproc.Tail) > 0 {
			gnp.parser.SN.Tail += string(preproc.Tail)
		}
	}
	gnp.parser.SN.AddVerbatim(s)
	gnp.parser.SN.ParserVersion = gnp.Version()
}

// ParseAndFormat function parses input and formats results according
// to format setting of GNparser.
func (gnp GNparser) ParseAndFormat(s string) (string, error) {
	var err error
	if gnp.Format == Debug {
		bs := gnp.Debug(s)
		return string(bs), nil
	}
	gnp.Parse(s)
	var bs []byte
	switch gnp.Format {
	case Compact:
		bs, err = gnp.ToJSON()
		if err != nil {
			return "", err
		}
		s = string(bs)
	case Pretty:
		bs, err = gnp.ToPrettyJSON()
		if err != nil {
			return "", err
		}
		s = string(bs)
	case Simple:
		s = output.ToCSV(gnp.ToSlice())
	}
	return s, nil
}

// ParseToObject function parses input and
// returns result as output.
func (gnp GNparser) ParseToObject(s string) *pb.Parsed {
	gnp.Parse(s)
	return pb.ToPB(output.NewOutput(gnp.parser.SN))
}

// ToPrettyJSON function creates pretty JSON output out of parsed results.
func (gnp GNparser) ToPrettyJSON() ([]byte, error) {
	o := output.NewOutput(gnp.parser.SN)
	return o.ToJSON(true)
}

// ToJSON function creates a 'compact' output out of parsed results.
func (gnp GNparser) ToJSON() ([]byte, error) {
	o := output.NewOutput(gnp.parser.SN)
	return o.ToJSON(false)
}

// ToSlice function creates a flat simplified output of parsed results.
func (gnp GNparser) ToSlice() []string {
	so := output.NewSimpleOutput(gnp.parser.SN)
	return so.ToSlice()
}

// Debug returns byte representation of complete and 'output' syntax trees.
func (gnp GNparser) Debug(s string) []byte {
	ppr := preprocess.Preprocess([]byte(s))
	var b bytes.Buffer
	if ppr.NoParse || ppr.Virus {
		b.WriteString("\n*** Preprocessing: NO PARSE ***\n")
		b.WriteString(fmt.Sprintf("\n%s\n", s))
		return b.Bytes()
	}
	gnp.parser.Buffer = string(ppr.Body)
	gnp.parser.FullReset()
	_ = gnp.parser.Parse()
	gnp.parser.OutputAST()
	b.WriteString("\n*** Complete Syntax Tree ***\n")
	gnp.parser.AST().PrettyPrint(&b, gnp.parser.Buffer)
	b.WriteString("\n*** Output Syntax Tree ***\n")
	gnp.parser.PrintOutputSyntaxTree(&b)
	return b.Bytes()
}

// ParsedName returns the string of parsed result without a tail.
func (gnp GNparser) ParsedName() string {
	return gnp.parser.ParsedName()
}

// Version function returns version number of `gnparser`.
func (gnp GNparser) Version() string {
	if gnp.isTest {
		return "test_version"
	}
	return output.Version
}

// Build returns date and time when gnparser was built
func (gnp GNparser) Build() string {
	return output.Build
}
