package parsed_test

import (
	"strings"
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/stretchr/testify/assert"
)

func TestOutput_JSON_Flatten(t *testing.T) {
	tests := []struct {
		name             string
		parsed           parsed.Parsed
		flatten          bool
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name: "Flatten JSON removes nested structures",
			parsed: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Homo sapiens",
				Cardinality:   2,
				VerbatimID:    "test-id",
				ParserVersion: "v1.0.0",
				Canonical: &parsed.Canonical{
					Simple:  "Homo sapiens",
					Full:    "Homo sapiens",
					Stemmed: "Homo sapiens",
				},
			},
			flatten: true,
			shouldContain: []string{
				`"canonicalSimple":"Homo sapiens"`,
				`"canonicalFull":"Homo sapiens"`,
				`"canonicalStemmed":"Homo sapiens"`,
			},
			shouldNotContain: []string{
				`"canonical":`,
			},
		},
		{
			name: "Non-flatten JSON keeps nested structures",
			parsed: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Homo sapiens",
				Cardinality:   2,
				VerbatimID:    "test-id",
				ParserVersion: "v1.0.0",
				Canonical: &parsed.Canonical{
					Simple:  "Homo sapiens",
					Full:    "Homo sapiens",
					Stemmed: "Homo sapiens",
				},
			},
			flatten: false,
			shouldContain: []string{
				`"canonical":`,
				`"simple":"Homo sapiens"`,
			},
			shouldNotContain: []string{
				`"canonicalSimple"`,
			},
		},
		{
			name: "Flatten JSON with authorship details",
			parsed: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Homo sapiens Linnaeus, 1758",
				Cardinality:   2,
				VerbatimID:    "test-id",
				ParserVersion: "v1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Homo sapiens",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "Linnaeus, 1758",
					Authors:  []string{"Linnaeus"},
					Combination: &parsed.AuthGroup{
						Authors: []string{"Linnaeus"},
						Year:    &parsed.Year{Value: "1758"},
					},
				},
			},
			flatten: true,
			shouldContain: []string{
				`"authorship":"Linnaeus, 1758"`,
				`"authors":"Linnaeus"`,
				`"combinationAuthorship":"Linnaeus"`,
				`"combinationAuthorshipYear":"1758"`,
			},
			shouldNotContain: []string{
				`"authorship":{`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.parsed.Output(gnfmt.CompactJSON, tt.flatten)
			for _, s := range tt.shouldContain {
				assert.Contains(t, output, s, "Should contain: %s", s)
			}
			for _, s := range tt.shouldNotContain {
				assert.NotContains(t, output, s, "Should not contain: %s", s)
			}
		})
	}
}

func TestOutput_CSV_AlwaysFlattened(t *testing.T) {
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "Homo sapiens",
		Cardinality:   2,
		ParseQuality:  1,
		VerbatimID:    "test-id",
		ParserVersion: "v1.0.0",
		Canonical: &parsed.Canonical{
			Simple:  "Homo sapiens",
			Full:    "Homo sapiens",
			Stemmed: "Homo sapiens",
		},
		Authorship: &parsed.Authorship{
			Verbatim: "Linnaeus 1758",
			Year:     "1758",
			Authors:  []string{"Linnaeus"},
			Combination: &parsed.AuthGroup{
				Authors: []string{"Linnaeus"},
				Year:    &parsed.Year{Value: "1758"},
			},
		},
	}

	// Test CSV output without details (simple, 10 fields)
	csvOutput := p.Output(gnfmt.CSV, false)
	assert.Contains(t, csvOutput, "test-id")
	assert.Contains(t, csvOutput, "Homo sapiens")
	assert.Contains(t, csvOutput, "Linnaeus 1758")
	assert.Contains(t, csvOutput, "1758")

	// CSV without details should have 10 fields (traditional 9 + NomCodeSetting)
	// Count by splitting on commas (works when authorship has no commas)
	fields := strings.Split(csvOutput, ",")
	assert.Equal(t, 10, len(fields), "CSV without details should have 10 fields")
}

func TestOutput_TSV_AlwaysFlattened(t *testing.T) {
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "Bubo bubo",
		Cardinality:   2,
		ParseQuality:  1,
		VerbatimID:    "test-id",
		ParserVersion: "v1.0.0",
		Canonical: &parsed.Canonical{
			Simple:  "Bubo bubo",
			Full:    "Bubo bubo",
			Stemmed: "Bubo bubo",
		},
	}

	// Test TSV output without details (simple, 10 fields)
	tsvOutput := p.Output(gnfmt.TSV, false)
	assert.Contains(t, tsvOutput, "test-id")
	assert.Contains(t, tsvOutput, "Bubo bubo")

	// TSV without details should have 10 fields
	fields := strings.Split(tsvOutput, "\t")
	assert.Equal(t, 10, len(fields), "TSV without details should have 10 fields")
}

