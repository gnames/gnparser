package gnparser

import (
	"log"
	"runtime"

	"gitlab.com/gogna/gnparser/grammar"
	"gitlab.com/gogna/gnparser/output"
)

// GNparser is responsible for parsing operations.
type GNparser struct {
	// workersNum defines the number of goroutines running parser in parallel.
	workersNum int
	// format defines the output format of the parser.
	format
	// sn keeps AST resulting parsing.
	// sn *grm.ScientificNameNode
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
	return gnp
}

func (gnp *GNparser) WorkersNum() int {
	return gnp.workersNum
}

// Parse function parses input using GNparser's supplied options.
// The abstract syntax tree formed by the parser is stored in an
// `sn` private field.
func (gnp *GNparser) Parse(s string) (*grammar.Engine, error) {
	e := &grammar.Engine{Buffer: s, Pretty: true}
	e.Init()
	err := e.Parse()
	if err != nil {
		log.Println(s)
		log.Printf("No parse for '%s': %s", s, err)
		return e, err
	}
	return e, nil
}

// ParseAndFormat function parses input and formats results according
// to format setting of GNparser.
func (gnp *GNparser) ParseAndFormat(s string) (string, error) {
	e, err := gnp.Parse(s)
	if err != nil {
		return "", err
	}
	return e.ParsedName(), nil
}

// 	var s []byte
// 	var err error
// 	err = gnp.Parse(n)
// 	if err != nil {
// 		return "", err
// 	}
// 	switch gnp.Format {
// 	case output.Compact:
// 		s, err = gnp.ToJSON()
// 		if err != nil {
// 			return "", err
// 		}
// 	case output.Pretty:
// 		s, err = gnp.ToPrettyJSON()
// 		if err != nil {
// 			return "", err
// 		}
// 	case output.Simple:
// 		s = []byte(strings.Join(gnp.ToSlice(), "|"))
// 	}
// 	return string(s), nil
// }

// // ToPrettyJSON function creates pretty JSON output out of parsed results.
// func (gnp *GNparser) ToPrettyJSON() ([]byte, error) {
// 	o := output.NewOutput(gnp.sn)
// 	return o.ToJSON(true)
// }

// // ToJSON function creates a 'compact' output out of parsed results.
// func (gnp *GNparser) ToJSON() ([]byte, error) {
// 	o := output.NewOutput(gnp.sn)
// 	return o.ToJSON(false)
// }

// // ToSlice function creates a flat simplified output of parsed results.
// func (gnp *GNparser) ToSlice() []string {
// 	so := output.NewSimpleOutput(gnp.sn)
// 	return so.ToSlice()
// }

// Version function returns version number of `gnparser`.
func Version() string {
	return output.Version
}

// Build returns date and time when gnparser was built
func Build() string {
	return output.Build
}
