package icvcn

import (
	"slices"
	"strings"

	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnuuid"
)

type Parsed struct {
	Input         string
	Rank          Rank
	Uninomial     string
	Genus         string
	Species       string
	Words         []parsed.Word
	Error         error
	Parsed        bool
	ParserVersion string
}

func (p *Parsed) ToOutput(withDetails, _ bool) parsed.Parsed {
	normalized, warns := p.normalize()
	quality := 1
	if len(warns) > 0 {
		quality = slices.MaxFunc(warns, func(x, y parsed.QualityWarning) int {
			if x.Quality < y.Quality {
				return -1
			} else if x.Quality > y.Quality {
				return 1
			}
			return 0
		}).Quality
	}

	// If parsing failed, return unparsed result
	if !p.Parsed {
		return parsed.Parsed{
			Parsed:        false,
			NomCode:       "ICVCN",
			ParseQuality:  0,
			Verbatim:      p.Input,
			Virus:         false,
			VerbatimID:    gnuuid.New(p.Input).String(),
			ParserVersion: p.ParserVersion,
		}
	}

	res := parsed.Parsed{
		Parsed:          true,
		NomCode:         "ICVCN",
		ParseQuality:    quality,
		QualityWarnings: warns,
		Verbatim:        p.Input,
		Normalized:      normalized,
		Canonical: &parsed.Canonical{
			Simple:  normalized,
			Stemmed: normalized,
			Full:    normalized,
		},
		Cardinality:   p.cardinality(),
		Rank:          p.Rank.String(),
		Virus:         true,
		VerbatimID:    gnuuid.New(p.Input).String(),
		ParserVersion: p.ParserVersion,
	}
	if withDetails {
		res.Words = p.Words
		res.Details = p.details()
	}
	return res
}

func (p *Parsed) normalize() (string, []parsed.QualityWarning) {
	res := p.Uninomial
	verbatim := strings.TrimSpace(p.Input)
	if p.Species != "" {
		res = res + " " + p.Species
	}
	var warns []parsed.QualityWarning
	if len(res) != len(verbatim) && res != verbatim {
		qw := parsed.SpaceNonStandardWarn.NewQualityWarning()
		warns = append(warns, qw)
	}
	if verbatim != p.Input {
		qw := parsed.WhiteSpaceTrailWarn.NewQualityWarning()
		warns = append(warns, qw)
	}
	return res, warns
}

func (p *Parsed) cardinality() int {
	res := 1
	if p.Species != "" {
		res = 2
	}
	return res
}

func (p *Parsed) details() parsed.Details {
	if p.Species != "" {
		// Binomial virus name
		return parsed.DetailsSpeciesICVCN{
			SpeciesICVCN: parsed.SpeciesICVCN{
				Genus:   p.Genus,
				Species: p.Species,
				Rank:    p.Rank.String(),
			},
		}
	}

	// Uninomial virus name
	return parsed.DetailsUninomialICVCN{
		UninomialICVCN: parsed.UninomialICVCN{
			Value: p.Uninomial,
			Rank:  p.Rank.String(),
		},
	}
}
