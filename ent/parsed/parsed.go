// Package parsed provides a user-friendly output of parsing result, as well
// as functions to convert the result to CSV or JSON-encoded strings.
package parsed

import (
	tb "github.com/gnames/tribool"
)

// Parsed is the result of a scientific name-string parsing. It can
// be converted into JSON or CSV formats.
type Parsed struct {
	// Parsed is false if parsing did not succeed.
	Parsed bool `json:"parsed"`

	// NomCode modifies parsing rules according to provided nomenclatural code.
	NomCode string `json:"nomenclaturalCode,omitempty"`

	// ParseQuality is a number that represents the quality of the
	// parsing.
	//
	//  0 - name-string is not parseable
	//  1 - no parsing problems encountered
	//  2 - small parsing problems
	//  3 - serious parsing problems
	//  4 - severe problems, name could not be parsed completely
	//
	// The ParseQuality is equal to the quality of the most
	// severe warning (see qualityWarnings). If no problems
	// are encountered, and the parsing succeeded, the parseQuality
	// is set to 1. If parsing failed, the parseQuality is 0.
	ParseQuality int `json:"quality"`

	// QualityWarnings contains encountered parsing problems.
	QualityWarnings []QualityWarning `json:"qualityWarnings,omitempty"`

	// Verbatim is input name-string without modifications.
	Verbatim string `json:"verbatim"`

	// Normalized is a normalized version of the input name-string.
	Normalized string `json:"normalized,omitempty"`

	// Canonical are simplified versions of a name-string more suitable for
	// matching and comparing name-strings than the verbatim version.
	Canonical *Canonical `json:"canonical,omitempty"`

	// Cardinality allows to sort, partition names according to number of
	// elements in their canonical forms.
	//
	// 0 - cardinality cannot be calculated
	// 1 - uninomial
	// 2 - binomial
	// 3 - trinomial
	// 4 - quadrinomial
	Cardinality int `json:"cardinality"`

	// Rank provides information about the rank of the name. It is not
	// always possible to infer rank correctly, so this field will be
	// omitted when the data for it does not exist.
	Rank string `json:"rank,omitempty"`

	// Authorship describes provided metainformation about authors of a name.
	// This authorship provided outside of Details belongs to
	// the most fine-grained element of a name.
	Authorship *Authorship `json:"authorship,omitempty"`

	// Bacteria is not nil if the input name has a genus
	// that is registered as bacterial. Possible
	// values are "maybe" - if the genus has homonyms in other groups
	// and "yes" if GNparser dictionary does not detect any homonyms
	//
	// The bacterial names often contain strain information which are
	// not parseable and are placed into the "tail" field.
	Bacteria *tb.Tribool `json:"bacteria,omitempty"`

	// Candidatus indicates that the parsed string is a candidatus bacterial name.
	Candidatus bool `json:"candidatus,omitempty"`

	// Virus is set to true in case if name is not parsed, and probably
	// belongs to a wide variety of sub-cellular entities like
	//
	// - viruses
	// - plasmids
	// - prions
	// - RNA
	// - DNA
	//
	// Viruses are the vast majority in this group of names,
	// as a result they gave (very imprecise) name to
	// the field.
	//
	// We do plan to create a parser for viruses at some point,
	// which will expand this group into more precise categories.
	Virus bool `json:"virus,omitempty"`

	// Cultivar is true if a name was parsed as a cultivar.
	Cultivar bool `json:"cultivar,omitempty"`

	// DaggerChar if true if a name-string includes '†' rune.
	// This rune might mean a fossil, or be indication of the clade extinction.
	DaggerChar bool `json:"daggerChar,omitempty"`

	// Hybrid is not nil if a name is detected as one of the hybrids
	//
	// - a non-categorized hybrid
	// - named hybrid
	// - notho- hybrid
	// - hybrid formula
	Hybrid *Annotation `json:"hybrid,omitempty"`

	// GraftChimera is not nil if a name is detected as one of the graft chimeras
	//
	// - a non-categorized graft chimera
	// - named graft chimera
	// - graft chimera formula
	GraftChimera *Annotation `json:"graftchimera,omitempty"`

	// Surrogate is a wide category of names that do not follow
	// nomenclatural rules

	// - a non-categorized surrogates
	// - surrogate names from BOLD project
	// - comparisons (Homo cf. sapiens)
	// - approximations (names for specimen that not fully identified)
	Surrogate *Annotation `json:"surrogate,omitempty"`

	// Tail is an unparseable tail of a name. It might contain "junk",
	// annotations, malformed parts of a scientific name, taxonomic concept
	// indications, bacterial strains etc.  If there is an unparseable tail, the
	// quality of the name-parsing is set to the worst category.
	Tail string `json:"tail,omitempty"`

	// Details contain more fine-grained information about parsed name.
	Details Details `json:"details,omitempty"`

	// Words contain description of every parsed word of a name.
	Words []Word `json:"words,omitempty"`

	// VerbatimID is a UUID v5 generated from the verbatim value of the
	// input name-string. Every unique string always generates the same
	// UUID.
	VerbatimID string `json:"id"`

	// ParserVersion is the version number of the GNparser.
	ParserVersion string `json:"parserVersion"`
}

