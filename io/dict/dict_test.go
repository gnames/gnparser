package dict_test

import (
	"testing"

	"github.com/gnames/gnparser/io/dict"
	"github.com/stretchr/testify/assert"
)

func TestLoadDictionary(t *testing.T) {
	d := dict.LoadDictionary()
	t.Run("loads bacterial dictionary", func(t *testing.T) {
		assert.Greater(t, len(d.Bacteria), 100)
	})
	t.Run("finds non homopypic genus", func(t *testing.T) {
		hom, ok := d.Bacteria["Sphingomonas"]
		assert.True(t, ok)
		assert.False(t, hom)
	})
	t.Run("finds homotypic genus", func(t *testing.T) {
		hom, ok := d.Bacteria["Arizona"]
		assert.True(t, ok)
		assert.True(t, hom)
	})
	t.Run("does not find non-bacterial genus", func(t *testing.T) {
		hom, ok := d.Bacteria["Homo"]
		assert.False(t, ok)
		assert.False(t, hom)
	})
	t.Run("does not find not ICN author", func(t *testing.T) {
		_, ok := d.AuthorICN["Arizona"]
		assert.False(t, ok)
	})
	t.Run("finds ICN author", func(t *testing.T) {
		_, ok := d.AuthorICN["Abramov"]
		assert.True(t, ok)
	})
}
