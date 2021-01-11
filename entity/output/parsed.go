package output

import (
	tb "github.com/gnames/gnlib/tribool"
)

type Parsed struct {
	Parsed          bool             `json:"parsed"`
	ParseQuality    int              `json:"quality"`
	QualityWarnings []QualityWarning `json:"qualityWarnings,omitempty"`
	Verbatim        string           `json:"verbatim"`
	Normalized      string           `json:"normalized,omitempty"`
	Canonical       *Canonical       `json:"canonical,omitempty"`
	Cardinality     int              `json:"cardinality"`
	Authorship      *Authorship      `json:"authorship,omitempty"`
	Bacteria        *tb.Tribool      `json:"bacteria,omitempty"`
	Virus           bool             `json:"virus,omitempty"`
	Hybrid          *Annotation      `json:"hybrid,omitempty"`
	Surrogate       *Annotation      `json:"surrogate,omitempty"`
	Tail            string           `json:"tail,omitempty"`
	Details         Details          `json:"details,omitempty"`
	Positions       []Position       `json:"pos,omitempty"`
	VerbatimID      string           `json:"id"`
	ParserVersion   string           `json:"parserVersion"`
}

type Canonical struct {
	Stemmed string `json:"stemmed"`
	Simple  string `json:"simple"`
	Full    string `json:"full"`
}

type Authorship struct {
	Verbatim    string     `json:"verbatim"`
	Normalized  string     `json:"normalized"`
	Year        string     `json:"year,omitempty"`
	Authors     []string   `json:"authors,omitempty"`
	Original    *AuthGroup `json:"originalAuth,omitempty"`
	Combination *AuthGroup `json:"combinationAuth,omitempty"`
}

type AuthGroup struct {
	Authors      []string `json:"authors"`
	Year         *Year    `json:"year,omitempty"`
	ExAuthors    *Authors `json:"exAuthors,omitempty"`
	EmendAuthors *Authors `json:"emendAuthors,omitempty"`
}

type Authors struct {
	Authors []string `json:"authors"`
	Year    *Year    `json:"year,omitempty"`
}

type Year struct {
	Value         string `json:"year"`
	IsApproximate bool   `json:"isApproximate,omitempty"`
}

type Position struct {
	Type  WordType `json:"wordType"`
	Start int      `json:"start"`
	End   int      `json:"end"`
}