// ParsedFlat is the result of a scientific name-string parsing flattened
// for the convenience.
type ParsedFlat struct {
	// Parsed is false if parsing did not succeed.
	Parsed bool `json:"parsed"`

	// NomCode modifies parsing rules according to provided nomenclatural code.
	NomCode string `json:"nomenclaturalCode,omitempty"`

	// ParseQuality is a number that represents the quality of the
	// parsing.
	//
	//  0 - name-string is not parseable
	//  1 - no parsing problems encountered
	//  2 - small parsing problems
	//  3 - serious parsing problems
	//  4 - severe problems, name could not be parsed completely
	//
	// The ParseQuality is equal to the quality of the most
	// severe warning (see qualityWarnings). If no problems
	// are encountered, and the parsing succeeded, the parseQuality
	// is set to 1. If parsing failed, the parseQuality is 0.
	ParseQuality int `json:"quality"`

	// Verbatim is input name-string without modifications.
	Verbatim string `json:"verbatim"`

	// Normalized is a normalized version of the input name-string.
	Normalized string `json:"normalized,omitempty"`

	// Canonical are simplified versions of a name-string more suitable for
	// matching and comparing name-strings than the verbatim version.
	CanonicalSimple string `json:"canonical,omitempty"`

	// Cardinality allows to sort, partition names according to number of
	// elements in their canonical forms.
	//
	// 0 - cardinality cannot be calculated
	// 1 - uninomial
	// 2 - binomial
	// 3 - trinomial
	// 4 - quadrinomial
	Cardinality int `json:"cardinality"`

	// Rank provides information about the rank of the name. It is not
	// always possible to infer rank correctly, so this field will be
	// omitted when the data for it does not exist.
	Rank string `json:"rank,omitempty"`

	// Authorship is the verbatim authorship of the name.
	Authorship string `json:"authorship,omitempty"`

	// Bacteria is not nil if the input name has a genus
	// that is registered as bacterial. Possible
	// values are "maybe" - if the genus has homonyms in other groups
	// and "yes" if GNparser dictionary does not detect any homonyms
	//
	// The bacterial names often contain strain information which are
	// not parseable and are placed into the "tail" field.
	Bacteria *tb.Tribool `json:"bacteria,omitempty"`

	// Candidatus indicates that the parsed string is a candidatus bacterial name.
	Candidatus bool `json:"candidatus,omitempty"`

	// Virus is set to true in case if name is not parsed, and probably
	// belongs to a wide variety of sub-cellular entities like
	//
	// - viruses
	// - plasmids
	// - prions
	// - RNA
	// - DNA
	//
	// Viruses are the vast majority in this group of names,
	// as a result they gave (very imprecise) name to
	// the field.
	//
	// We do plan to create a parser for viruses at some point,
	// which will expand this group into more precise categories.
	Virus bool `json:"virus,omitempty"`

	// Cultivar is true if a name was parsed as a cultivar.
	Cultivar bool `json:"cultivar,omitempty"`

	// DaggerChar if true if a name-string includes '†' rune.
	// This rune might mean a fossil, or be indication of the clade extinction.
	DaggerChar bool `json:"daggerChar,omitempty"`

	// Hybrid is a string representation of a hybrid type.
	//
	// - a non-categorized hybrid
	// - named hybrid
	// - notho- hybrid
	// - hybrid formula
	Hybrid string `json:"hybrid,omitempty"`

	// GraftChimera is a string representation of graft chimera.
	//
	// - a non-categorized graft chimera
	// - named graft chimera
	// - graft chimera formula
	GraftChimera string `json:"graftchimera,omitempty"`

	// Surrogate is a string repsresentation of a surrogate type.

	// - a non-categorized surrogates
	// - surrogate names from BOLD project
	// - comparisons (Homo cf. sapiens)
	// - approximations (names for specimen that not fully identified)
	Surrogate string `json:"surrogate,omitempty"`

	// Tail is an unparseable tail of a name. It might contain "junk",
	// annotations, malformed parts of a scientific name, taxonomic concept
	// indications, bacterial strains etc.  If there is an unparseable tail, the
	// quality of the name-parsing is set to the worst category.
	Tail string `json:"tail,omitempty"`

	// Uninomial represents the single name used for uninomial nomenclature,
	// typically applied to higher taxonomic ranks (e.g., family or order names
	// like "Asteraceae"). This field is populated only for uninomial names and
	// omitted otherwise.
	Uninomial string `json:"uninomial,omitempty"`

	// Genus specifies the genus part of a binomial or trinomial scientific name
	// (e.g., "Quercus" in "Quercus robur"). This field is empty if the name is
	// uninomial.
	Genus string `json:"genus,omitempty"`

	// InfragenericEpithet indicates the infrageneric epithet when present.
	// This field is omitted if not applicable.
	InfragenericEpithet string `json:"infragenericEpithet,omitempty"`

	// CultivarEpithet contains the cultivar name for cultivated plant varieties
	// (e.g., "Golden Delicious" in "Malus domestica 'Golden Delicious'"). This
	// field is populated only for names that include a cultivar designation.
	CultivarEpithet string `json:"cultivarEpithet,omitempty"`

	// Notho denotes the hybrid status of a name, indicating whether it is a
	// hybrid (e.g., "nothosubsp." or "nothovar." in "Salvia × sylvestris"). This
	// field is empty if not given.
	Notho string `json:"notho,omitempty"`

	// CombinationAuthorship provides the authorship for the current combination
	// of the name, typically the authors who transferred the species to a new
	// genus. (e.g., "K." in "Aus bus (L.) K."). This field is
	// omitted if no combination authorship is specified.
	CombinationAuthorship string `json:"combinationAuthorship,omitempty"`

	// CombinationExAuthorship captures the "ex" part of the combination
	// authorship (e.g., "ex DC." in "Quercus robur L. ex DC."). This field is
	// empty if no "ex" authorship exists.
	CombinationExAuthorship string `json:"combinationExAuthorship,omitempty"`

	// CombinationAuthorshipYear records the year associated with the combination
	// authorship, if provided (e.g., "1753" in "Homo sapiens (L.) K. 1753").
	// This field is omitted if the year is not specified.
	CombinationAuthorshipYear string `json:"combinationAuthorshipYear,omitempty"`

	// BasionymAuthorship identifies the authorship of the original combination
	// of the name (e.g., "Mill." in "Quercus robur (Mill.) L." where Mill. is
	// the original author). This field is populated only if basionym authorship
	// is present.
	BasionymAuthorship string `json:"basionymAuthorship,omitempty"`

	// BasionymExAuthorship specifies the "ex" part of the basionym authorship,
	// if applicable (e.g., "ex Torr." in "Pinus ponderosa Douglas ex Torr.").
	// This field is empty when no "ex" basionym authorship is provided.
	BasionymExAuthorship string `json:"basionymExAuthorship,omitempty"`

	// BasionymAuthorshipYear indicates the year tied to the basionym authorship
	// (e.g., "1820" in "Pinus ponderosa Douglas, 1820"). This field is included
	// only when the basionym year is explicitly stated.
	BasionymAuthorshipYear string `json:"basionymAuthorshipYear,omitempty"`

	// VerbatimID is a UUID v5 generated from the verbatim value of the
	// input name-string. Every unique string always generates the same
	// UUID.
	VerbatimID string `json:"id"`

	// ParserVersion is the version number of the GNparser.
	ParserVersion string `json:"parserVersion"`
}

