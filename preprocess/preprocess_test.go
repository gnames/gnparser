package preprocess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/extensions/table"
	. "gitlab.com/gogna/gnparser/preprocess"
)

var _ = Describe("Preprocess", func() {
	DescribeTable("normalization of hybrid char",
		func(s string, expected string) {
			Expect(NormalizeHybridChar(s)).To(Equal(expected))
		},
		Entry(
			"'×', no space at the start",
			"×Agropogon P. Fourn. 1934",
			"×Agropogon P. Fourn. 1934",
		),
		Entry(
			"'x', no space at the start",
			"xAgropogon P. Fourn. 1934",
			"×Agropogon P. Fourn. 1934",
		),
		Entry(
			"'X', no space at the start",
			"XAgropogon P. Fourn. 1934",
			"×Agropogon P. Fourn. 1934",
		),
		Entry(
			"'×', space at the start",
			"× Agropogon P. Fourn. 1934",
			"× Agropogon P. Fourn. 1934",
		),
		Entry(
			"'x', space at the start",
			"x Agropogon P. Fourn. 1934",
			"× Agropogon P. Fourn. 1934",
		),
		Entry(
			"'X', space at the start",
			"X Agropogon P. Fourn. 1934",
			"× Agropogon P. Fourn. 1934",
		),
		Entry(
			"'×', no space at species",
			"Mentha ×smithiana ",
			"Mentha ×smithiana ",
		),
		Entry(
			"'X', spaces at species",
			"Asplenium X inexpectatum",
			"Asplenium × inexpectatum",
		),
		Entry(
			"'x', spaces at species",
			"Salix x capreola Andersson",
			"Salix × capreola Andersson",
		),
		Entry(
			"'x', spaces in formula",
			"Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939",
			"Asplenium rhizophyllum DC. × ruta-muraria E.L. Braun 1939",
		),
		// This one is brittle!
		Entry(
			"'X', spaces in formula",
			"Arthopyrenia hyalospora Hall X Hydnellum scrobiculatum D.E. Stuntz",
			"Arthopyrenia hyalospora Hall × Hydnellum scrobiculatum D.E. Stuntz",
		),
		Entry(
			"'x', in the end",
			"Arthopyrenia hyalospora x",
			"Arthopyrenia hyalospora ×",
		),
	)
})
