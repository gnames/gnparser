package gnparser

import (
	"bytes"
	"log"
	"runtime"
	"strings"

	"gitlab.com/gogna/gnparser/grammar"
	"gitlab.com/gogna/gnparser/output"
)

// GNparser is responsible for parsing operations.
type GNparser struct {
	// workersNum defines the number of goroutines running parser in parallel.
	workersNum int
	// format defines the output format of the parser.
	format
	// e keeps parsing engine
	parser *grammar.Engine
}

// Option is a function that creates a new option for GNparser.
type Option func(*GNparser)

// WorkersNum Option sets the quantity of workers to run parsing jobs.
func WorkersNum(wn int) Option {
	return func(gnp *GNparser) {
		gnp.workersNum = wn
	}
}

// Format Option sets the output format to return/display parsing results.
func Format(f string) Option {
	return func(gnp *GNparser) {
		fo := newFormat(f)
		gnp.format = fo
	}
}

// NewGNparser constructor function takes options and returns
// configured GNparser.
func NewGNparser(opts ...Option) GNparser {
	gnp := GNparser{workersNum: runtime.NumCPU(), format: Compact}
	for _, opt := range opts {
		opt(&gnp)
	}
	e := &grammar.Engine{Buffer: ""}
	e.Init()
	gnp.parser = e
	return gnp
}

func (gnp *GNparser) WorkersNum() int {
	return gnp.workersNum
}

// Parse function parses input using GNparser's supplied options.
// The abstract syntax tree formed by the parser is stored in an
// `sn` private field.
func (gnp *GNparser) Parse(s string) error {
	gnp.parser.Buffer = s
	gnp.parser.Reset()
	err := gnp.parser.Parse()
	if err != nil {
		log.Println(s)
		log.Printf("No parse for '%s': %s", s, err)
		return err
	}
	gnp.parser.OutputAST()
	return nil
}

// ParseAndFormat function parses input and formats results according
// to format setting of GNparser.
func (gnp *GNparser) ParseAndFormat(s string) (string, error) {
	err := gnp.Parse(s)
	if err != nil {
		return "", err
	}
	var bs []byte
	switch gnp.format {
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
		s = strings.Join(gnp.ToSlice(), "|")
	case Debug:
		s = string(gnp.Debug())
	}
	return s, nil
}

// ToPrettyJSON function creates pretty JSON output out of parsed results.
func (gnp *GNparser) ToPrettyJSON() ([]byte, error) {
	gnp.parser.NewScientificNameNode()
	o := output.NewOutput(gnp.parser.SN)
	return o.ToJSON(true)
}

// ToJSON function creates a 'compact' output out of parsed results.
func (gnp *GNparser) ToJSON() ([]byte, error) {
	gnp.parser.NewScientificNameNode()
	o := output.NewOutput(gnp.parser.SN)
	return o.ToJSON(false)
}

// ToSlice function creates a flat simplified output of parsed results.
func (gnp *GNparser) ToSlice() []string {
	gnp.parser.NewScientificNameNode()
	so := output.NewSimpleOutput(gnp.parser.SN)
	return so.ToSlice()
}

// Debug returns byte representation of complete and 'output' syntax trees.
func (gnp *GNparser) Debug() []byte {
	var b bytes.Buffer
	b.WriteString("\n*** Complete Syntax Tree ***\n")
	gnp.parser.AST().PrettyPrint(&b, gnp.parser.Buffer)
	b.WriteString("\n*** Output Syntax Tree ***\n")
	gnp.parser.PrintOutputSyntaxTree(&b)
	return b.Bytes()
}

// Version function returns version number of `gnparser`.
func Version() string {
	return output.Version
}

// Build returns date and time when gnparser was built
func Build() string {
	return output.Build
}
