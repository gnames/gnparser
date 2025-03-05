package parsed_test

import (
	"testing"

	"github.com/gnames/gnparser/ent/parsed"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		input    parsed.Parsed
		expected parsed.ParsedFlat
	}{
		{
			name: "Parsed with all fields",
			input: parsed.Parsed{
				Parsed:        true,
				NomCode:       "ICZN",
				ParseQuality:  1,
				Verbatim:      "Aus bus",
				Normalized:    "Aus bus",
				Cardinality:   2,
				Rank:          "species",
				Bacteria:      nil,
				Candidatus:    true,
				Virus:         false,
				Cultivar:      false,
				DaggerChar:    false,
				Hybrid:        nil,
				GraftChimera:  nil,
				Surrogate:     nil,
				Tail:          "tail",
				VerbatimID:    "12345",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple:  "Aus bus",
					Full:    "Aus bus",
					Stemmed: "Aus bus",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "L.",
					Original: &parsed.AuthGroup{
						Authors: []string{"Linnaeus"},
						Year:    &parsed.Year{Value: "1758"},
					},
					Combination: &parsed.AuthGroup{
						Authors: []string{"Smith"},
						Year:    &parsed.Year{Value: "1800"},
					},
				},
				Details: parsed.DetailsSpecies{
					Species: parsed.Species{
						Genus:   "Aus",
						Species: "bus",
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:                    true,
				NomCode:                   "ICZN",
				ParseQuality:              1,
				Verbatim:                  "Aus bus",
				Normalized:                "Aus bus",
				Cardinality:               2,
				Rank:                      "species",
				Bacteria:                  "",
				Candidatus:                true,
				Virus:                     false,
				Cultivar:                  false,
				DaggerChar:                false,
				Hybrid:                    "",
				GraftChimera:              "",
				Surrogate:                 "",
				Tail:                      "tail",
				VerbatimID:                "12345",
				ParserVersion:             "1.0.0",
				CanonicalSimple:           "Aus bus",
				CanonicalFull:             "Aus bus",
				CanonicalStemmed:          "Aus bus",
				Authorship:                "L.",
				BasionymAuthorship:        "Linnaeus",
				BasionymAuthorshipYear:    "1758",
				CombinationAuthorship:     "Smith",
				CombinationAuthorshipYear: "1800",
				Genus:                     "Aus",
				Subgenus:                  "",
				Species:                   "bus",
			},
		},
		{
			name: "Parsed with minimal fields",
			input: parsed.Parsed{
				Parsed:        false,
				NomCode:       "ICZN",
				ParseQuality:  0,
				Verbatim:      "Unknown",
				VerbatimID:    "67890",
				ParserVersion: "1.0.0",
			},
			expected: parsed.ParsedFlat{
				Parsed:        false,
				NomCode:       "ICZN",
				ParseQuality:  0,
				Verbatim:      "Unknown",
				VerbatimID:    "67890",
				ParserVersion: "1.0.0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Flatten()
			assert.Equal(t, tt.expected, result)
		})
	}
}
