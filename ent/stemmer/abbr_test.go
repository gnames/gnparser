package stemmer_test

import (
	"strings"
	"testing"

	"github.com/gnames/gnparser/ent/stemmer"
	"github.com/stretchr/testify/assert"
)

// TestAbbrKey keeps keys stable. If a change in the stemmer breaks
// this test, databases built with AbbrKey (gndb) and clients that
// query them (gnames) have to be updated together.
func TestAbbrKey(t *testing.T) {
	tests := []struct {
		msg     string
		genus   string
		epithet string
		want    string
	}{
		{"full genus", "Caenorhabditis", "elegans", "c-elegans"},
		{"abbreviated genus", "C.", "elegans", "c-elegans"},
		{"several genus letters", "Cae.", "elegans", "c-elegans"},
		{"capitalized epithet", "C.", "Elegans", "c-elegans"},
		{"extra spaces", " C. ", " elegans ", "c-elegans"},
		{"feminine ending", "Carex", "alba", "c-alb"},
		{"masculine ending", "C.", "albus", "c-alb"},
		{"neuter ending", "C.", "album", "c-alb"},
		{"species epithet", "P.", "tigris", "p-tigr"},
		{"infraspecific epithet", "P.", "altaica", "p-altaic"},
		{"v becomes u", "S.", "vulgaris", "s-uulgar"},
		{"short epithet is not stemmed", "Aus", "jo", "a-jo"},
		{"diaeresis", "Aus", "coëlestis", "a-coelest"},
		{"empty genus", "", "elegans", ""},
		{"empty epithet", "C.", "", ""},
		{"empty stem", "Bisetocreagris", "que", ""},
	}

	for _, v := range tests {
		t.Run(v.msg, func(t *testing.T) {
			assert.Equal(t, v.want, stemmer.AbbrKey(v.genus, v.epithet))
		})
	}
}

// TestAbbrKey_MatchesStemCanonical checks that a key built from text
// matches a key built from the stemmed canonical form.
func TestAbbrKey_MatchesStemCanonical(t *testing.T) {
	canonicals := []string{
		"Caenorhabditis elegans",
		"Panthera tigris altaica",
		"Bison bison bison",
		"Aus jo",
		"Aus coëlestis",
	}

	for _, can := range canonicals {
		words := strings.Fields(can)
		stems := strings.Fields(stemmer.StemCanonical(can))
		for i := 1; i < len(words); i++ {
			want := strings.ToLower(words[0][:1]) + "-" + stems[i]
			got := stemmer.AbbrKey(words[0], words[i])
			assert.Equal(t, want, got, can)
		}
	}
}
