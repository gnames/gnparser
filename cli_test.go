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

func TestFlattenOutput(t *testing.T) {
	t.Run("flatten with JSON compact format", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens Linnaeus, 1758", "-f", "compact", "-F")
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		// Flattened JSON should not have nested canonical
		assert.NotContains(t, output, `"canonical":`)
		// Should have flat fields
		assert.Contains(t, output, `"canonicalSimple"`)
		assert.Contains(t, output, `"canonicalFull"`)
		assert.Contains(t, output, `"authorship"`)
	})

	t.Run("flatten with JSON pretty format", func(t *testing.T) {
		c := testcli.Command("gnparser", "Bubo bubo", "-f", "pretty", "-F")
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		assert.NotContains(t, output, `"canonical":`)
		assert.Contains(t, output, `"canonicalSimple"`)
	})

	t.Run("without flatten flag uses nested JSON", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", "compact")
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		// Non-flattened JSON should have nested canonical
		assert.Contains(t, output, `"canonical":`)
		assert.NotContains(t, output, `"canonicalSimple"`)
	})

	t.Run("CSV without details is simple", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", "csv")
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		// CSV without -d should have simple header (10 fields)
		lines := strings.Split(output, "\n")
		header := lines[0]
		assert.Contains(t, header, "Id,Verbatim,Cardinality")
		assert.Contains(t, header, "NomCodeSetting")
		// Should NOT include extended fields
		assert.NotContains(t, header, "Parsed,")
		assert.NotContains(t, header, ",Genus,")
		assert.NotContains(t, header, "CultivarEpithet")
	})

	t.Run("CSV with details includes all fields", func(t *testing.T) {
		c := testcli.Command("gnparser", "Homo sapiens", "-f", "csv", "-d")
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		// CSV with details should include all extended fields
		assert.Contains(t, output, ",Genus,")
		assert.Contains(t, output, ",Species,")
		assert.Contains(t, output, ",Infraspecies")
		assert.Contains(t, output, "Parsed,")
		assert.Contains(t, output, ",Authors,")
	})

	t.Run("flatten from stdin", func(t *testing.T) {
		c := testcli.Command("gnparser", "-f", "compact", "-F")
		c.SetStdin(strings.NewReader("Homo sapiens"))
		c.Run()
		assert.True(t, c.Success())
		output := c.Stdout()
		assert.NotContains(t, output, `"canonical":`)
		assert.Contains(t, output, `"canonicalSimple"`)
	})
}
