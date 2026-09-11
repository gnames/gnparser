package parser_test

import (
	"testing"

	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/parser"
	"github.com/stretchr/testify/assert"
)

type tailAnnot struct {
	kind, verbatim, normalized, author string
}

func tailAnnots(ta *parsed.TailAnnotations) []tailAnnot {
	if ta == nil {
		return nil
	}
	var res []tailAnnot
	for _, v := range ta.Sensu {
		res = append(res, tailAnnot{"sensu", v.Verbatim, v.Normalized, v.Author})
	}
	for _, v := range ta.Status {
		res = append(res, tailAnnot{"status", v.Verbatim, v.Normalized, ""})
	}
	for _, v := range ta.Publication {
		res = append(res,
			tailAnnot{"publication", v.Verbatim, v.Normalized, v.Author})
	}
	return res
}

// parseTail runs tail parsing on a parsed name-string with the given tail.
func parseTail(p parser.Parser, tail string, withAuthorship bool) parsed.Parsed {
	res := parsed.Parsed{
		Parsed:       true,
		ParseQuality: 4,
		QualityWarnings: []parsed.QualityWarning{
			parsed.TailWarn.NewQualityWarning(),
		},
		Tail: tail,
	}
	if withAuthorship {
		res.Authorship = &parsed.Authorship{Verbatim: "Smith", Normalized: "Smith"}
	}
	res = p.ParseTail(res)
	return res
}

