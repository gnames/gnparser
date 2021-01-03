package parser

import (
	"sort"

	o "github.com/gnames/gnparser/entity/output"
)

func (sn *ScientificNameNode) ToOutput(withDetails bool) o.Parsed {
	res := o.Parsed{
		Verbatim:      sn.Verbatim,
		Canonical:     sn.Canonical(),
		Virus:         sn.Virus,
		VerbatimID:    sn.VerbatimID,
		ParserVersion: sn.ParserVersion,
	}

	if res.Canonical == nil {
		return res
	}

	res.Parsed = true
	res.OverallQuality, res.QualityWarnings = processWarnings(sn.Warnings)
	res.Normalized = sn.Normalized()
	res.Cardinality = sn.Cardinality
	res.Authorship = sn.LastAuthorship(withDetails)
	res.Hybrid = sn.Hybrid
	res.Bacteria = sn.Bacteria
	res.Tail = sn.Tail
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
