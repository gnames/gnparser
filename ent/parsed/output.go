package parsed

import (
	"strconv"

	"github.com/gnames/gnfmt"
)

// Output creates a JSON or CSV representation of Parsed results.
func (p Parsed) Output(f gnfmt.Format) string {
	switch f {
	case gnfmt.CSV:
		return p.csvOutput(',')
	case gnfmt.TSV:
		return p.csvOutput('\t')
	case gnfmt.CompactJSON:
		return p.jsonOutput(false)
	case gnfmt.PrettyJSON:
		return p.jsonOutput(true)
	default:
		return "N/A"
	}
}

// HeadersCSV returns the CSV header for parsing output.
func HeaderCSV(f gnfmt.Format) string {
	header := []string{"Id", "Verbatim", "Cardinality", "CanonicalStem",
		"CanonicalSimple", "CanonicalFull", "Authorship", "Year", "Quality"}
	switch f {
	case gnfmt.CSV:
		return gnfmt.ToCSV(header, ',')
	case gnfmt.TSV:
		return gnfmt.ToCSV(header, '\t')
	default:
		return ""
	}
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

func (p Parsed) jsonOutput(pretty bool) string {
	enc := gnfmt.GNjson{Pretty: pretty}
	res, _ := enc.Encode(p)
	return string(res)
}
