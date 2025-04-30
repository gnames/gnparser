package parser_test

import (
	"fmt"
	"testing"

	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser/ent/parser"
	"github.com/stretchr/testify/assert"
)

// TTestPreNParse tests PreprocessAndParse method
func TestPreNParse(t *testing.T) {
	p := parser.New()
	testData := []struct {
		name, can string
	}{
		{"Pardosa moesta L.", "Pardosa moesta"},
		{"something", ""},
	}
	for _, v := range testData {
		sn := p.PreprocessAndParse(
			v.name, "test_version", nomcode.Unknown, true, false, false,
		)
		parsed := sn.ToOutput(false, false)
		can := parsed.Canonical
		msg := v.name
		if v.can == "" {
			assert.Nil(t, can, msg)
			continue
		}
		assert.Equal(t, v.can, can.Simple, msg)
	}
}

// TestToOutput tests ToOutput method of ScientificNameNode
func TestToOutput(t *testing.T) {
	p := parser.New()
	testData := []struct {
		name, can, au string
		det, parsed   bool
	}{
		{"Pardosa moesta L.", "Pardosa moesta", "L.", false, true},
		{
			"Bacillus subtilis (Ehrenberg, 1835) Cohn, 1872",
			"Bacillus subtilis", "(Ehrenberg 1835) Cohn 1872",
			false, true,
		},
		{
			"Bacillus subtilis (Ehrenberg, 1835) Cohn, 1872 sec. Miller",
			"Bacillus subtilis", "(Ehrenberg 1835) Cohn 1872",
			false, true,
		},
		{
			"Aconitum napellus var. formosum (Rchb.) W. D. J. Koch (nom. ambig.)",
			"Aconitum napellus formosum", "(Rchb.) W. D. J. Koch",
			true, true,
		},
		{"something", "", "", false, false},
	}
	for _, v := range testData {
		sn := p.PreprocessAndParse(
			v.name, "test_version", nomcode.Unknown, true, false, false,
		)
		out := sn.ToOutput(v.det, false)
		msg := v.name
		if !out.Parsed {
			assert.Nil(t, out.Canonical, msg)
			continue
		}
		assert.Equal(t, v.can, out.Canonical.Simple, msg)
		assert.Equal(t, v.au, out.Authorship.Normalized, msg)
	}
}

// TestSpecGroupOption checks if stem is cut when WithSpeciesGroupCut is true.
func TestSpecGroupOption(t *testing.T) {
	assert := assert.New(t)

	p := parser.New()
	testData := []struct {
		name, stemmed string
		spGrp         bool
	}{
		{"Aus alba alba", "Aus alb alb", false},
		{"Aus alba alba", "Aus alb", true},
		{"Aus alba albus", "Aus alb alb", true},
		{
			"Bacillus subtilis subtilis (Ehrenberg, 1835) Cohn, 1872",
			"Bacillus subtil subtil", false,
		},
		{
			"Bacillus subtilis subtilis (Ehrenberg, 1835) Cohn, 1872",
			"Bacillus subtil", true,
		},
		{
			"Bacillus subtila subtilis (Ehrenberg, 1835) Cohn, 1872",
			"Bacillus subtil subtil", true,
		},
	}
	for _, v := range testData {
		sn := p.PreprocessAndParse(
			v.name, "test_version",
			nomcode.Unknown,
			true, false, false,
		)
		out := sn.ToOutput(false, v.spGrp)
		msg := v.name
		fmt.Println(out.Canonical.Simple)
		assert.Equal(v.stemmed, out.Canonical.Stemmed, msg)
	}
}