func TestParseTailVariants(t *testing.T) {
	p := parser.New()
	tests := []struct {
		tail  string
		annot tailAnnot
	}{
		{" sensu lato", tailAnnot{"sensu", "sensu lato", "sensu lato", ""}},
		{" Sensu lato", tailAnnot{"sensu", "Sensu lato", "sensu lato", ""}},
		{" s.l.", tailAnnot{"sensu", "s.l.", "sensu lato", ""}},
		{" s. l.", tailAnnot{"sensu", "s. l.", "sensu lato", ""}},
		{" s.lat.", tailAnnot{"sensu", "s.lat.", "sensu lato", ""}},
		{" s. lat.", tailAnnot{"sensu", "s. lat.", "sensu lato", ""}},
		{" sensu stricto", tailAnnot{"sensu", "sensu stricto", "sensu stricto", ""}},
		{" s.s.", tailAnnot{"sensu", "s.s.", "sensu stricto", ""}},
		{" s. s.", tailAnnot{"sensu", "s. s.", "sensu stricto", ""}},
		{" s.str.", tailAnnot{"sensu", "s.str.", "sensu stricto", ""}},
		{" s. str.", tailAnnot{"sensu", "s. str.", "sensu stricto", ""}},
		{" sensu Jones", tailAnnot{"sensu", "sensu", "sensu", "Jones"}},
		{" sensu. Jones", tailAnnot{"sensu", "sensu.", "sensu", "Jones"}},
		{" sensu Lavallée, 1880", tailAnnot{"sensu", "sensu", "sensu", "Lavallée, 1880"}},
		{" auct.", tailAnnot{"sensu", "auct.", "auct.", ""}},
		{" Auct.", tailAnnot{"sensu", "Auct.", "auct.", ""}},
		{" auct", tailAnnot{"sensu", "auct", "auct.", ""}},
		{" sensu auct.", tailAnnot{"sensu", "sensu auct.", "sensu auct.", ""}},
		{" auct. non Jones", tailAnnot{"sensu", "auct. non", "auct. non", "Jones"}},
		{" auct., non Jones", tailAnnot{"sensu", "auct., non", "auct. non", "Jones"}},
		{" Auct non L.", tailAnnot{"sensu", "Auct non", "auct. non", "L."}},
		{
			" sensu auct., non (Smith) Jones",
			tailAnnot{"sensu", "sensu auct., non", "sensu auct. non", "(Smith) Jones"},
		},
		{" pro parte", tailAnnot{"sensu", "pro parte", "pro parte", ""}},
		{" p.p.", tailAnnot{"sensu", "p.p.", "pro parte", ""}},
		{" p. p.", tailAnnot{"sensu", "p. p.", "pro parte", ""}},
		{" P. P.", tailAnnot{"sensu", "P. P.", "pro parte", ""}},
		{" pro p.", tailAnnot{"sensu", "pro p.", "pro parte", ""}},

		{" nom. dub.", tailAnnot{"status", "nom. dub.", "nomen dubium", ""}},
		{" nomen dubium", tailAnnot{"status", "nomen dubium", "nomen dubium", ""}},
		{" nom. nud.", tailAnnot{"status", "nom. nud.", "nomen nudum", ""}},
		{" nom.nud.", tailAnnot{"status", "nom.nud.", "nomen nudum", ""}},
		{" nomen nudum", tailAnnot{"status", "nomen nudum", "nomen nudum", ""}},
		{" nom. illeg.", tailAnnot{"status", "nom. illeg.", "nomen illegitimum", ""}},
		{
			" nomen illegitimum",
			tailAnnot{"status", "nomen illegitimum", "nomen illegitimum", ""},
		},
		{" nom. cons.", tailAnnot{"status", "nom. cons.", "nomen conservandum", ""}},
		{
			" nomen conservandum",
			tailAnnot{"status", "nomen conservandum", "nomen conservandum", ""},
		},
		{" nom. rej.", tailAnnot{"status", "nom. rej.", "nomen rejiciendum", ""}},
		{
			" nomen rejiciendum",
			tailAnnot{"status", "nomen rejiciendum", "nomen rejiciendum", ""},
		},
		{" nom. prov.", tailAnnot{"status", "nom. prov.", "nomen provisorium", ""}},
		{
			" nomen provisorium",
			tailAnnot{"status", "nomen provisorium", "nomen provisorium", ""},
		},
		{" stat. nov.", tailAnnot{"status", "stat. nov.", "status novus", ""}},
		{" status novus", tailAnnot{"status", "status novus", "status novus", ""}},
		{" comb. nov.", tailAnnot{"status", "comb. nov.", "combinatio nova", ""}},
		{
			" combinatio nova",
			tailAnnot{"status", "combinatio nova", "combinatio nova", ""},
		},
		{" stat. rev.", tailAnnot{"status", "stat. rev.", "status restitutus", ""}},
		{
			" status restitutus",
			tailAnnot{"status", "status restitutus", "status restitutus", ""},
		},

		{" ined.", tailAnnot{"publication", "ined.", "ineditus", ""}},
		{" hort.", tailAnnot{"publication", "hort.", "hortorum", ""}},
		{" Hort.", tailAnnot{"publication", "Hort.", "hortorum", ""}},
		{" fide Jones, 1900", tailAnnot{"publication", "fide", "fide", "Jones, 1900"}},
		{
			" fide Müller & Söhngen, 1906",
			tailAnnot{"publication", "fide", "fide", "Müller & Söhngen, 1906"},
		},
		{" em. Jones", tailAnnot{"publication", "em.", "emend.", "Jones"}},
		{" emend. Jones", tailAnnot{"publication", "emend.", "emend.", "Jones"}},
		{" emend Jones", tailAnnot{"publication", "emend", "emend.", "Jones"}},
		{" ex Jones", tailAnnot{"publication", "ex", "ex", "Jones"}},

		{" (s.str.)", tailAnnot{"sensu", "s.str.", "sensu stricto", ""}},
		{" [nom. nud.]", tailAnnot{"status", "nom. nud.", "nomen nudum", ""}},
		{" ( nom. nud. )", tailAnnot{"status", "nom. nud.", "nomen nudum", ""}},
		{", nom. illeg.", tailAnnot{"status", "nom. illeg.", "nomen illegitimum", ""}},
	}

	for _, v := range tests {
		res := parseTail(p, v.tail, true)
		assert.Equal(t, []tailAnnot{v.annot}, tailAnnots(res.TailAnnotations), v.tail)
		assert.Equal(t, "", res.Tail, v.tail)
		assert.Equal(t, 1, res.ParseQuality, v.tail)
		if res.TailAnnotations != nil {
			assert.Equal(t, v.tail, res.TailAnnotations.Verbatim, v.tail)
		}
	}
}

