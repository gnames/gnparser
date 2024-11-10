package nomcode_test

import (
	"testing"

	"github.com/gnames/gnparser/ent/nomcode"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, inp string
		out      nomcode.Code
	}{
		{"bad", "something", nomcode.Unknown},
		{"zoo1", "zoo", nomcode.Zoological},
		{"zoo2", "Zoological", nomcode.Zoological},
		{"zoo3", "ICZN", nomcode.Zoological},
		{"bot1", "bot", nomcode.Botanical},
		{"bot2", "botanical", nomcode.Botanical},
		{"bot2", "icn", nomcode.Botanical},
		{"cult1", "CULT", nomcode.Cultivar},
		{"cult2", "CultiVar", nomcode.Cultivar},
		{"cult3", "icncp", nomcode.Cultivar},
		{"bact1", "bact", nomcode.Bacterial},
		{"bact2", "bacterial", nomcode.Bacterial},
		{"bact3", "ICNP", nomcode.Bacterial},
	}

	for _, v := range tests {
		res := nomcode.New(v.inp)
		assert.Equal(v.out, res, v.msg)
	}
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, out string
		inp      nomcode.Code
	}{
		{"zoo", "ICZN", nomcode.Zoological},
		{"bot", "ICN", nomcode.Botanical},
		{"bact", "ICNP", nomcode.Bacterial},
		{"cult", "ICNCP", nomcode.Cultivar},
	}

	for _, v := range tests {
		res := v.inp.String()
		assert.Equal(v.out, res, v.msg)
	}
}
