package dict_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnames/gnparser/dict"
)

var _ = Describe("Dict", func() {
	Describe("LoadDictionary", func() {
		d := LoadDictionary()
		It("loads bacterial dictionary", func() {
			Expect(len(d.Bacteria)).To(BeNumerically(">", 100))
		})
		It("finds non homopypic genus", func() {
			hom, ok := d.Bacteria["Sphingomonas"]
			Expect(ok).To(Equal(true))
			Expect(hom).To(Equal(false))
		})
		It("finds homotypic genus", func() {
			hom, ok := d.Bacteria["Arizona"]
			Expect(ok).To(Equal(true))
			Expect(hom).To(Equal(true))
		})
		It("does not find non-bacterial genus", func() {
			hom, ok := d.Bacteria["Homo"]
			Expect(ok).To(Equal(false))
			Expect(hom).To(Equal(false))
		})
		It("does not find not ICN author", func() {
			_, ok := d.AuthorICN["Arizona"]
			Expect(ok).To(Equal(false))
		})
		It("finds ICN author", func() {
			_, ok := d.AuthorICN["Abramov"]
			Expect(ok).To(Equal(true))
		})
	})
})
