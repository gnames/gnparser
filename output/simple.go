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
