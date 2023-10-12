package stemmer_test

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gnames/gnparser/ent/stemmer"
	"github.com/stretchr/testify/assert"
)

func TestStemmer(t *testing.T) {
	stemsDict := stemData(t)
	t.Run("treats que suffix with exceptions", func(t *testing.T) {
		assert.Equal(t, "detorque", stemmer.Stem("detorque").Stem)
		assert.Equal(t, "something", stemmer.Stem("somethingque").Stem)
	})
	t.Run("removes suffixes correctly", func(t *testing.T) {
		for k, v := range stemsDict {
			assert.Equal(t, v, stemmer.Stem(k).Stem)
		}
	})

	t.Run("StemCanonical", func(t *testing.T) {
		data := []struct {
			msg string
			in  string
			out string
		}{
			{"Uninomial", "Pomatomus", "Pomatomus"},
			{"Binomial1", "Betula naturae", "Betula natur"},
			{"Binomial2", "Betula alba", "Betula alb"},
			{"Binomial3", "Leptochloöpsis virgata", "Leptochloopsis uirgat"},
			{"Trinomial", "Betula alba naturae", "Betula alb natur"},
			{"SpGroup", "Betula alba alba", "Betula alb alb"},
			{"SpGroup", "Betula alba albus", "Betula alb alb"},
			{"GraftChimeraFormula", "Crataegus + Mespilus", "Crataegus + Mespilus"},
			{"GraftChimeraFormula2", "Cytisus purpureus + Laburnum anagyroides", "Cytisus purpure + Laburnum anagyroid"},
		}
		for _, v := range data {
			assert.Equal(t, v.out, stemmer.StemCanonical(v.in), v.msg)
		}
	})
}

func stemData(t *testing.T) map[string]string {
	res := make(map[string]string)
	path := filepath.Join("..", "..", "testdata", "stems.txt")
	f, err := os.Open(path)
	assert.Nil(t, err)
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		l := strings.TrimSpace(scan.Text())
		ws := regexp.MustCompile(`\s+`).Split(l, 2)
		res[ws[0]] = ws[1]
	}

	assert.Nil(t, scan.Err())

	return res
}
