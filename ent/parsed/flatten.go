package parsed

import "strings"

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

	// CanonicalSimple is a simplified version of a name where some elements like ranks,
	// or hybrid signs "×" are omitted (hybrid signs are present for hybrid
	// formulas).
	//
	// It is most useful to match names in general.
	CanonicalSimple string `json:"canonicalSimple,omitempty"`

	// CanonicalFull is a canonical form that keeps hybrid signs "×" for named
	// hybrids and shows infra-specific ranks.
	//
	// It is most useful for detection of the best matches from
	// multiple results. It is also recommended for displaying
	// canonical forms of botanical names.
	CanonicalFull string `json:"canonicalFull,omitempty"`

	// CanonicalStemmed is the most "normalized" and simplified version of the name.
	// Species epithets are stripped of suffixes, "j" character converted to "i",
	// "v" character converted to "u" according to "Schinke R, Greengrass M,
	// Robertson AM and Willett P (1996)"
	//
	// It is most useful to match names when a variability in suffixes is
	// possible.
	CanonicalStemmed string `json:"canonicalStemmed,omitempty"`

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

	// Authors is the list of all authors separated by pipe character.
	Authors string `json:"authors,omitempty"`

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

	// Subgenus indicates the infrageneric epithet when present.
	// This field is omitted if not applicable.
	Subgenus string `json:"infragenericEpithet,omitempty"`

	// Species is the specific epithet of a binomial or trinomial.
	Species string `json:"specificEpithet,omitempty"`

	// Infraspecies is the infraspecificEpither of trinomials (names with
	// cardinality 3). We do not provide details for names with higher
	// cardinality.
	Infraspecies string `json:"infraspecificEpithet,omitempty"`

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
	// authorship, if provided (e.g., "1754" in "Homo sapiens (L.) K. 1753").
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

// Flatten converts a Parsed struct into a ParsedFlat struct, which is a
// flattened representation of the parsed data.
func (p Parsed) Flatten() ParsedFlat {
	var hybrid string
	if p.Hybrid != nil {
		hybrid = p.Hybrid.String()
	}
	var graft string
	if p.GraftChimera != nil {
		graft = p.GraftChimera.String()
	}
	var surrogate string
	if p.Surrogate != nil {
		surrogate = p.Surrogate.String()
	}

	res := ParsedFlat{
		Parsed:        p.Parsed,
		NomCode:       p.NomCode,
		ParseQuality:  p.ParseQuality,
		Verbatim:      p.Verbatim,
		Normalized:    p.Normalized,
		Cardinality:   p.Cardinality,
		Rank:          p.Rank,
		Candidatus:    p.Candidatus,
		Virus:         p.Virus,
		Cultivar:      p.Cultivar,
		DaggerChar:    p.DaggerChar,
		Hybrid:        hybrid,
		GraftChimera:  graft,
		Surrogate:     surrogate,
		Tail:          p.Tail,
		VerbatimID:    p.VerbatimID,
		ParserVersion: p.ParserVersion,
	}
	if !p.Parsed {
		return res
	}

	res.CanonicalSimple = p.Canonical.Simple
	res.CanonicalFull = p.Canonical.Full
	res.CanonicalStemmed = p.Canonical.Stemmed

	if p.Authorship != nil {
		au := p.Authorship
		res.Authorship = au.Verbatim
		res.Authors = strings.Join(au.Authors, "|")

		if au.Original != nil {
			res.BasionymAuthorship = authorship(au.Original)
			res.BasionymExAuthorship = exAuthorship(au.Original)
			res.BasionymAuthorshipYear = year(au.Original)
		}

		if au.Combination != nil {
			res.CombinationAuthorship = authorship(au.Combination)
			res.CombinationExAuthorship = exAuthorship(au.Combination)
			res.CombinationAuthorshipYear = year(au.Combination)
		}
	}

	switch detail := p.Details.(type) {
	case DetailsUninomial:
		res.Uninomial = detail.Uninomial.Value
	case DetailsSpecies:
		res.Genus = detail.Species.Genus
		res.Subgenus = detail.Species.Subgenus
		res.Species = detail.Species.Species
	case DetailsInfraspecies:
		if len(detail.Infraspecies.Infraspecies) == 1 {
			res.Genus = detail.Infraspecies.Genus
			res.Species = detail.Infraspecies.Species.Species
			res.Rank = detail.Infraspecies.Infraspecies[0].Rank
			res.Infraspecies = detail.Infraspecies.Infraspecies[0].Value
		}
	}
	return res
}

func authorship(ag *AuthGroup) string {
	if ag == nil {
		return ""
	}
	return joinAuthors(ag.Authors)
}

func joinAuthors(aus []string) string {
	var res string
	switch len(aus) {
	case 0:
		res = ""
	case 1:
		res = aus[0]
	case 2:
		res = strings.Join(aus, " & ")
	default:
		res = strings.Join(aus[0:len(aus)-1], ", ")
		res = res + " & " + aus[len(aus)-1]
	}
	return res
}

func exAuthorship(ag *AuthGroup) string {
	if ag == nil || ag.ExAuthors == nil {
		return ""
	}
	return joinAuthors(ag.ExAuthors.Authors)
}

func year(ag *AuthGroup) string {
	if ag == nil || ag.Year == nil {
		return ""
	}
	if ag.Year.IsApproximate {
		return "(" + ag.Year.Value + ")"
	}
	return ag.Year.Value
}
