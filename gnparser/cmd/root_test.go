package cmd_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/gnames/gnparser/gnparser/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// runCLI invokes the gnparser CLI in-process. Tests no longer require the
// compiled `gnparser` binary on PATH.
func runCLI(t *testing.T, stdin string, args ...string) (string, string, error) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	var in io.Reader
	if stdin != "" {
		in = strings.NewReader(stdin)
	}
	err := cmd.ExecuteWith(args, in, &stdout, &stderr)
	return stdout.String(), stderr.String(), err
}

func TestVersion(t *testing.T) {
	out, _, err := runCLI(t, "", "-V")
	require.NoError(t, err)
	assert.Contains(t, out, "version:")

	out, _, err = runCLI(t, "", "-V", "-f", "simple", "-j", "200", "-p", "8000")
	require.NoError(t, err)
	assert.Contains(t, out, "version:")
}

func TestFormat(t *testing.T) {
	t.Run("runs csv format", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", "csv")
		require.NoError(t, err)
		assert.Contains(t, out, ",Homo sapiens,2")
	})

	t.Run("ignores parsing with --version", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", "simple", "--version")
		require.NoError(t, err)
		assert.NotContains(t, out, ",Homo sapiens,")
		assert.Contains(t, out, "version:")
	})

	t.Run("sets format to default if -f value is unknown", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", ":)")
		require.NoError(t, err)
		assert.Contains(t, out, `Id,Verbatim,Cardinality,`)
	})
}

func TestStdin(t *testing.T) {
	t.Run("takes data from Stdin", func(t *testing.T) {
		out, _, err := runCLI(t, "Homo sapiens", "-f", "simple")
		require.NoError(t, err)
		assert.Contains(t, out, ",Homo sapiens,")
	})

	t.Run("takes multiple names from Stdin", func(t *testing.T) {
		out, _, err := runCLI(t, "Plantago\nBubo L.\n", "-f", "simple")
		require.NoError(t, err)
		assert.Contains(t, out, ",Plantago,")
		assert.Contains(t, out, ",Bubo,")
	})
}

func TestFlattenOutput(t *testing.T) {
	t.Run("flatten with JSON compact format", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens Linnaeus, 1758", "-f", "compact", "-F")
		require.NoError(t, err)
		assert.NotContains(t, out, `"canonical":`)
		assert.Contains(t, out, `"canonicalSimple"`)
		assert.Contains(t, out, `"canonicalFull"`)
		assert.Contains(t, out, `"authorship"`)
	})

	t.Run("flatten with JSON pretty format", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Bubo bubo", "-f", "pretty", "-F")
		require.NoError(t, err)
		assert.NotContains(t, out, `"canonical":`)
		assert.Contains(t, out, `"canonicalSimple"`)
	})

	t.Run("without flatten flag uses nested JSON", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", "compact")
		require.NoError(t, err)
		assert.Contains(t, out, `"canonical":`)
		assert.NotContains(t, out, `"canonicalSimple"`)
	})

	t.Run("CSV without details is simple", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", "csv")
		require.NoError(t, err)
		lines := strings.Split(out, "\n")
		header := lines[0]
		assert.Contains(t, header, "Id,Verbatim,Cardinality")
		assert.Contains(t, header, "NomCodeSetting")
		assert.NotContains(t, header, "Parsed,")
		assert.NotContains(t, header, ",Genus,")
		assert.NotContains(t, header, "CultivarEpithet")
	})

	t.Run("CSV with details includes all fields", func(t *testing.T) {
		out, _, err := runCLI(t, "", "Homo sapiens", "-f", "csv", "-d")
		require.NoError(t, err)
		assert.Contains(t, out, ",Genus,")
		assert.Contains(t, out, ",Species,")
		assert.Contains(t, out, ",Infraspecies")
		assert.Contains(t, out, "Parsed,")
		assert.Contains(t, out, ",Authors,")
	})

	t.Run("flatten from stdin", func(t *testing.T) {
		out, _, err := runCLI(t, "Homo sapiens", "-f", "compact", "-F")
		require.NoError(t, err)
		assert.NotContains(t, out, `"canonical":`)
		assert.Contains(t, out, `"canonicalSimple"`)
	})
}
