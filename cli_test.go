package gnparser_test

import (
	"strings"
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
)

func TestTMP(t *testing.T) {
	assert.True(t, true)
}

// Run make install before these tests to get meaningful
// results.

func TestVersion(t *testing.T) {
	c := testcli.Command("gnparser", "-V")
	c.Run()
	assert.True(t, c.Success())
	assert.Contains(t, c.Stdout(), "version:")

	c = testcli.Command("gnparser", "-V", "-f", "simple",
		"-j", "200", "-p", "8000")
	c.Run()
	assert.True(t, c.Success())
	assert.Contains(t, c.Stdout(), "version:")
}

func TestFormat(t *testing.T) {
	t.Run("runs csv format", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", "csv")
		c.Run()
		assert.True(t, c.Success())
		assert.Contains(t, c.Stdout(), ",Homo sapiens,2")
	})

	t.Run("ignores parsing with --version", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", "simple", "--version")
		c.Run()
		assert.True(t, c.Success())
		assert.NotContains(t, c.Stdout(), ",Homo sapiens,")
		assert.Contains(t, c.Stdout(), "version:")
	})

	t.Run("sets format to default if -f value is unknown", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", ":)")
		c.Run()
		assert.True(t, c.Success())
		assert.Contains(t, c.Stdout(), `Id,Verbatim,Cardinality,`)
	})
}

func TestStdin(t *testing.T) {
	t.Run("takes data from Stdin", func(t *testing.T) {
		c := testcli.Command("gnparser", "-f", "simple")
		c.SetStdin(strings.NewReader("Homo sapiens"))
		c.Run()
		assert.True(t, c.Success())
		assert.Contains(t, c.Stdout(), ",Homo sapiens,")
	})

	t.Run("takes multiple names from Stdin", func(t *testing.T) {
		c := testcli.Command("gnparser", "-f", "simple")
		c.SetStdin(strings.NewReader("Plantago\nBubo L.\n"))
		c.Run()
		assert.True(t, c.Success())
		assert.Contains(t, c.Stdout(), ",Plantago,")
		assert.Contains(t, c.Stdout(), ",Bubo,")
	})
}
