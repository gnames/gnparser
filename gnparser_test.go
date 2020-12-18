package gnparser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnparser/output"
	"github.com/gnames/gnparser/pb"
	"github.com/gnames/gnparser/preprocess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("grammar", func() {
	DescribeTable("AST formation",
		func(s string, expected string) {
			Expect(s).To(Equal(expected))
		}, astEntries()...,
	)
})

var _ = Describe("GNparser", func() {
	DescribeTable("full stack input to output",
		func(compactRes, compact, simpleRes, simple string) {
			Expect(compactRes).To(Equal(compact))
			Expect(simpleRes).To(Equal(simple))
		}, outputEntries()...,
	)

	Describe("ParseToObject", func() {
		It("returns output", func() {
			gnp := NewGNparser()
			o := gnp.ParseToObject("Homo sapiens")
			Expect(o.Parsed).To(Equal(true))
			Expect(o.Canonical.Simple).To(Equal("Homo sapiens"))
			Expect(o.Canonical.Stem).To(Equal("Homo sapiens"))
			switch d := o.Details.(type) {
			case *pb.Parsed_Species:
				Expect(d.Species.Genus).To(Equal("Homo"))
			default:
				Expect(2).To(Equal(3))
			}
		})

		It("parses hybrid formula", func() {
			gnp := NewGNparser()
			o := gnp.ParseToObject("Stanhopea tigrina Bateman ex Lindl. x S. ecornuta Lem.")
			Expect(o.Parsed).To(Equal(true))
			Expect(o.Cardinality).To(Equal(int32(0)))
			Expect(pb.NameType_name[int32(o.NameType)]).To(Equal("HYBRID_FORMULA"))
			Expect(o.Canonical.Full).To(Equal("Stanhopea tigrina × Stanhopea ecornuta"))
			Expect(o.Details).To(BeNil())
			det := o.DetailsHybridFormula
			Expect(len(det)).To(Equal(2))
			for _, v := range det {
				switch d := v.Element.(type) {
				case *pb.HybridFormula_Species:
					Expect(d.Species.Genus).To(Equal("Stanhopea"))
				default:
					Expect(2).To(Equal(3))
				}
			}
		})
	})
})

func outputEntries() []TableEntry {
	var entries []TableEntry
	tests, err := testData()
	if err != nil {
		panic(err)
	}
	gnp := NewGNparser(OptIsTest())
	for i, v := range tests {
		gnp.Parse(v.NameString)
		res, err := gnp.ToJSON()
		if err != nil {
			fmt.Println(v.NameString)
			panic(err)
		}
		json := string(res)

		gnp.Parse(v.NameString)
		simple := output.ToCSV(gnp.ToSlice())
		testName := fmt.Sprintf("%000d: |%s|", i+1, v.NameString)
		te := Entry(testName, json, v.Compact, simple, v.Simple)
		entries = append(entries, te)
	}
	return entries
}

func astEntries() []TableEntry {
	var entries []TableEntry
	tests, err := testData()
	if err != nil {
		fmt.Println(err)
	}
	gnp := NewGNparser()
	for i, v := range tests {
		testName := fmt.Sprintf("AST-%03d: |%s|", i+1, v.NameString)
		ppr := preprocess.Preprocess([]byte(v.NameString))
		if ppr.NoParse {
			parsedStr := "noparse"
			te := Entry(testName, parsedStr, v.Parsed)
			entries = append(entries, te)
			continue
		}
		gnp.parser.Buffer = string(ppr.Body)
		gnp.parser.FullReset()
		gnp.parser.Error = gnp.parser.Parse()
		parsedStr := gnp.parser.ParsedName()
		te := Entry(testName, parsedStr, v.Parsed)
		entries = append(entries, te)
	}
	return entries
}

// BenchmarkParse checks parsing event speed. Run it with:
// `go test -bench=. -benchmem -count=10 -run=XXX > bench.txt && benchstat bench.txt`
func BenchmarkParse(b *testing.B) {
	path := filepath.Join("testdata", "200k-lines.txt")
	count := 1000
	test := make([]string, count)
	gnp := NewGNparser()
	ops := []Option{OptFormat("simple")}
	gnpSimple := NewGNparser(ops...)
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if count == 0 {
			break
		}
		test = append(test, scanner.Text())
		count--
	}
	b.Run("ParseToObjectOnce", func(b *testing.B) {
		var p *pb.Parsed
		for i := 0; i < b.N; i++ {
			p = gnp.ParseToObject("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("ParseToObjectOnceWithInit", func(b *testing.B) {
		var p *pb.Parsed
		for i := 0; i < b.N; i++ {
			gnp1 := NewGNparser()
			p = gnp1.ParseToObject("Abarema clypearia (Jack) Kosterm., p.p.")
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})
	b.Run("ParseToObject", func(b *testing.B) {
		var p *pb.Parsed
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p = gnp.ParseToObject(v)
			}
		}
		_ = fmt.Sprintf("%v", p.Parsed)
	})

	b.Run("ParseAndFormat", func(b *testing.B) {
		var p string
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p, err = gnp.ParseAndFormat(v)
				if err != nil {
					panic(err)
				}
			}
		}
		_ = fmt.Sprintf("%d", len(p))
	})
	b.Run("ParseAndFormat(Simple)", func(b *testing.B) {
		var p string
		for i := 0; i < b.N; i++ {
			for _, v := range test {
				p, err = gnpSimple.ParseAndFormat(v)
				if err != nil {
					panic(err)
				}
			}
		}
		_ = fmt.Sprintf("%d", len(p))
	})
}
