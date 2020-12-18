package output

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gnames/gnparser/grammar"
	"github.com/gnames/gnparser/stemmer"
)

type simple struct {
	ID              string
	Verbatim        string
	Cardinality     int
	CanonicalRanked string
	Canonical       string
	CanonicalStem   string
	Authorship      string
	Year            string
	Quality         int
}

func NewSimpleOutput(sn *grammar.ScientificNameNode) *simple {
	ao := sn.LastAuthorship()
	authorship := ""
	if ao != nil {
		authorship = ao.Value
	}
	yr := ""
	if ao != nil && ao.Original.Year != nil {
		yr = ao.Original.Year.Value
		if ao.Original.Year.Approximate {
			yr = fmt.Sprintf("(%s)", yr)
		}
	}

	quality := 0
	c := sn.Canonical()
	if c == nil {
		c = &grammar.Canonical{}
	} else {
		_, quality = qualityAndWarnings(sn.Warnings)
	}

	so := simple{
		ID:              sn.VerbatimID,
		Verbatim:        sn.Verbatim,
		Cardinality:     sn.Cardinality,
		CanonicalRanked: c.ValueRanked,
		Canonical:       c.Value,
		CanonicalStem:   stemmer.StemCanonical(c.Value),
		Authorship:      authorship,
		Year:            yr,
		Quality:         quality,
	}
	return &so
}

func CSVHeader() string {
	header := ([]string{
		"Id",
		"Verbatim",
		"Cardinality",
		"CanonicalFull",
		"CanonicalSimple",
		"CanonicalStem",
		"Authorship",
		"Year",
		"Quality",
	})
	return strings.Join(header, ",")
}

func (so *simple) ToSlice() []string {
	yr := so.Year
	if yr == "0" {
		yr = ""
	}

	qual := strconv.Itoa(so.Quality)
	card := strconv.Itoa(so.Cardinality)
	res := []string{
		so.ID,
		so.Verbatim,
		card,
		so.CanonicalRanked,
		so.Canonical,
		so.CanonicalStem,
		so.Authorship,
		yr,
		qual,
	}
	return res
}
