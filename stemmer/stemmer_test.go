package stemmer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "gitlab.com/gogna/gnparser/stemmer"
)

var _ = Describe("Stemmer", func() {
	Describe("Stem", func() {
		It("treats que suffix with exceptions", func() {
			Expect(Stem("detorque").Stem).To(Equal("detorque"))
			Expect(Stem("somethingque").Stem).To(Equal("something"))
		})
		It("removes suffixes correctly", func() {
			for k, v := range stemsDict {
				Expect(Stem(k).Stem).To(Equal(v))
			}
		})

		DescribeTable("StemCanonical",
			func(given string, expected string) {
				Expect(StemCanonical(given)).To(Equal(expected))
			}, []TableEntry{
				Entry("Uninomial", "Pomatomus", "Pomatomus"),
				Entry("Binomial1", "Betula naturae", "Betula natur"),
				Entry("Binomial2", "Betula alba", "Betula alb"),
				Entry("Trinomial", "Betula alba naturae", "Betula alb natur"),
			}...,
		)
	})

})
