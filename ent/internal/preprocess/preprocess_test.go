package preprocess

import (
	"strings"
	"testing"

	"github.com/gnames/gnparser/ent/internal/preparser"
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
			assert.Equal(t, v.notags, StripTags(v.tags), v.msg)
		}
	})
	t.Run("does not return nil", func(t *testing.T) {
		assert.NotNil(t, StripTags("<!--"))
		assert.NotNil(t, StripTags("<!--\r\n"))
	})
}

func TestPreprocess(t *testing.T) {
	t.Run("NoParseLikeName", func(t *testing.T) {
		data := []struct {
			msg            string
			name           string
			likeAnnotation bool
		}{
			{"name", "Navicula bacterium", true},
		}
		for _, v := range data {
			words := strings.Split(v.name, " ")
			assert.Equal(t, v.likeAnnotation, isException(words, NoParseException), v.msg)
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
			words := strings.Split(v.name, " ")
			assert.Equal(t, v.likeVirus, isException(words, VirusException), v.msg)
		}
	})

	t.Run("IsVirus", func(t *testing.T) {
		data := []struct {
			msg     string
			name    string
			isVirus bool
		}{
			{"No match", "Homo sapiens", false},
			{"Match word 1", "Arv1virus ", true},
			{"Match word 2", "Turtle herpesviruses", true},
			{"Match word 3", "Cre expression vector", true},
			{"Match word 4", "Abutilon mosaic vir. ICTV", true},
			{"Match word 5", "Aeromonas phage 65", true},
			{"Match word 6", "Apple scar skin viroid", true},
			{
				"Match word 7",
				"Agents of Spongiform Encephalopathies CWD prion Chronic wasting disease",
				true,
			},
			{"Match word 8", "Phi h-like viruses", true},
			{"Match word 9", "Viroids", true},
			{"Match word 10", "Human rhinovirus A11", true},
			{"Match word 11", "Gossypium mustilinum symptomless alphasatellite", true},
			{"Match word 12", "Bemisia betasatellite LW-2014", true},
			{"Match word 13", "Intracisternal A-particles", true},
			{"Match word 14", "Uranotaenia sapphirina NPV", true},
			{"Match word 15", "Spodoptera frugiperda MNPV", true},
			{"Match word 16", "Mamestra configurata NPV-A", true},
			{"Match word 17", "Bacteriophage PH75", true},
		}
		for _, v := range data {
			res := IsVirus([]byte(v.name))
			assert.Equal(t, v.isVirus, res, v.msg)
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
			res := NoParse([]byte(v.name))
			assert.Equal(t, v.parsed, res, v.msg)
		}
	})

	t.Run("Annotations", func(t *testing.T) {
		tests := []struct {
			msg  string
			in   string
			out  string
			tail string
		}{

			{"No tail", "Homo sapiens", "Homo sapiens", ""},
			{"S. S.", "Homo sapiens S. S.", "Homo sapiens S. S.", ""},
			{"s. s.", "Homo sapiens s. s.", "Homo sapiens", " s. s."},
			{"sensu", "Homo sapiens sensu Linn.", "Homo sapiens", " sensu Linn."},
			{"nomen", "Homo sapiens nomen nudum", "Homo sapiens", " nomen nudum"},
		}
		ppr := preparser.New()

		for _, v := range tests {
			bs := []byte(v.in)
			i := procAnnot(ppr, bs)
			assert.Equal(t, v.out, string(bs[0:i]), v.msg)
			assert.Equal(t, v.tail, string(bs[i:]), v.msg)
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
			changed2, _ := UnderscoreToSpace(bs)
			assert.Equal(t, v.out, string(bs), v.msg)
			assert.Equal(t, v.changed, changed2)
		}
	})

	t.Run("does not remove spaces", func(t *testing.T) {
		name := "    Asplenium       × inexpectatum(E. L. Braun ex Friesner      )Morton"
		ppr := preparser.New()
		res := Preprocess(ppr, []byte(name))
		assert.Equal(t, name, string(res.Body))
	})
}
