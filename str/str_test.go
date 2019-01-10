package str_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "gitlab.com/gogna/gnparser/str"
)

var _ = Describe("string tools", func() {
	DescribeTable("ToASCII",
		func(s string, expected string) {
			res, _ := ToASCII([]byte(s))
			Expect(string(res)).To(Equal(expected))
		},
		Entry("Döringina", "Döringina", "Doeringina"),
		Entry("Aëtosaurus", "Aëtosaurus", "Aetosaurus"),
		Entry("thomæ", "thomæ", "thomae"),
		Entry("many ö", "ööö", "oeoeoe"),
	)

	DescribeTable("NumToStr",
		func(num string, expected string) {
			res := NumToStr(num)
			Expect(res).To(Equal(expected))
		},
		Entry("1", "1", "uni"),
		Entry("2", "2", "bi"),
		Entry("3", "3", "tri"),
		Entry("4", "4", "quadri"),
		Entry("5", "5", "quinque"),
		Entry("6", "6", "sex"),
		Entry("7", "7", "septem"),
		Entry("8", "8", "octo"),
		Entry("9", "9", "novem"),
		Entry("10", "10", "decem"),
		Entry("11", "11", "undecim"),
		Entry("12", "12", "duodecim"),
		Entry("13", "13", "tredecim"),
		Entry("14", "14", "quatuordecim"),
		Entry("15", "15", "quindecim"),
		Entry("16", "16", "sedecim"),
		Entry("17", "17", "septendecim"),
		Entry("18", "18", "octodecim"),
		Entry("19", "19", "novemdecim"),
		Entry("20", "20", "viginti"),
		Entry("21", "21", "vigintiuno"),
		Entry("22", "22", "vigintiduo"),
		Entry("23", "23", "vigintitre"),
		Entry("24", "24", "vigintiquatuor"),
		Entry("25", "25", "vigintiquinque"),
		Entry("26", "26", "vigintisex"),
		Entry("27", "27", "vigintiseptem"),
		Entry("28", "28", "vigintiocto"),
		Entry("30", "30", "triginta"),
		Entry("31", "31", "trigintauno"),
		Entry("32", "32", "trigintaduo"),
		Entry("38", "38", "trigintaocto"),
		Entry("40", "40", "quadraginta"),
		Entry("400", "400", "400"),
		Entry("something", "something", "something"),
	)

	DescribeTable("FixAllCaps",
		func(caps string, expected string) {
			res := FixAllCaps(caps)
			Expect(res).To(Equal(expected))
		},
		Entry("KURNAKOV", "KURNAKOV", "Kurnakov"),
		Entry("GÓMEZ-BOLEA", "GÓMEZ-BOLEA", "Gómez-Bolea"),
		Entry("hello", "hello", "hello"),
	)
})
