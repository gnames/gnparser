package gnparser

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"gitlab.com/gogna/gnparser/preprocess"
)

var _ = FDescribe("grammar", func() {
	DescribeTable("parsing rules execution",
		func(s string, expected string) {
			Expect(s).To(Equal(expected))
		}, astEntries()...,
	)
})

var _ = Describe("GNparser", func() {
	DescribeTable("predictable outputs",
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
	for _, v := range tests {
		gnp.Parse(v.NameString)
		res, err := gnp.ToJSON()
		if err != nil {
			fmt.Println(v.NameString)
			panic(err)
		}
		json := strings.Replace(string(res), "\\u0026", "&", -1)

		gnp.Parse(v.NameString)
		simple := strings.Join(gnp.ToSlice(), "|")
		te := Entry(v.NameString, json, v.Compact, simple, v.Simple)
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
	for _, v := range tests {
		gnp.parser.Buffer = preprocess.NormalizeHybridChar(v.NameString)
		gnp.parser.Reset()
		gnp.parser.Parse()
		parsedStr := gnp.parser.ParsedName()
		te := Entry(v.NameString, parsedStr, v.Parsed)
		entries = append(entries, te)
	}
	return entries
}
