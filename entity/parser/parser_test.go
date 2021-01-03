package parser_test

import (
	"testing"

	"github.com/gnames/gnparser/entity/parser"
	"github.com/stretchr/testify/assert"
)

// TTestPreNParse tests PreprocessAndParse method
func TestPreNParse(t *testing.T) {
	p := &parser.Engine{Buffer: ""}
	p.Init()
	testData := []struct {
		name, can string
	}{
		{"Pardosa moesta L.", "Pardosa moesta"},
		{"something", ""},
	}
	for _, v := range testData {
		sn := p.PreprocessAndParse(v.name, "test_version", true)
		can := sn.Canonical()
		msg := v.name
		if v.can == "" {
			assert.Nil(t, can)
			continue
		}
		assert.Equal(t, can.Simple, v.can, msg)
	}
}

// TTestToOutput tests ToOutput method of ScientificNameNode
func TestToOutput(t *testing.T) {
	p := &parser.Engine{Buffer: ""}
	p.Init()
	testData := []struct {
		name, can, auNorm string
		det, parsed       bool
	}{
		{"Pardosa moesta L.", "Pardosa moesta", "L.", false, true},
		{"something", "", "", false, false},
	}
	for _, v := range testData {
		sn := p.PreprocessAndParse(v.name, "test_version", true)
		out := sn.ToOutput(v.det)
		msg := v.name
		if !out.Parsed {
			assert.Nil(t, out.Canonical)
			continue
		}
		assert.Equal(t, out.Canonical.Simple, v.can, msg)
	}
}