func TestParseTailRelations(t *testing.T) {
	p := parser.New()
	tests := []struct {
		tail string
		rel  *parsed.ConceptRelation
	}{
		{
			" sensu lato",
			&parsed.ConceptRelation{Type: parsed.RelationBroader, RCC5: ">"},
		},
		{
			" s.l. Jones, 1900",
			&parsed.ConceptRelation{
				Type: parsed.RelationBroader, RCC5: ">", ReferenceAuthor: "Jones, 1900",
			},
		},
		{
			" sensu stricto",
			&parsed.ConceptRelation{Type: parsed.RelationNarrower, RCC5: "<"},
		},
		{
			" s. str. Jones",
			&parsed.ConceptRelation{
				Type: parsed.RelationNarrower, RCC5: "<", ReferenceAuthor: "Jones",
			},
		},
		{
			" sensu Jones, 1900",
			&parsed.ConceptRelation{
				Type: parsed.RelationSameAs, RCC5: "==", ReferenceAuthor: "Jones, 1900",
			},
		},
		{
			" auct. non Jones, 1900",
			&parsed.ConceptRelation{
				Type:            parsed.RelationMisapplicationOf,
				RCC5:            "|",
				ReferenceAuthor: "Jones, 1900",
			},
		},
		{
			" sensu auct. non Jones",
			&parsed.ConceptRelation{
				Type:            parsed.RelationMisapplicationOf,
				RCC5:            "|",
				ReferenceAuthor: "Jones",
			},
		},
		{" auct.", nil},
		{" sensu auct.", nil},
		{" pro parte", nil},
	}

	for _, v := range tests {
		res := parseTail(p, v.tail, true)
		ta := res.TailAnnotations
		if !assert.NotNil(t, ta, v.tail) || !assert.Len(t, ta.Sensu, 1, v.tail) {
			continue
		}
		assert.Equal(t, v.rel, ta.Sensu[0].ConceptRelation, v.tail)
		if v.rel != nil {
			assert.Equal(t, ta.Sensu[0].Author, v.rel.ReferenceAuthor, v.tail)
		}
	}
}

func TestParseTailResidual(t *testing.T) {
	p := parser.New()
	tests := []struct {
		msg, tail, residual string
		annotsNum           int
	}{
		{"junk after annotation", " sensu lato foo", " foo", 1},
		{"junk after dash", " s.lat. - Aus bus", " - Aus bus", 1},
		{"non after sensu", " sensu Smith, non Jones", ", non Jones", 1},
		{"non without author", " auct. non", " non", 1},
		{"junk before annotation", " foo s.l.", " foo s.l.", 0},
		{"sensu without author", " sensu", " sensu", 0},
		{"no boundary", " p.p.B", " p.p.B", 0},
		{"question mark", " ined.?", " ined.?", 0},
		{"unclosed bracket", " (s.str.", " (s.str.", 0},
		{"standalone non", " non Jones, 1900", " non Jones, 1900", 0},
		{"trailing separator", " s.l. ,", "", 1},
		{"semicolon", " nom. nud.; p.p.", "", 2},
		{"several", " sensu lato Smith, 1850 p.p.", "", 2},
		{"hort. ex", " hort. ex Smith", "", 2},
	}

	for _, v := range tests {
		res := parseTail(p, v.tail, true)
		assert.Equal(t, v.residual, res.Tail, v.msg)
		assert.Len(t, tailAnnots(res.TailAnnotations), v.annotsNum, v.msg)
		if v.annotsNum == 0 {
			assert.Nil(t, res.TailAnnotations, v.msg)
		}
		quality := 1
		if v.residual != "" {
			quality = 4
		}
		assert.Equal(t, quality, res.ParseQuality, v.msg)
	}
}

