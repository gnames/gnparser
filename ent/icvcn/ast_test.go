package icvcn_test

import (
	"testing"

	"github.com/gnames/gnparser/ent/icvcn"
)

func TestParseToAST_Species(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantParsed    bool
		wantGenus     string
		wantSpecies   string
		wantRank      icvcn.Rank
		wantUninomial string
	}{
		{
			name:          "species with viriform genus",
			input:         "Bracoviriform congregatae",
			wantParsed:    true,
			wantGenus:     "Bracoviriform",
			wantSpecies:   "congregatae",
			wantRank:      icvcn.Species,
			wantUninomial: "Bracoviriform",
		},
		{
			name:          "species with alphanumeric epithet",
			input:         "Plasmavirus L2",
			wantParsed:    true,
			wantGenus:     "Plasmavirus",
			wantSpecies:   "L2",
			wantRank:      icvcn.Species,
			wantUninomial: "Plasmavirus",
		},
		{
			name:          "species with capitalized epithet",
			input:         "Omtjevirus Omtje",
			wantParsed:    true,
			wantGenus:     "Omtjevirus",
			wantSpecies:   "Omtje",
			wantRank:      icvcn.Species,
			wantUninomial: "Omtjevirus",
		},
		{
			name:          "species with complex alphanumeric epithet",
			input:         "Cebaduodecimvirus phi12duo",
			wantParsed:    true,
			wantGenus:     "Cebaduodecimvirus",
			wantSpecies:   "phi12duo",
			wantRank:      icvcn.Species,
			wantUninomial: "Cebaduodecimvirus",
		},
		{
			name:          "species with lowercase epithet and number",
			input:         "Ahphunavirus yong1",
			wantParsed:    true,
			wantGenus:     "Ahphunavirus",
			wantSpecies:   "yong1",
			wantRank:      icvcn.Species,
			wantUninomial: "Ahphunavirus",
		},
		{
			name:          "species with complex epithet containing dash",
			input:         "Seodaemunguvirus YMC16-01N133",
			wantParsed:    true,
			wantGenus:     "Seodaemunguvirus",
			wantSpecies:   "YMC16-01N133",
			wantRank:      icvcn.Species,
			wantUninomial: "Seodaemunguvirus",
		},
		{
			name:          "valid species name",
			input:         "Triavirus phi2958PVL",
			wantParsed:    true,
			wantGenus:     "Triavirus",
			wantSpecies:   "phi2958PVL",
			wantRank:      icvcn.Species,
			wantUninomial: "Triavirus",
		},
		{
			name:          "species with lowercase epithet",
			input:         "Orthobunyavirus encephalitidis",
			wantParsed:    true,
			wantGenus:     "Orthobunyavirus",
			wantSpecies:   "encephalitidis",
			wantRank:      icvcn.Species,
			wantUninomial: "Orthobunyavirus",
		},
		{
			name:          "species of satellite",
			input:         "Deltasatellite desmodii",
			wantParsed:    true,
			wantGenus:     "Deltasatellite",
			wantSpecies:   "desmodii",
			wantRank:      icvcn.Species,
			wantUninomial: "Deltasatellite",
		},
		{
			name:          "species of viroid",
			input:         "Pospiviroid chloronani",
			wantParsed:    true,
			wantGenus:     "Pospiviroid",
			wantSpecies:   "chloronani",
			wantRank:      icvcn.Species,
			wantUninomial: "Pospiviroid",
		},
		{
			name:       "species of mammal",
			input:      "Homo sapiens",
			wantParsed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &icvcn.Parser{Buffer: tt.input}
			p.Init()
			result := p.ParseToStruct()

			if result.Parsed != tt.wantParsed {
				t.Errorf("ParseToAST() Parsed = %v, want %v", result.Parsed, tt.wantParsed)
				if result.Error != nil {
					t.Errorf("ParseToAST() Error = %v", result.Error)
				}
				return
			}

			if result.Genus != tt.wantGenus {
				t.Errorf("ParseToAST() Genus = %v, want %v", result.Genus, tt.wantGenus)
			}

			if result.Species != tt.wantSpecies {
				t.Errorf("ParseToAST() Species = %v, want %v", result.Species, tt.wantSpecies)
			}

			if result.Rank != tt.wantRank {
				t.Errorf(
					"ParseToAST() UninomialRank = %v, want %v",
					result.Rank,
					tt.wantRank,
				)
			}

			if result.Uninomial != tt.wantUninomial {
				t.Errorf("ParseToAST() Uninomial = %v, want %v", result.Uninomial, tt.wantUninomial)
			}
		})
	}
}

