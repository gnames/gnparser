package output

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	grm "gitlab.com/gogna/gnparser/grammar"
)

var _ = Describe("Output", func() {
	DescribeTable("ToCSV", func(r []string, expected string) {
		Expect(ToCSV(r)).To(Equal(expected))
	},
		Entry("simple case", []string{"abc"}, "abc"),
		Entry("", []string{`"abc"`}, `"""abc"""`),
		Entry("", []string{`a"b`}, `"a""b"`),
		Entry("", []string{`"a"b"`}, `"""a""b"""`),
		Entry("", []string{" abc"}, `" abc"`),
		Entry("", []string{"abc,def"}, `"abc,def"`),
		Entry("", []string{"abc", "def"}, "abc,def"),
		Entry("", []string{"abc\ndef"}, "\"abc\ndef\""),
		Entry("", []string{"abc\ndef"}, "\"abc\ndef\""),
		Entry("", []string{"abc\rdef"}, "\"abc\rdef\""),
		Entry("", []string{""}, ""),
		Entry("", []string{"", ""}, ","),
		Entry("", []string{"", "", ""}, ",,"),
		Entry("", []string{"", "", "a"}, ",,a"),
		Entry("", []string{"", "a", ""}, ",a,"),
		Entry("", []string{"", "a", "a"}, ",a,a"),
		Entry("", []string{"a", "", ""}, "a,,"),
		Entry("", []string{"a", "", "a"}, "a,,a"),
		Entry("", []string{"a", "a", ""}, "a,a,"),
		Entry("", []string{"a", "a", "a"}, "a,a,a"),
		Entry("", []string{`\.`}, "\"\\.\""),
		Entry("", []string{"x09\x41\xb4\x1c", "aktau"}, "x09\x41\xb4\x1c,aktau"),
		Entry("", []string{",x09\x41\xb4\x1c", "aktau"}, "\",x09\x41\xb4\x1c\",aktau"),
	)
})

var _ = Describe("Private Functions", func() {
	Describe("prepareWarnings", func() {
		It("sorts incoming data consistently", func() {
			ws := []grm.Warning{
				grm.YearParensWarn,
				grm.YearCharWarn,
				grm.CharBadWarn,
				grm.YearQuestionWarn,
				grm.TailWarn,
			}
			res := prepareWarnings(ws)
			output := []Warning{
				{
					Quality: 3,
					Message: "Unparsed tail",
				},
				{
					Quality: 2,
					Message: "Non-standard characters in canonical",
				},
				{
					Quality: 2,
					Message: "Year with latin character",
				},
				{
					Quality: 2,
					Message: "Year with parentheses",
				},
				{
					Quality: 2,
					Message: "Year with question mark",
				},
			}
			Expect(res).To(Equal(output))
		})
	})
})

// func randIntSlice(sl []int) []int {
// 	res := make([]int, len(sl))
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	for i, v := range r.Perm(len(sl)) {
// 		res[i] = sl[v]
// 	}
// 	return res
// }
