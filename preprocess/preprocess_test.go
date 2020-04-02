package preprocess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/extensions/table"
	. "gitlab.com/gogna/gnparser/preprocess"
)

var _ = Describe("Cleanup", func() {
	DescribeTable("StripTags",
		func(s string, expected string) {
			Expect(StripTags(s)).To(Equal(expected))
		},
		Entry("no html", "Hello", "Hello"),
		Entry("html tags", "<i>Hello</i>", "Hello"),
		Entry("html tags", "<I>Hello</I>", "Hello"),
		Entry("bad tag", "<!--", ""),
		Entry("bad tag with newline", "<!--\n", ""),
		Entry("keep other tags",
			"<code>Hello</code> & you",
			"<code>Hello</code> & you"),
		Entry("preserve case for other tags",
			"<CODE>Hello & you</CODE>",
			"<CODE>Hello & you</CODE>"),
		Entry("unknown tags", "<NA>Hello</NA> & you", "<NA>Hello</NA> & you"),
		Entry("entities", "Hello &amp; you", "Hello & you"),
	)

	Describe("StripTags no nil output", func() {
		It("does not return nil", func() {
			Expect(StripTags("<!--")).ToNot(Equal(nil))
			Expect(StripTags("<!--\r\n")).ToNot(Equal(nil))
		})
	})
})

var _ = Describe("Preprocess", func() {
	DescribeTable("NormalizeHybridChar",
		func(s string, expected string) {
			Expect(NormalizeHybridChar([]byte(s))).To(Equal([]byte(expected)))
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

	DescribeTable("VirusLikeName",
		func(s string, expected bool) {
			Expect(VirusLikeName(s)).To(Equal(expected))
		},
		Entry("name1", "Aspilota vector Belokobylskij, 2007", true),
		Entry("name2", "Ceylonesmus vector Chamberlin, 1941", true),
		Entry("name3", "Cryptops (Cryptops) vector Chamberlin, 1939", true),
		Entry("name4", "Culex vector Dyar & Knab, 1906", true),
		Entry("name5", "Dasyproctus cevirus Leclercq, 1963", true),
		Entry("name6", "Desmoxytes vector (Chamberlin, 1941)", true),
		Entry("name7", "Dicathais vector Thornley, 1952", true),
		Entry("name8", "Euragallia prion Kramer, 1976", true),
		Entry("name9", "Exochus virus Gauld & Sithole, 2002", true),
		Entry("name10", "Hilara vector Miller, 1923", true),
		Entry("name11", "Microgoneplax prion Castro, 2007", true),
		Entry("name12", "Neoaemula vector Mackinnon, Hiller, Long & Marshall, 2008", true),
		Entry("name13", "Ophion virus Gauld & Mitchell, 1981", true),
		Entry("name14", "Psenulus trevirus Leclercq, 1961", true),
		Entry("name15", "Tidabius vector Chamberlin, 1931", true),
		Entry("name16", "Ceylonesmus prion", false),
		Entry("name17", "Homo sapiens coronavirus", false),
	)

	DescribeTable("IsVirus",
		func(s string, itIs bool) {
			res := IsVirus([]byte(s))
			Expect(res).To(Equal(itIs))
		},
		Entry("No match", "Homo sapiens", false),
		Entry("Match word", "Arv1virus ", true),
		Entry("Match word", "Turtle herpesviruses", true),
		Entry("Match word", "Cre expression vector", true),
		Entry("Match word", "Abutilon mosaic vir. ICTV", true),
		Entry("Match word", "Aeromonas phage 65", true),
		Entry("Match word", "Apple scar skin viroid", true),
		Entry("Match word", "Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease", true),
		Entry("Match word", "Phi h-like viruses", true),
		Entry("Match word", "Viroids", true),
		Entry("Match word", "Human rhinovirus A11", true),
		Entry("Match word", "Gossypium mustilinum symptomless alphasatellite", true),
		Entry("Match word", "Bemisia betasatellite LW-2014", true),
		Entry("Match word", "Intracisternal A-particles", true),
		Entry("Match word", "Uranotaenia sapphirina NPV", true),
		Entry("Match word", "Spodoptera frugiperda MNPV", true),
		Entry("Match word", "Mamestra configurata NPV-A", true),
		Entry("Match word", "Bacteriophage PH75", true),
	)

	DescribeTable("NoParse",
		func(s string, itIs bool) {
			res := NoParse([]byte(s))
			Expect(res).To(Equal(itIs))
		},
		Entry("No match", "Homo sapiens", false),
		Entry("No word at the start", "Not Homo sapiens", true),
		Entry("Noword at the start", "Nothomo sapiens", false),
		Entry("Not word at the start", "Not Homo sapiens", true),
		Entry("None word at the start", "None Homo sapiens", true),
		Entry("Unidentified at the start", "Unidentified species", true),
		Entry("Incertae sedis1", "incertae sedis", true),
		Entry("Incertae sedis2", "Incertae Sedis", true),
		Entry("Incertae sedis3", "Something incertae sedis", true),
		Entry("Incertae sedis4", "Homo sapiens inc.sed.", true),
		Entry("Incertae sedis5", "Incertae sedis", true),
		Entry("Phytoplasma in the middle", "Homo sapiensphytoplasmaoid", false),
		Entry("Phytoplasma in the end", "Homo sapiensphytoplasma Linn", true),
		Entry("Phytoplasma in the end", "Homo sapiensphytoplasma Linn", true),
		Entry("Plasmid1", "E. coli plasmids", true),
		Entry("Plasmid2", "E. coli plasmidia", false),
		Entry("Plasmid3", "E. coli plasmid", true),
		Entry("RNA1", "E. coli RNA", true),
		Entry("RNA2", "E. coli 32RNA", true),
		Entry("RNA3", "KURNAKOV", false),
		Entry("RNA4", "E. coli mRNA", true),
	)

	DescribeTable("Annotations",
		func(s string, body string, tail string) {
			bs := []byte(s)
			i := Annotation(bs)
			Expect(string(bs[0:i])).To(Equal(body))
			Expect(string(bs[i:])).To(Equal(tail))
		},
		Entry("No tail", "Homo sapiens", "Homo sapiens", ""),
		Entry("No tail", "Homo sapiens S. S.", "Homo sapiens S. S.", ""),
		Entry("No tail", "Homo sapiens s. s.", "Homo sapiens", " s. s."),
		Entry("No tail", "Homo sapiens sensu Linn.", "Homo sapiens", " sensu Linn."),
		Entry("No tail", "Homo sapiens nomen nudum", "Homo sapiens", " nomen nudum"),
	)

	DescribeTable("UnderscoreToSpace",
		func(s string, expected string, changed bool) {
			bs := []byte(s)
			changed2, _ := UnderscoreToSpace(bs)
			Expect(string(bs)).To(Equal(expected))
			Expect(changed).To(Equal(changed2))
		},
		Entry("no nothing", "Hello", "Hello", false),
		Entry("has spaces", "Hello_you !", "Hello_you !", false),
		Entry("has spaces", "Hello_you\t!", "Hello_you\t!", false),
		Entry("has only underscores", "Hello_you_!_", "Hello you ! ", true),
	)

	Describe("Preprocess", func() {
		It("does not remove spaces at the start of a string", func() {
			name := "    Asplenium       × inexpectatum(E. L. Braun ex Friesner      )Morton"
			res := Preprocess([]byte(name))
			Expect(string(res.Body)).To(Equal(name))
		})
	})
})