func TestParseTailQuality(t *testing.T) {
	p := parser.New()
	tests := []struct {
		msg, tail      string
		withAuthorship bool
		quality        int
		warnings       []parsed.Warning
	}{
		{"s.l. with authorship", " s.l.", true, 1, nil},
		{
			"s.l. without author", " s.l.", false, 3,
			[]parsed.Warning{parsed.SensuNoAuthorWarn},
		},
		{
			"p.p. without author", " p.p.", false, 3,
			[]parsed.Warning{parsed.SensuNoAuthorWarn},
		},
		{
			"s.l. p.p. without author", " s.l. p.p.", false, 3,
			[]parsed.Warning{parsed.SensuNoAuthorWarn},
		},
		{
			"s.l. author without authorship", " sensu lato Jones", false, 3,
			[]parsed.Warning{parsed.SensuAuthorAmbiguousWarn},
		},
		{"s.l. author with authorship", " sensu lato Jones", true, 1, nil},
		{
			"sensu author without authorship", " sensu Jones", false, 3,
			[]parsed.Warning{parsed.SensuAuthorAmbiguousWarn},
		},
		{"sensu author with authorship", " sensu Jones", true, 1, nil},
		{
			"two ambiguous authors", " sensu lato Jones sensu Brown", false, 3,
			[]parsed.Warning{parsed.SensuAuthorAmbiguousWarn},
		},
		{"auct. non author", " auct. non Jones", false, 1, nil},
		{"auct.", " auct.", false, 1, nil},
		{"status", " nom. nud.", false, 1, nil},
		{"s.l. with fide author", " s.l. fide Jones", false, 1, nil},
		{
			"residual", " sensu lato foo", true, 4,
			[]parsed.Warning{parsed.TailWarn},
		},
		{
			"residual without author", " sensu lato foo", false, 4,
			[]parsed.Warning{parsed.TailWarn, parsed.SensuNoAuthorWarn},
		},
	}

	for _, v := range tests {
		res := parseTail(p, v.tail, v.withAuthorship)
		assert.Equal(t, v.quality, res.ParseQuality, v.msg)
		var ws []parsed.Warning
		for _, w := range res.QualityWarnings {
			ws = append(ws, w.Warning)
		}
		assert.Equal(t, v.warnings, ws, v.msg)
	}
}

func TestParseTailKeepsWarnings(t *testing.T) {
	p := parser.New()
	res := parsed.Parsed{
		Parsed:       true,
		ParseQuality: 4,
		QualityWarnings: []parsed.QualityWarning{
			parsed.TailWarn.NewQualityWarning(),
			parsed.AuthExWarn.NewQualityWarning(),
		},
		Authorship: &parsed.Authorship{Verbatim: "Smith ex Jones"},
		Tail:       " nom. nud.",
	}
	res = p.ParseTail(res)
	assert.Equal(t, 2, res.ParseQuality)
	assert.Equal(t,
		[]parsed.QualityWarning{parsed.AuthExWarn.NewQualityWarning()},
		res.QualityWarnings,
	)
}

func TestParseTailNoChange(t *testing.T) {
	p := parser.New()
	tests := []struct {
		msg string
		res parsed.Parsed
	}{
		{"not parsed", parsed.Parsed{Tail: " s.l."}},
		{"no tail", parsed.Parsed{Parsed: true, ParseQuality: 1}},
		{
			"not recognized",
			parsed.Parsed{
				Parsed:       true,
				ParseQuality: 4,
				QualityWarnings: []parsed.QualityWarning{
					parsed.TailWarn.NewQualityWarning(),
				},
				Tail: " foo",
			},
		},
	}

	for _, v := range tests {
		res := v.res
		res = p.ParseTail(res)
		assert.Equal(t, v.res, res, v.msg)
	}
}

// TestParseTailThenName checks that tail parsing does not affect parsing
// of the next name-string by the same parser.
func TestParseTailThenName(t *testing.T) {
	p := parser.New()
	res := parseTail(p, " sensu lato Smith, 1850", false)
	assert.NotNil(t, res.TailAnnotations)

	sn := p.PreprocessAndParse(
		"Pardosa moesta Banks, 1892", "test_version", nomcode.Unknown,
		true, false, false, false,
	)
	out := sn.ToOutput(true, false)
	assert.Equal(t, "Pardosa moesta", out.Canonical.Simple)
	assert.Equal(t, "Banks, 1892", out.Authorship.Verbatim)
	assert.Equal(t, 1, out.ParseQuality)
	assert.Equal(t, "", out.Tail)
}
