package str_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.com/gogna/gnparser/str"
)

var _ = Describe("Str", func() {
	var _ = Describe("ToASCII", func() {
		It("converts diacritics to ASCII", func() {
			data := [][]string{
				{"Döringina", "Doringina"},
				{"Aëtosaurus", "Aetosaurus"},
			}

			for _, v := range data {
				res, err := ToASCII(v[0])
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(v[1]))
			}
		})
	})

})
