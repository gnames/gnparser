package icvcn_test

import (
	"testing"

	"github.com/gnames/gnparser/ent/icvcn"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, input, output string
		ok                 bool
	}{
		{"realm Adnaviria", "Adnaviria", "Adnaviria", true},
		{"realm Riboviria", "Riboviria", "Riboviria", true},
		{"genus Flavivirus", "Flavivirus", "Flavivirus", true},
		{"genus Triavirus", "Triavirus", "Triavirus", true},
		{"family Flaviviridae", "Flaviviridae", "Flaviviridae", true},
		{"species Triavirus", "Triavirus phi2958PVL", "Triavirus phi2958PVL", true},
		{"species Flavivirus", "Flavivirus encephalitidis", "Flavivirus encephalitidis", true},
	}

	for _, v := range tests {
		p := &icvcn.Parser{Buffer: v.input}
		p.Init()
		err := p.Parse()
		if err != nil {
			t.Errorf("%s failed: %v", v.msg, err)
		} else {
			t.Logf("%s: OK", v.msg)
		}
		assert.Equal(v.ok, (err == nil), v.msg)
	}
}
