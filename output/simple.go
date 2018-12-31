package output

import (
	"fmt"
	"strconv"

	"gitlab.com/gogna/gnparser/grammar"
)

type simple struct {
	ID              string
	Verbatim        string
	Canonical       string
	CanonicalRanked string
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
	if ao != nil && len(ao.Original.Years) > 0 {
		yrs := ao.Original.Years
		yr = yrs[0].Value
		if yrs[0].Approximate {
			yr = fmt.Sprintf("(%s)", yr)
		}
	}

	c, _ := sn.Canonical()

	_, quality := qualityAndWarnings(sn.Warnings)
	so := simple{
		ID:              sn.VerbatimID,
		Verbatim:        sn.Verbatim,
		Canonical:       c.Value,
		CanonicalRanked: c.ValueRanked,
		Authorship:      authorship,
		Year:            yr,
		Quality:         quality,
	}
	return &so
}

func (so *simple) ToSlice() []string {
	yr := so.Year
	if yr == "0" {
		yr = ""
	}

	qual := strconv.Itoa(so.Quality)
	if qual == "0" {
		qual = ""
	}
	res := []string{
		so.ID,
		so.Verbatim,
		so.Canonical,
		so.CanonicalRanked,
		so.Authorship,
		yr,
		qual,
	}
	return res
}
