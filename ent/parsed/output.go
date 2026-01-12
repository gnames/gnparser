package parsed

import (
	"strconv"

	"github.com/gnames/gnfmt"
)

// Output creates a JSON or CSV representation of Parsed results.
// When flatten is true, JSON output uses the flattened structure.
// For CSV/TSV, flatten is always used and withDetails is determined by
// whether the Details field is populated.
func (p Parsed) Output(f gnfmt.Format, flatten bool) string {
	switch f {
	case gnfmt.CSV:
		return p.csvOutputFlat(',', p.hasDetails())
	case gnfmt.TSV:
		return p.csvOutputFlat('\t', p.hasDetails())
	case gnfmt.CompactJSON:
		return p.jsonOutput(false, flatten)
	case gnfmt.PrettyJSON:
		return p.jsonOutput(true, flatten)
	default:
		return "N/A"
	}
}

// hasDetails returns true if the Details field is populated,
// indicating that WithDetails was true during parsing.
func (p Parsed) hasDetails() bool {
	return p.Details != nil
}

// HeadersCSV returns the CSV header for parsing output.
// The withDetails parameter determines whether to include name component columns.
func HeaderCSV(f gnfmt.Format, withDetails bool) string {
	return HeaderCSVFlat(f, withDetails)
}

func (p Parsed) csvOutput(sep rune) string {
	var stem, simple, full, authorship, year string
	if p.Canonical != nil {
		stem = p.Canonical.Stemmed
		simple = p.Canonical.Simple
		full = p.Canonical.Full
	}

	if p.Authorship != nil {
		authorship = p.Authorship.Verbatim
		year = p.Authorship.Year
	}

	res := []string{
		p.VerbatimID,
		p.Verbatim,
		strconv.Itoa(p.Cardinality),
		stem,
		simple,
		full,
		authorship,
		year,
		strconv.Itoa(p.ParseQuality),
	}
	return gnfmt.ToCSV(res, sep)
}

func (p Parsed) jsonOutput(pretty bool, flatten bool) string {
	enc := gnfmt.GNjson{Pretty: pretty}
	var res []byte
	if flatten {
		res, _ = enc.Encode(p.Flatten())
	} else {
		res, _ = enc.Encode(p)
	}
	return string(res)
}

// HeaderCSVFlat returns the CSV/TSV header for flattened output.
// When withDetails is false, shows only the most commonly used fields.
// When withDetails is true, includes all flattened fields.
func HeaderCSVFlat(f gnfmt.Format, withDetails bool) string {
	// Traditional 9 columns + NomCodeSetting (backward compatible, simple)
	header := []string{
		"Id",
		"Verbatim",
		"Cardinality",
		"CanonicalStem",
		"CanonicalSimple",
		"CanonicalFull",
		"Authorship",
		"Year",
		"Quality",
		"NomCodeSetting",
	}

	// Extended columns only with WithDetails
	if withDetails {
		header = append(header,
			"Parsed",
			"ParserVersion",
			"Normalized",
			"Rank",
			"Authors",
			"BasionymAuthorship",
			"BasionymExAuthorship",
			"BasionymAuthorshipYear",
			"CombinationAuthorship",
			"CombinationExAuthorship",
			"CombinationAuthorshipYear",
			"Candidatus",
			"Virus",
			"Cultivar",
			"DaggerChar",
			"Hybrid",
			"GraftChimera",
			"Surrogate",
			"Tail",
			"CultivarEpithet",
			"Notho",
			"Uninomial",
			"Genus",
			"Subgenus",
			"Species",
			"Infraspecies",
		)
	}

	switch f {
	case gnfmt.CSV:
		return gnfmt.ToCSV(header, ',')
	case gnfmt.TSV:
		return gnfmt.ToCSV(header, '\t')
	default:
		return ""
	}
}

// csvOutputFlat returns a CSV/TSV row using the flattened structure.
// When withDetails is false, returns only most commonly used fields.
// When withDetails is true, includes all flattened fields.
func (p Parsed) csvOutputFlat(sep rune, withDetails bool) string {
	pf := p.Flatten()

	// Derive Year field for backward compatibility from Authorship.Year
	var year string
	if p.Authorship != nil {
		year = p.Authorship.Year
	}

	row := []string{
		pf.VerbatimID,
		pf.Verbatim,
		strconv.Itoa(pf.Cardinality),
		pf.CanonicalStemmed,
		pf.CanonicalSimple,
		pf.CanonicalFull,
		pf.Authorship,
		year,
		strconv.Itoa(pf.ParseQuality),
		pf.NomCodeSetting,
	}

	// Extended columns only with WithDetails
	if withDetails {
		row = append(row,
			strconv.FormatBool(pf.Parsed),
			pf.ParserVersion,
			pf.Normalized,
			pf.Rank,
			pf.Authors,
			pf.BasionymAuthorship,
			pf.BasionymExAuthorship,
			pf.BasionymAuthorshipYear,
			pf.CombinationAuthorship,
			pf.CombinationExAuthorship,
			pf.CombinationAuthorshipYear,
			strconv.FormatBool(pf.Candidatus),
			strconv.FormatBool(pf.Virus),
			strconv.FormatBool(pf.Cultivar),
			strconv.FormatBool(pf.DaggerChar),
			pf.Hybrid,
			pf.GraftChimera,
			pf.Surrogate,
			pf.Tail,
			pf.CultivarEpithet,
			pf.Notho,
			pf.Uninomial,
			pf.Genus,
			pf.Subgenus,
			pf.Species,
			pf.Infraspecies,
		)
	}

	return gnfmt.ToCSV(row, sep)
}