func TestHeaderCSV_WithDetails(t *testing.T) {
	tests := []struct {
		name             string
		format           gnfmt.Format
		withDetails      bool
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name:        "CSV header without details (simple)",
			format:      gnfmt.CSV,
			withDetails: false,
			shouldContain: []string{
				"Id", "Verbatim", "Cardinality",
				"CanonicalStem", "CanonicalSimple", "CanonicalFull",
				"Authorship", "Year", "Quality",
				"NomCodeSetting",
			},
			shouldNotContain: []string{
				",Parsed,", "ParserVersion", "Normalized",
				",Authors,", ",Rank,", "Candidatus",
				"Uninomial", ",Genus,", "Subgenus", ",Species,", "Infraspecies",
			},
		},
		{
			name:        "CSV header with details (extended)",
			format:      gnfmt.CSV,
			withDetails: true,
			shouldContain: []string{
				"Id", "Verbatim", "Cardinality",
				",Parsed,", "ParserVersion", "Normalized",
				",Authors,", "CultivarEpithet", "Notho",
				"Uninomial", ",Genus,", "Subgenus", ",Species,", "Infraspecies",
			},
			shouldNotContain: []string{},
		},
		{
			name:        "TSV header without details (simple)",
			format:      gnfmt.TSV,
			withDetails: false,
			shouldContain: []string{
				"Id", "Verbatim", "NomCodeSetting",
			},
			shouldNotContain: []string{
				"\tParsed\t", "\tAuthors\t", "Uninomial", "Genus",
			},
		},
		{
			name:        "TSV header with details (extended)",
			format:      gnfmt.TSV,
			withDetails: true,
			shouldContain: []string{
				"Id", "Verbatim", "\tAuthors\t", "Genus", "Species",
			},
			shouldNotContain: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := parsed.HeaderCSV(tt.format, tt.withDetails)
			for _, s := range tt.shouldContain {
				assert.Contains(t, header, s, "Header should contain: %s", s)
			}
			for _, s := range tt.shouldNotContain {
				assert.NotContains(t, header, s, "Header should not contain: %s", s)
			}

			// Verify separator
			if tt.format == gnfmt.CSV {
				assert.Contains(t, header, ",")
			} else if tt.format == gnfmt.TSV {
				assert.Contains(t, header, "\t")
			}
		})
	}
}

func TestHeaderCSVFlat_FieldCount(t *testing.T) {
	tests := []struct {
		name           string
		format         gnfmt.Format
		withDetails    bool
		separator      string
		expectedFields int
	}{
		{
			name:           "CSV without details has 10 fields (simple)",
			format:         gnfmt.CSV,
			withDetails:    false,
			separator:      ",",
			expectedFields: 10,
		},
		{
			name:           "CSV with details has 36 fields (extended)",
			format:         gnfmt.CSV,
			withDetails:    true,
			separator:      ",",
			expectedFields: 36,
		},
		{
			name:           "TSV without details has 10 fields (simple)",
			format:         gnfmt.TSV,
			withDetails:    false,
			separator:      "\t",
			expectedFields: 10,
		},
		{
			name:           "TSV with details has 36 fields (extended)",
			format:         gnfmt.TSV,
			withDetails:    true,
			separator:      "\t",
			expectedFields: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := parsed.HeaderCSV(tt.format, tt.withDetails)
			fields := strings.Split(header, tt.separator)
			assert.Equal(t, tt.expectedFields, len(fields),
				"Expected %d fields, got %d", tt.expectedFields, len(fields))
		})
	}
}

func TestOutput_CSV_WithDetails(t *testing.T) {
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "Homo sapiens Linnaeus",
		Cardinality:   2,
		ParseQuality:  1,
		VerbatimID:    "test-id",
		ParserVersion: "v1.0.0",
		Canonical: &parsed.Canonical{
			Simple: "Homo sapiens",
		},
		Details: parsed.DetailsSpecies{
			Species: parsed.Species{
				Genus:   "Homo",
				Species: "sapiens",
			},
		},
	}

	// hasDetails() should return true when Details is set
	csvOutput := p.Output(gnfmt.CSV, false)
	fields := strings.Split(csvOutput, ",")

	// With details, should have 36 fields (10 base + 26 extended)
	assert.Equal(t, 36, len(fields), "CSV with details should have 36 fields")

	// Check that genus and species are in the output
	assert.Contains(t, csvOutput, "Homo")
	assert.Contains(t, csvOutput, "sapiens")
}

func TestOutput_InvalidFormat(t *testing.T) {
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "test",
		VerbatimID:    "test-id",
		ParserVersion: "v1.0.0",
	}

	// Unknown format should return "N/A"
	output := p.Output(gnfmt.Format(999), false)
	assert.Equal(t, "N/A", output)
}
