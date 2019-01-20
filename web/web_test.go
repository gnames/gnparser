package web

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Web", func() {
	DescribeTable("namesFromText",
		func(s string, expected []string) {
			res := namesFromText(s)
			Expect(res).To(Equal(expected))
		},
		Entry("Bubo|Homo", "Bubo|Homo", []string{"Bubo", "Homo"}),
		Entry("Bubo", "Bubo", []string{"Bubo"}),
		Entry("Bubo\r\n", "Bubo\r\n", []string{"Bubo"}),
		Entry("Bubo\r\nHomo", "Bubo\r\nHomo", []string{"Bubo", "Homo"}),
		Entry("\r\nBubo\r\nHomo\n\n", "\r\nBubo\r\nHomo\n\n", []string{"Bubo", "Homo"}),
		Entry("Bubo\r\nHomo   \n", "\r\nBubo\r\nHomo   \n", []string{"Bubo", "Homo"}),
	)
})
