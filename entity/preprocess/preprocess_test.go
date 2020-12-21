package preprocess_test

import (
	"testing"

	ppr "github.com/gnames/gnparser/entity/preprocess"
	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {
	t.Run("StripTags", func(t *testing.T) {
		data := []struct {
			msg    string
			tags   string
			notags string
		}{
			{"no html", "Hello", "Hello"},
			{"html tags", "<i>Hello</i>", "Hello"},
			{"html tags", "<I>Hello</I>", "Hello"},
			{"bad tag", "<!--", ""},
			{"bad tag with newline", "<!--\n", ""},
			{"keep other tags",
				"<code>Hello</code> & you",
				"<code>Hello</code> & you"},
			{"preserve case for other tags",
				"<CODE>Hello & you</CODE>",
				"<CODE>Hello & you</CODE>"},
			{"unknown tags", "<NA>Hello</NA> & you", "<NA>Hello</NA> & you"},
			{"entities", "Hello &amp; you", "Hello & you"},
		}
		for _, v := range data {
			assert.Equal(t, ppr.StripTags(v.tags), v.notags, v.msg)
		}
	})
	t.Run("does not return nil", func(t *testing.T) {
		assert.NotNil(t, ppr.StripTags("<!--"))
		assert.NotNil(t, ppr.StripTags("<!--\r\n"))
	})
}

func TestPreprocess(t *testing.T) {
	t.Run("NormalizeHybridChar", func(t *testing.T) {
		data := []struct {
			msg string
			in  string
			out string
		}{
			{
				"'×', no space at the start",
				"×Agropogon P. Fourn. 1934",
				"×Agropogon P. Fourn. 1934",
			},
			{
				"'x', no space at the start",
				"xAgropogon P. Fourn. 1934",
				"×Agropogon P. Fourn. 1934",
			},
			{
				"'X', no space at the start",
				"XAgropogon P. Fourn. 1934",
				"×Agropogon P. Fourn. 1934",
			},
			{
				"'×', space at the start",
				"× Agropogon P. Fourn. 1934",
				"× Agropogon P. Fourn. 1934",
			},
			{
				"'x', space at the start",
				"x Agropogon P. Fourn. 1934",
				"× Agropogon P. Fourn. 1934",
			},
			{
				"'X', space at the start",
				"X Agropogon P. Fourn. 1934",
				"× Agropogon P. Fourn. 1934",
			},
			{
				"'×', no space at species",
				"Mentha ×smithiana ",
				"Mentha ×smithiana ",
			},
			{
				"'X', spaces at species",
				"Asplenium X inexpectatum",
				"Asplenium × inexpectatum",
			},
			{
				"'x', spaces at species",
				"Salix x capreola Andersson",
				"Salix × capreola Andersson",
			},
			{
				"'x', spaces in formula",
				"Asplenium rhizophyllum DC. x ruta-muraria E.L. Braun 1939",
				"Asplenium rhizophyllum DC. × ruta-muraria E.L. Braun 1939",
			},
			// This one is brittle!
			{
				"'X', spaces in formula",
				"Arthopyrenia hyalospora Hall X Hydnellum scrobiculatum D.E. Stuntz",
				"Arthopyrenia hyalospora Hall × Hydnellum scrobiculatum D.E. Stuntz",
			},
			{
				"'x', in the end",
				"Arthopyrenia hyalospora x",
				"Arthopyrenia hyalospora ×",
			},
		}

		for _, v := range data {
			assert.Equal(t, ppr.NormalizeHybridChar([]byte(v.in)),
				[]byte(v.out), v.msg)
		}
	})

	t.Run("VirusLikeName", func(t *testing.T) {
		data := []struct {
			msg       string
			name      string
			likeVirus bool
		}{
			{"name1", "Aspilota vector Belokobylskij, 2007", true},
			{"name2", "Ceylonesmus vector Chamberlin, 1941", true},
			{"name3", "Cryptops (Cryptops) vector Chamberlin, 1939", true},
			{"name4", "Culex vector Dyar & Knab, 1906", true},
			{"name5", "Dasyproctus cevirus Leclercq, 1963", true},
			{"name6", "Desmoxytes vector (Chamberlin, 1941)", true},
			{"name7", "Dicathais vector Thornley, 1952", true},
			{"name8", "Euragallia prion Kramer, 1976", true},
			{"name9", "Exochus virus Gauld & Sithole, 2002", true},
			{"name10", "Hilara vector Miller, 1923", true},
			{"name11", "Microgoneplax prion Castro, 2007", true},
			{"name12", "Neoaemula vector Mackinnon, Hiller, Long & Marshall, 2008", true},
			{"name13", "Ophion virus Gauld & Mitchell, 1981", true},
			{"name14", "Psenulus trevirus Leclercq, 1961", true},
			{"name15", "Tidabius vector Chamberlin, 1931", true},
			{"name16", "Ceylonesmus prion", false},
			{"name17", "Homo sapiens coronavirus", false},
		}
		for _, v := range data {
			assert.Equal(t, ppr.VirusLikeName(v.name), v.likeVirus, v.msg)
		}
	})

	t.Run("IsVirus", func(t *testing.T) {
		data := []struct {
			msg     string
			name    string
			isVirus bool
		}{
			{"No match", "Homo sapiens", false},
			{"Match word", "Arv1virus ", true},
			{"Match word", "Turtle herpesviruses", true},
			{"Match word", "Cre expression vector", true},
			{"Match word", "Abutilon mosaic vir. ICTV", true},
			{"Match word", "Aeromonas phage 65", true},
			{"Match word", "Apple scar skin viroid", true},
			{"Match word", "Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease", true},
			{"Match word", "Phi h-like viruses", true},
			{"Match word", "Viroids", true},
			{"Match word", "Human rhinovirus A11", true},
			{"Match word", "Gossypium mustilinum symptomless alphasatellite", true},
			{"Match word", "Bemisia betasatellite LW-2014", true},
			{"Match word", "Intracisternal A-particles", true},
			{"Match word", "Uranotaenia sapphirina NPV", true},
			{"Match word", "Spodoptera frugiperda MNPV", true},
			{"Match word", "Mamestra configurata NPV-A", true},
			{"Match word", "Bacteriophage PH75", true},
		}
		for _, v := range data {
			res := ppr.IsVirus([]byte(v.name))
			assert.Equal(t, res, v.isVirus, v.msg)
		}
	})

	t.Run("NoParse", func(t *testing.T) {
		data := []struct {
			msg    string
			name   string
			parsed bool
		}{
			{"No match", "Homo sapiens", false},
			{"No word at the start", "Not Homo sapiens", true},
			{"Noword at the start", "Nothomo sapiens", false},
			{"Not word at the start", "Not Homo sapiens", true},
			{"None word at the start", "None Homo sapiens", true},
			{"Unidentified at the start", "Unidentified species", true},
			{"Incertae sedis1", "incertae sedis", true},
			{"Incertae sedis2", "Incertae Sedis", true},
			{"Incertae sedis3", "Something incertae sedis", true},
			{"Incertae sedis4", "Homo sapiens inc.sed.", true},
			{"Incertae sedis5", "Incertae sedis", true},
			{"Phytoplasma in the middle", "Homo sapiensphytoplasmaoid", false},
			{"Phytoplasma in the end", "Homo sapiensphytoplasma Linn", true},
			{"Phytoplasma in the end", "Homo sapiensphytoplasma Linn", true},
			{"Plasmid1", "E. coli plasmids", true},
			{"Plasmid2", "E. coli plasmidia", false},
			{"Plasmid3", "E. coli plasmid", true},
			{"RNA1", "E. coli RNA", true},
			{"RNA2", "E. coli 32RNA", true},
			{"RNA3", "KURNAKOV", false},
			{"RNA4", "E. coli mRNA", true},
		}
		for _, v := range data {
			res := ppr.NoParse([]byte(v.name))
			assert.Equal(t, res, v.parsed, v.msg)
		}
	})

	t.Run("Annotations", func(t *testing.T) {
		data := []struct {
			msg  string
			in   string
			out  string
			tail string
		}{

			{"No tail", "Homo sapiens", "Homo sapiens", ""},
			{"No tail", "Homo sapiens S. S.", "Homo sapiens S. S.", ""},
			{"No tail", "Homo sapiens s. s.", "Homo sapiens", " s. s."},
			{"No tail", "Homo sapiens sensu Linn.", "Homo sapiens", " sensu Linn."},
			{"No tail", "Homo sapiens nomen nudum", "Homo sapiens", " nomen nudum"},
		}
		for _, v := range data {
			bs := []byte(v.in)
			i := ppr.Annotation(bs)
			assert.Equal(t, string(bs[0:i]), v.out, v.msg)
			assert.Equal(t, string(bs[i:]), v.tail, v.msg)
		}
	})

	t.Run("UnderscoreToSpace", func(t *testing.T) {
		data := []struct {
			msg     string
			in      string
			out     string
			changed bool
		}{
			{"no nothing", "Hello", "Hello", false},
			{"has spaces", "Hello_you !", "Hello_you !", false},
			{"has spaces", "Hello_you\t!", "Hello_you\t!", false},
			{"has only underscores", "Hello_you_!_", "Hello you ! ", true},
		}
		for _, v := range data {
			bs := []byte(v.in)
			changed2, _ := ppr.UnderscoreToSpace(bs)
			assert.Equal(t, string(bs), v.out, v.msg)
			assert.Equal(t, changed2, v.changed)
		}
	})

	t.Run("does not remove spaces", func(t *testing.T) {
		name := "    Asplenium       × inexpectatum(E. L. Braun ex Friesner      )Morton"
		res := ppr.Preprocess([]byte(name))
		assert.Equal(t, string(res.Body), name)
	})
}
