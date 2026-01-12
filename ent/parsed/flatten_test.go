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
				Parsed:         true,
				NomCodeSetting: "ICZN",
				ParseQuality:   1,
				Verbatim:       "Aus bus",
				Normalized:     "Aus bus",
				Cardinality:    2,
				Rank:           "species",
				Candidatus:     true,
				Virus:          false,
				Cultivar:       false,
				DaggerChar:     false,
				Hybrid:         nil,
				GraftChimera:   nil,
				Surrogate:      nil,
				Tail:           "tail",
				VerbatimID:     "12345",
				ParserVersion:  "1.0.0",
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
				NomCodeSetting:            "ICZN",
				ParseQuality:              1,
				Verbatim:                  "Aus bus",
				Normalized:                "Aus bus",
				Cardinality:               2,
				Rank:                      "species",
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
				Parsed:         false,
				NomCodeSetting: "ICZN",
				ParseQuality:   0,
				Verbatim:       "Unknown",
				VerbatimID:     "67890",
				ParserVersion:  "1.0.0",
			},
			expected: parsed.ParsedFlat{
				Parsed:         false,
				NomCodeSetting: "ICZN",
				ParseQuality:   0,
				Verbatim:       "Unknown",
				VerbatimID:     "67890",
				ParserVersion:  "1.0.0",
			},
		},
		{
			name: "Uninomial with rank",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Asteraceae",
				Cardinality:   1,
				Rank:          "family",
				VerbatimID:    "test-123",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Asteraceae",
				},
				Details: parsed.DetailsUninomial{
					Uninomial: parsed.Uninomial{
						Value: "Asteraceae",
						Rank:  "family",
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "Asteraceae",
				Cardinality:     1,
				Rank:            "family",
				VerbatimID:      "test-123",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Asteraceae",
				Uninomial:       "Asteraceae",
			},
		},
		{
			name: "Infraspecies with rank",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus bus var. cus",
				Cardinality:   3,
				VerbatimID:    "test-456",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus bus cus",
				},
				Details: parsed.DetailsInfraspecies{
					Infraspecies: parsed.Infraspecies{
						Species: parsed.Species{
							Genus:   "Aus",
							Species: "bus",
						},
						Infraspecies: []parsed.InfraspeciesElem{
							{
								Value: "cus",
								Rank:  "var.",
							},
						},
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "Aus bus var. cus",
				Cardinality:     3,
				Rank:            "var.",
				VerbatimID:      "test-456",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Aus bus cus",
				Genus:           "Aus",
				Species:         "bus",
				Infraspecies:    "cus",
			},
		},
		{
			name: "Ex-authorship in combination",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus bus",
				VerbatimID:    "test-789",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus bus",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "ex Smith Jones",
					Combination: &parsed.AuthGroup{
						Authors: []string{"Jones"},
						ExAuthors: &parsed.Authors{
							Authors: []string{"Smith"},
						},
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:                  true,
				Verbatim:                "Aus bus",
				VerbatimID:              "test-789",
				ParserVersion:           "1.0.0",
				CanonicalSimple:         "Aus bus",
				Authorship:              "ex Smith Jones",
				CombinationAuthorship:   "Jones",
				CombinationExAuthorship: "Smith",
			},
		},
		{
			name: "Ex-authorship in basionym",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus bus",
				VerbatimID:    "test-101",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus bus",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "(ex Brown White) Black",
					Original: &parsed.AuthGroup{
						Authors: []string{"White"},
						ExAuthors: &parsed.Authors{
							Authors: []string{"Brown"},
						},
					},
					Combination: &parsed.AuthGroup{
						Authors: []string{"Black"},
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:                true,
				Verbatim:              "Aus bus",
				VerbatimID:            "test-101",
				ParserVersion:         "1.0.0",
				CanonicalSimple:       "Aus bus",
				Authorship:            "(ex Brown White) Black",
				BasionymAuthorship:    "White",
				BasionymExAuthorship:  "Brown",
				CombinationAuthorship: "Black",
			},
		},
		{
			name: "Approximate year",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus bus",
				VerbatimID:    "test-202",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus bus",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "Smith (1800)",
					Combination: &parsed.AuthGroup{
						Authors: []string{"Smith"},
						Year: &parsed.Year{
							Value:         "1800",
							IsApproximate: true,
						},
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:                    true,
				Verbatim:                  "Aus bus",
				VerbatimID:                "test-202",
				ParserVersion:             "1.0.0",
				CanonicalSimple:           "Aus bus",
				Authorship:                "Smith (1800)",
				CombinationAuthorship:     "Smith",
				CombinationAuthorshipYear: "(1800)",
			},
		},
		{
			name: "Multiple authors formatting",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus bus",
				VerbatimID:    "test-303",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus bus",
				},
				Authorship: &parsed.Authorship{
					Verbatim: "Smith, Jones & Brown",
					Authors:  []string{"Smith", "Jones", "Brown"},
					Combination: &parsed.AuthGroup{
						Authors: []string{"Smith", "Jones", "Brown"},
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:                true,
				Verbatim:              "Aus bus",
				VerbatimID:            "test-303",
				ParserVersion:         "1.0.0",
				CanonicalSimple:       "Aus bus",
				Authorship:            "Smith, Jones & Brown",
				Authors:               "Smith|Jones|Brown",
				CombinationAuthorship: "Smith, Jones & Brown",
			},
		},
		{
			name: "Hybrid string conversion",
			input: func() parsed.Parsed {
				hybrid := parsed.NamedHybridAnnot
				return parsed.Parsed{
					Parsed:        true,
					Verbatim:      "× Aus bus",
					VerbatimID:    "test-404",
					ParserVersion: "1.0.0",
					Canonical: &parsed.Canonical{
						Simple: "Aus bus",
					},
					Hybrid: &hybrid,
				}
			}(),
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "× Aus bus",
				VerbatimID:      "test-404",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Aus bus",
				Hybrid:          "NAMED_HYBRID",
			},
		},
		{
			name: "GraftChimera string conversion",
			input: func() parsed.Parsed {
				graft := parsed.NamedGraftChimeraAnnot
				return parsed.Parsed{
					Parsed:        true,
					Verbatim:      "+ Aus bus",
					VerbatimID:    "test-505",
					ParserVersion: "1.0.0",
					Canonical: &parsed.Canonical{
						Simple: "Aus bus",
					},
					GraftChimera: &graft,
				}
			}(),
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "+ Aus bus",
				VerbatimID:      "test-505",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Aus bus",
				GraftChimera:    "NAMED_GRAFT_CHIMERA",
			},
		},
		{
			name: "Surrogate string conversion",
			input: func() parsed.Parsed {
				surrogate := parsed.ComparisonAnnot
				return parsed.Parsed{
					Parsed:        true,
					Verbatim:      "Aus cf. bus",
					VerbatimID:    "test-606",
					ParserVersion: "1.0.0",
					Canonical: &parsed.Canonical{
						Simple: "Aus cf. bus",
					},
					Surrogate: &surrogate,
				}
			}(),
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "Aus cf. bus",
				VerbatimID:      "test-606",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Aus cf. bus",
				Surrogate:       "COMPARISON",
			},
		},
		{
			name: "Species with subgenus",
			input: parsed.Parsed{
				Parsed:        true,
				Verbatim:      "Aus (Bus) cus",
				Cardinality:   2,
				VerbatimID:    "test-707",
				ParserVersion: "1.0.0",
				Canonical: &parsed.Canonical{
					Simple: "Aus cus",
				},
				Details: parsed.DetailsSpecies{
					Species: parsed.Species{
						Genus:    "Aus",
						Subgenus: "Bus",
						Species:  "cus",
					},
				},
			},
			expected: parsed.ParsedFlat{
				Parsed:          true,
				Verbatim:        "Aus (Bus) cus",
				Cardinality:     2,
				VerbatimID:      "test-707",
				ParserVersion:   "1.0.0",
				CanonicalSimple: "Aus cus",
				Genus:           "Aus",
				Subgenus:        "Bus",
				Species:         "cus",
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

// TestFlattenCultivarEpithet tests that cultivar epithet is properly flattened
// This requires parsing with cultivar support to populate the field
func TestFlattenCultivarEpithet(t *testing.T) {
	// Note: This test would ideally use actual parser output with cultivar names
	// For now, we test that the field is properly copied if present
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "Rosa 'Peace'",
		Cultivar:      true,
		VerbatimID:    "test-808",
		ParserVersion: "1.0.0",
		Canonical: &parsed.Canonical{
			Simple: "Rosa",
		},
	}

	flat := p.Flatten()
	assert.True(t, flat.Cultivar)
	assert.Equal(t, "Rosa 'Peace'", flat.Verbatim)
	assert.Equal(t, "Rosa", flat.CanonicalSimple)
}

// TestFlattenInfraspeciesMultiple tests that only single infraspecies rank is flattened
// Flatten does not cover cardinality > 3 for now.
func TestFlattenInfraspeciesMultiple(t *testing.T) {
	// Multiple infraspecies elements should not populate the flattened fields
	p := parsed.Parsed{
		Parsed:        true,
		Verbatim:      "Aus bus var. cus f. dus",
		Cardinality:   4,
		VerbatimID:    "test-909",
		ParserVersion: "1.0.0",
		Canonical: &parsed.Canonical{
			Simple: "Aus bus cus dus",
		},
		Details: parsed.DetailsInfraspecies{
			Infraspecies: parsed.Infraspecies{
				Species: parsed.Species{
					Genus:   "Aus",
					Species: "bus",
				},
				Infraspecies: []parsed.InfraspeciesElem{
					{Value: "cus", Rank: "var."},
					{Value: "dus", Rank: "f."},
				},
			},
		},
	}

	flat := p.Flatten()
	// With multiple infraspecies, genus/species/infraspecies should be empty
	assert.Equal(t, "", flat.Genus)
	assert.Equal(t, "", flat.Species)
	assert.Equal(t, "", flat.Infraspecies)
}