// Canonical are simplified forms of a name-string more suitable for
// matching and comparing name-strings than the verbatim version.
type Canonical struct {
	// Stemmed is the most "normalized" and simplified version of the name.
	// Species epithets are stripped of suffixes, "j" character converted to "i",
	// "v" character converted to "u" according to "Schinke R, Greengrass M,
	// Robertson AM and Willett P (1996)"
	//
	// It is most useful to match names when a variability in suffixes is
	// possible.
	Stemmed string `json:"stemmed"`
	// Simple is a simplified version of a name where some elements like ranks,
	// or hybrid signs "×" are omitted (hybrid signs are present for hybrid
	// formulas).
	//
	// It is most useful to match names in general.
	Simple string `json:"simple"`
	// Full is a canonical form that keeps hybrid signs "×" for named
	// hybrids and shows infra-specific ranks.
	//
	// It is most useful for detection of the best matches from
	// multiple results. It is also recommended for displaying
	// canonical forms of botanical names.
	Full string `json:"full"`
}

// Authorship describes provided metainformation about authors of a name.
// Sometimes authorship is provided for several elements of a name, for example
// in "Agalinis purpurea (L.) Briton var. borealis (Berg.) Peterson 1987"
//
// The authorship provided outside of "details" section belongs to the most
// fine-grained element of a name ("var. borealis" for the example above).
type Authorship struct {
	// Verbatim is an authorship string without modifications.
	Verbatim string `json:"verbatim"`
	// Normalized is a normalized value of the authorship.
	Normalized string `json:"normalized"`
	// Year is a string representing a year of original description of the name.
	// The year number is surrounded by parentheses "(1758)", in cases when a
	// year is approximate.
	Year string `json:"year,omitempty"`
	// Authors is a slice containing each author as an element.
	Authors []string `json:"authors,omitempty"`
	// Original is an AuthGroup that contains authors of the original
	// description of a name.
	Original *AuthGroup `json:"originalAuth,omitempty"`
	// Combination is an AuthGroup that contains authors of new combination,
	// rank etc.
	Combination *AuthGroup `json:"combinationAuth,omitempty"`
}

