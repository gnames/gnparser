package preparser_test

import (
	"testing"

	"github.com/gnames/gnparser/ent/internal/preparser"
	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	debug := true
	q := "The annulignatha group"
	assert := assert.New(t)
	ppr := preparser.New()
	if debug {
		err := ppr.Debug(q)
		assert.Nil(err)
	}
}

func TestPreParsed(t *testing.T) {
	tests := []struct {
		msg, str, tail string
	}{
		// Last 'junk' words/ annotations
		{"var", "Musca domeſtica Linnaeus 1758 var?  ", " var?  "},
		{"ined", "  Oxalis_barrelieri ined.?", " ined.?"},
		{"ssp.", "Peperomia non-alata Trel. ssp.", " ssp."},
		{"subsp.", "Sanogasta x-signata (Keyserling,1891) subsp.",
			" subsp."},
		{"subgen", "Sanogasta x-signata (Keyserling,1891) subgen?  ",
			" subgen?  "},
		{"sensu", "Pseudomonas methanica (Söhngen 1906) sensu. Dworkin and Foster 1956",
			" sensu. Dworkin and Foster 1956"},
		{"new", "Hegeter (Hegeter) intercedens Lindberg H 1950 new", " new"},
		{"non", "Anthoscopus Cabanis [1851?] non", " non"},
		{"nec", "Hegeter (Hegeter) intercedens Lindberg H 1950 nec", " nec"},
		{"hybrid", "  Arthopyrenia hyalospora x hybrid?", " hybrid?"},
		{"von$", "Nautilus asterizans von", " von"},

		// Pro Parte
		{"Pro Parte", "Abarema clypearia (Jack) Kosterm., Pro Parte",
			", Pro Parte"},
		{"nomen", "Akeratidae Nomen Nudum", " Nomen Nudum"},
		{"nom.", "Akeratidae nom. nudum", " nom. nudum"},
		{"nom illeg", "Abutilon avicennae Gaertn., nom. illeg.", ", nom. illeg."},
		{"comb", "Arthopyrenia hyalospora (Nyl.) R.C. Harris comb. nov.",
			" comb. nov."},
		{"p. p.", "Abarema clypearia (Jack) Kosterm., p. p.", ", p. p."},
		{"P. P.", "Abarema clypearia (Jack) Kosterm., P. P.", ", P. P."},

		// s.s.
		{", s. s.", "Bubo bubo, s. s. nov spec something",
			", s. s. nov spec something"},
		{"s.s.", "Bubo bubo s.s. nov spec something",
			" s.s. nov spec something"},
		{"s.l.", "Bubo bubo s.l. something",
			" s.l. something"},
		{"s. lat.", "Bubo bubo s. lat. something",
			" s. lat. something"},
		{"s. str.", "Bubo bubo s. str. something",
			" s. str. something"},
		{"no break space", " Canadensis Erxleben, 1777 s.str.", " s.str."},

		// Stop words
		{"env", "Ge Nicéville 1895 Environmental sample",
			" Environmental sample"},
		{"env samples", "Candidatus Anammoxoglobus environmental samples",
			" environmental samples"},
		{"enrichment", "Crenarchaeote enrichment culture clone OREC-B1022",
			" enrichment culture clone OREC-B1022"},
		{"samples", "Candidatus Anammoxoglobus samples",
			" samples"},

		{"sec", "Ataladoris Iredale & O'Donoghue 1923 sec Eschmeyer",
			" sec Eschmeyer"},
		{"sec.", "Ataladoris Iredale & O'Donoghue 1923 sec. Eschmeyer",
			" sec. Eschmeyer"},
		{"sp compl", "Acarospora cratericola cratericola Shenk 1974 species complex",
			" species complex"},
		{"utf8", "× Dialaeliopsis hort.", " hort."},
	}

	assert := assert.New(t)
	ppr := preparser.New()

	for _, v := range tests {
		idx := ppr.TailIndex(v.str)
		assert.True(idx >= 0, v.msg)
		assert.Equal(v.tail, string([]byte(v.str)[idx:]), v.msg)
	}
}

func TestNotPreParsed(t *testing.T) {
	tests := []struct {
		msg, str string
	}{
		{"no tail1", "Lachenalia tricolor var. nelsonii (anon.) Baker"},
		{"S. S.", "Bubo bubo, S. S. something"},
		{"dagger", "Heteralocha acutirostris (Gould, 1837) Huia N E†"},
		{"spaces", "Heteralocha acutirostris (Gould, 1837) Huia N E   "},
		{"comma", "Abantiadinus pusillus Broun, T. , 1914"},
		{"last comma", "Acalles foveopunctatus Fiedler,"},
		{"space comma", "Calamagrostis neglecta G.Gaertn. ,B.Mey. & Scherb."},
		{"all tail", "Non splenectomized mulatta"},
		{"several commas", "Naupliicola cystifingens Michajlow, ,1968"},
		{"spp", "Crataegus curvisepala nvar. naviculiformis T. Petauer Alaria spp."},
	}

	assert := assert.New(t)
	ppr := preparser.New()

	for _, v := range tests {
		idx := ppr.TailIndex(v.str)
		assert.Equal(-1, idx)
	}
}
