package gnparser

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"gitlab.com/gogna/gnparser/preprocess"
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
})

func outputEntries() []TableEntry {
	var entries []TableEntry
	tests, err := testData()
	if err != nil {
		panic(err)
	}
	gnp := NewGNparser()
	for i, v := range tests {
		gnp.Parse(v.NameString)
		res, err := gnp.ToJSON()
		if err != nil {
			fmt.Println(v.NameString)
			panic(err)
		}
		json := strings.Replace(string(res), "\\u0026", "&", -1)

		gnp.Parse(v.NameString)
		simple := strings.Join(gnp.ToSlice(), "|")
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
			parsedStr := "np"
			te := Entry(testName, parsedStr, v.Parsed)
			entries = append(entries, te)
			continue
		}
		gnp.parser.Buffer = string(ppr.Body)
		gnp.parser.Reset()
		err := gnp.parser.Parse()
		parsedStr := "np"
		if err == nil {
			parsedStr = gnp.parser.ParsedName()
		}
		te := Entry(testName, parsedStr, v.Parsed)
		entries = append(entries, te)
	}
	return entries
}