// AuthGroup are provided only if config.WithDetails is true. Group of
// authors belonging to a particular nomenclatural event.  We distinguish two
// possible situations when AuthGroup is used:
//
// - original - authors of the original description of a name:w
// - combination - authors of a new combination, rank etc.
type AuthGroup struct {
	// Authors is a slice of strings containing found outhors
	Authors []string `json:"authors"`
	// Year provided only if "with_details=true" Year of the original
	// publication. If a range of the years provided, the start year is kept,
	// with isApproximate flag set to true.
	Year *Year `json:"year,omitempty"`
	// ExAuthors provided only if "with_details=true" A "special" group of
	// authors, that sometimes appear in scientific names after "ex"
	// qualifier.
	ExAuthors *Authors `json:"exAuthors,omitempty"`
	// InAuthors provided only if "with_details=true" A "special" group of
	// authors, that sometimes appear in scientific names after "in"
	// qualifier.
	InAuthors *Authors `json:"inAuthors,omitempty"`
	// EmendAuthors provided only if "with_details=true" A "special" group of
	// authors, that sometimes appear in scientific names after "emend."
	// qualifier.
	EmendAuthors *Authors `json:"emendAuthors,omitempty"`
}

// Authors contains information about authors and a year of publication.
type Authors struct {
	// Authors is a slice of strings containing found outhors of an AuthGroup
	Authors []string `json:"authors"`
	// Year of publication by the AuthGroup.
	Year *Year `json:"year,omitempty"`
}

// Year provided only if "with_details=true" Year of the original
// publication. If a range of the years provided, the start year is kept,
// with isApproximate flag set to true.
type Year struct {
	// Value is a string value of a year.
	Value string `json:"year"`
	// IsApproximate is indication if the year was written as approximate.
	// Approximate year might be represented by a range of years, by
	// a question mark "188?", by parentheses "(1888)".
	IsApproximate bool `json:"isApproximate,omitempty"`
}
