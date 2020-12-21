package parser

import (
	"sort"

	o "github.com/gnames/gnparser/entity/output"
)

func (sn *ScientificNameNode) ToOutput(withDetails bool) o.Parsed {
	res := o.Parsed{
		VerbatimID:    sn.VerbatimID,
		Verbatim:      sn.Verbatim,
		Virus:         sn.Virus,
		ParserVersion: sn.ParserVersion,
	}
	res.Canonical = sn.Canonical()

	if res.Canonical != nil {
		res.Parsed = true
		res.OverallQuality, res.QualityWarnings = processWarnings(sn.Warnings)
		res.Hybrid = sn.Hybrid
		res.Bacteria = sn.Bacteria
		res.Authorship = sn.LastAuthorship(withDetails)
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
