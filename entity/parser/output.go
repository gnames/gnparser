package parser

import (
	"sort"

	o "github.com/gnames/gnparser/entity/output"
)

func (sn *scientificNameNode) ToOutput(withDetails bool) o.Parsed {
	res := o.Parsed{
		Verbatim:      sn.verbatim,
		Canonical:     sn.Canonical(),
		Virus:         sn.virus,
		VerbatimID:    sn.verbatimID,
		ParserVersion: sn.parserVersion,
	}

	if res.Canonical == nil {
		return res
	}

	res.Parsed = true
	res.OverallQuality, res.QualityWarnings = processWarnings(sn.warnings)
	res.Normalized = sn.Normalized()
	res.Cardinality = sn.cardinality
	res.Authorship = sn.LastAuthorship(withDetails)
	res.Hybrid = sn.hybrid
	res.Bacteria = sn.bacteria
	res.Tail = sn.tail
	if withDetails {
		res.Details = sn.Details()
		res.Positions = sn.Pos()
	}
	return res
}

func processWarnings(ws []o.Warning) (int, []o.QualityWarning) {
	warns := prepareWarnings(ws)
	quality := 1
	if len(warns) > 0 {
		quality = warns[0].Quality
	}
	return quality, warns
}

func prepareWarnings(ws []o.Warning) []o.QualityWarning {
	res := make([]o.QualityWarning, len(ws))
	for i := range ws {
		res[i] = ws[i].NewQualityWarning()
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Quality > res[j].Quality {
			return true
		}
		if res[i].Quality < res[j].Quality {
			return false
		}
		return res[i].Warning < res[j].Warning
	})
	return res
}