func TestParseToAST_Uninomial(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantParsed bool
		wantRank   icvcn.Rank
	}{
		{
			name:       "kingdom",
			input:      "Bamfordvirae",
			wantParsed: true,
			wantRank:   icvcn.Kingdom,
		},
		{
			name:       "phylum",
			input:      "Preplasmiviricota",
			wantParsed: true,
			wantRank:   icvcn.Phylum,
		},
		{
			name:       "subphylum",
			input:      "Polisuviricotina",
			wantParsed: true,
			wantRank:   icvcn.Subphylum,
		},
		{
			name:       "class",
			input:      "Virophaviricetes",
			wantParsed: true,
			wantRank:   icvcn.Class,
		},
		{
			name:       "order",
			input:      "Rowavirales",
			wantParsed: true,
			wantRank:   icvcn.Order,
		},
		{
			name:       "family",
			input:      "Adenoviridae",
			wantParsed: true,
			wantRank:   icvcn.Family,
		},
		{
			name:       "family of satellite",
			input:      "Alphasatellitidae",
			wantParsed: true,
			wantRank:   icvcn.Family,
		},
		{
			name:       "subfamily of satellite",
			input:      "Geminialphasatellitinae",
			wantParsed: true,
			wantRank:   icvcn.Subfamily,
		},
		{
			name:       "genus of satellite",
			input:      "Clecrusatellite",
			wantParsed: true,
			wantRank:   icvcn.Genus,
		},
		{
			name:       "family of viriform",
			input:      "Bartogtaviriformidae",
			wantParsed: true,
			wantRank:   icvcn.Family,
		},
		{
			name:       "genus of viriform",
			input:      "Ichnoviriform",
			wantParsed: true,
			wantRank:   icvcn.Genus,
		},
		{
			name:       "family of viroid",
			input:      "Pospiviroidae",
			wantParsed: true,
			wantRank:   icvcn.Family,
		},
		{
			name:       "genus of viroid",
			input:      "Avsunviroid",
			wantParsed: true,
			wantRank:   icvcn.Genus,
		},
		{
			name:       "realm",
			input:      "Riboviria",
			wantParsed: true,
			wantRank:   icvcn.Realm,
		},
		{
			name:       "realm",
			input:      "Riboviria",
			wantParsed: true,
			wantRank:   icvcn.Realm,
		},
		{
			name:       "family",
			input:      "Flaviviridae",
			wantParsed: true,
			wantRank:   icvcn.Family,
		},
		{
			name:       "genus",
			input:      "Flavivirus",
			wantParsed: true,
			wantRank:   icvcn.Genus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &icvcn.Parser{Buffer: tt.input}
			p.Init()
			result := p.ParseToStruct()

			if result.Parsed != tt.wantParsed {
				t.Errorf("ParseToAST() Parsed = %v, want %v", result.Parsed, tt.wantParsed)
				if result.Error != nil {
					t.Errorf("ParseToAST() Error = %v", result.Error)
				}
				return
			}

			if result.Rank != tt.wantRank {
				t.Errorf("ParseToAST() UninomialRank = %v (%s), want %v (%s)",
					result.Rank, result.Rank.String(),
					tt.wantRank, tt.wantRank.String())
			}

			if result.Uninomial == "" {
				t.Errorf("ParseToAST() Uninomial is empty, expected value")
			}
		})
	}
}

func TestParseToAST_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "invalid suffix",
			input: "Invalidname",
		},
		{
			name:  "empty input",
			input: "",
		},
		{
			name:  "random text",
			input: "not a virus name at all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &icvcn.Parser{Buffer: tt.input}
			p.Init()
			result := p.ParseToStruct()

			if result.Parsed {
				t.Errorf(
					"ParseToAST() Parsed = true for invalid input %q, expected false",
					tt.input,
				)
			}

			if result.Error == nil && tt.input != "" {
				t.Errorf("ParseToAST() Error = nil for invalid input %q, expected error", tt.input)
			}
		})
	}
}

func TestParse_Blacklist(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantParsed bool
		reason     string
	}{
		// All blacklisted names should not parse
		{
			name:       "blacklisted Calviria",
			input:      "Calviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		{
			name:       "blacklisted Caviria",
			input:      "Caviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		{
			name:       "blacklisted Corvira",
			input:      "Corvira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Dravira",
			input:      "Dravira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Elvira",
			input:      "Elvira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Eugivira",
			input:      "Eugivira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Euvira",
			input:      "Euvira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Givira",
			input:      "Givira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Kaviria",
			input:      "Kaviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		{
			name:       "blacklisted Lussanvira",
			input:      "Lussanvira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Mahavira",
			input:      "Mahavira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Miracavira",
			input:      "Miracavira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Navira",
			input:      "Navira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Paracalviria",
			input:      "Paracalviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		{
			name:       "blacklisted Roselviria",
			input:      "Roselviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		{
			name:       "blacklisted Rovira",
			input:      "Rovira",
			wantParsed: false,
			reason:     "botanical genus matching -vira suffix",
		},
		{
			name:       "blacklisted Selviria",
			input:      "Selviria",
			wantParsed: false,
			reason:     "botanical genus matching -viria suffix",
		},
		// Test whitespace trimming
		{
			name:       "blacklisted with leading space",
			input:      " Calviria",
			wantParsed: false,
			reason:     "should trim leading whitespace",
		},
		{
			name:       "blacklisted with trailing space",
			input:      "Elvira ",
			wantParsed: false,
			reason:     "should trim trailing whitespace",
		},
		{
			name:       "blacklisted with surrounding spaces",
			input:      " Rovira ",
			wantParsed: false,
			reason:     "should trim surrounding whitespace",
		},
		{
			name:       "blacklisted with tab",
			input:      "\tCalviria",
			wantParsed: false,
			reason:     "should trim tab characters",
		},
		// Valid ICVCN names should still parse
		{
			name:       "valid realm Riboviria",
			input:      "Riboviria",
			wantParsed: true,
			reason:     "legitimate ICVCN realm",
		},
		{
			name:       "valid realm Adnaviria",
			input:      "Adnaviria",
			wantParsed: true,
			reason:     "legitimate ICVCN realm",
		},
		{
			name:       "valid genus Flavivirus",
			input:      "Flavivirus",
			wantParsed: true,
			reason:     "legitimate ICVCN genus",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := icvcn.Parse(tt.input)

			if result.Parsed != tt.wantParsed {
				t.Errorf(
					"Parse(%q) Parsed = %v, want %v (%s)",
					tt.input,
					result.Parsed,
					tt.wantParsed,
					tt.reason,
				)
			}

			// Verify input is preserved
			if result.Input != tt.input {
				t.Errorf("Parse(%q) Input = %q, want %q", tt.input, result.Input, tt.input)
			}
		})
	}
}
