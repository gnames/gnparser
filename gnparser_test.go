package gnparser

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("grammar", func() {
	Describe("Parse", func() {
		FIt("tokenizes data by parsing rules", func() {
			var currentNameString string
			tests, err := testData()
			Expect(len(tests)).To(BeNumerically(">", 0))
			Expect(err).NotTo(HaveOccurred())
			gnp := NewGNparser()
			for _, v := range tests {
				currentNameString = v.NameString
				gnp.Parse(v.NameString)
				parsedStr := gnp.parser.ParsedName()
				Expect(parsedStr).To(Equal(v.Parsed))
			}
			fmt.Println(currentNameString)
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
				gnp.Parse(v.NameString)
				Expect(err).NotTo(HaveOccurred())
				res, err := gnp.ToJSON()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(res)).To(BeNumerically(">", 300))
				json := strings.Replace(string(res), "\\u0026", "&", -1)
				fmt.Println(json)
				Expect(json).To(Equal(string(v.Compact)))
				gnp.Parse(v.NameString)
				simple := strings.Join(gnp.ToSlice(), "|")
				fmt.Println(simple)
				Expect(simple).To(Equal(v.Simple))
			}
		})
	})
})
