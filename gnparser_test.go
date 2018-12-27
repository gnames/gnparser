package gnparser

import (
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/gogna/gnparser/grammar"
)

var _ = Describe("grammar", func() {
	Describe("Parse", func() {
		It("tokenizes data by parsing rules", func() {
			tests, err := testData()
			Expect(len(tests)).To(BeNumerically(">", 0))
			Expect(err).NotTo(HaveOccurred())
			e := grammar.Engine{}
			e.Init()
			for _, v := range tests {
				e.Buffer = v.NameString
				e.Reset()
				err := e.Parse()
				parsedStr := e.ParsedName()
				e.PrintOutputSyntaxTree(os.Stdout)
				Expect(err).NotTo(HaveOccurred())
				Expect(parsedStr).To(Equal(v.Parsed))
			}
		})
	})
})

var _ = Describe("GNparser", func() {
	Describe("Parse", func() {
		It("parses test data", func() {
			tests, err := testData()
			Expect(len(tests)).To(BeNumerically(">", 0))
			Expect(err).NotTo(HaveOccurred())
			gnp := NewGNparser()
			for _, v := range tests {
				fmt.Println(v.NameString)
				err := gnp.Parse(v.NameString)
				Expect(err).NotTo(HaveOccurred())
				res, err := gnp.ToJSON()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(res)).To(BeNumerically(">", 300))
				json := strings.Replace(string(res), "\\u0026", "&", -1)
				fmt.Println(json)
				Expect(json).To(Equal(string(v.Compact)))
				Expect(strings.Join(gnp.ToSlice(), "|")).To(Equal(v.Simple))
			}
		})
	})
})
