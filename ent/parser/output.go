package parser

import (
	"cmp"
	"slices"
	"strings"

	"github.com/gnames/gnparser/ent/parsed"
)

// ToOutput converts Abstract Syntax Tree of scientific name to a
// final output object.
func (sn *scientificNameNode) ToOutput(withDetails bool) parsed.Parsed {
	res := parsed.Parsed{
		Verbatim:      sn.verbatim,
		Canonical:     sn.Canonical(),
		Virus:         sn.virus,
		DaggerChar:    sn.daggerChar,
		VerbatimID:    sn.verbatimID,
		ParserVersion: sn.parserVersion,
	}

	if res.Canonical == nil {
		return res
	}

	res.Parsed = true
	res.ParseQuality, res.QualityWarnings = sn.qualityWarnings()
	res.Normalized = sn.Normalized()
	res.Cardinality = sn.cardinality
	res.Authorship = sn.LastAuthorship(withDetails)
	res.Hybrid = sn.hybrid
	res.Surrogate = sn.surrogate
	res.Bacteria = sn.bacteria
	res.Tail = sn.tail
	if withDetails {
		res.Details = sn.Details()
		res.Words = sn.Words()
	}

	if sn.ambiguousEpithet != "" {
		res.RestoreAmbiguous(sn.ambiguousEpithet, sn.ambiguousModif)
	}
	return res
}

func (sn *scientificNameNode) qualityWarnings() (int, []parsed.QualityWarning) {
	if sn.cardinality > 2 && sn.maybeFilius() {
		if sn.warnings == nil {
			sn.warnings = make(map[parsed.Warning]struct{})
		}
		sn.warnings[parsed.AuthAmbiguousFiliusWarn] = struct{}{}
	}

	warns := prepareWarnings(sn.warnings)
	quality := 1
	if len(warns) > 0 {
		quality = warns[0].Quality
	}
	return quality, warns
}

func (sn *scientificNameNode) maybeFilius() bool {
	words := sn.Words()
	for i := range words {
		if words[i].Verbatim != "f." {
			continue
		}
		if i == 0 || i == len(words)-1 {
			continue
		}

		betweenChars := sn.verbatim[words[i-1].End:words[i+1].Start]

		if words[i-1].Type == parsed.AuthorWordType &&
			words[i+1].Type == parsed.InfraspEpithetType &&
			!strings.Contains(betweenChars, ")") {
			return true
		}
	}
	return false
}

func prepareWarnings(ws map[parsed.Warning]struct{}) []parsed.QualityWarning {
	res := make([]parsed.QualityWarning, len(ws))
	var i int
	for k := range ws {
		res[i] = k.NewQualityWarning()
		i++
	}

	slices.SortFunc(res, func(a, b parsed.QualityWarning) int {
		res := cmp.Compare(b.Quality, a.Quality)
		if res != 0 {
			return res
		}
		return cmp.Compare(a.Warning.String(), b.Warning.String())
	})

	return res
}
