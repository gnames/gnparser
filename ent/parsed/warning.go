package parsed

import (
	"errors"
	"strings"
)

// Warning is a type to represent warnings found during parsing of a
// scientific name.
type Warning int

const (
	TailWarn Warning = iota
	ApostrOtherWarn
	AuthAmbiguousFiliusWarn
	AuthDoubleParensWarn
	AuthEmendWarn
	AuthEmendWithoutDotWarn
	AuthExWarn
	AuthExWithDotWarn
	AuthMissingOneParensWarn
	AuthQuestionWarn
	AuthShortWarn
	AuthUnknownWarn
	AuthUpperCaseWarn
	BacteriaMaybeWarn
	BotanyAuthorNotSubgenWarn
	CandidatusName
	CanonicalApostropheWarn
	CapWordQuestionWarn
	CharBadWarn
	ContainsIgnoredAnnotation
	CultivarEpithetWarn
	DashOtherWarn
	DotEpithetWarn
	GenusAbbrWarn
	GenusUpperCharAfterDash
	GraftChimeraCharNoSpaceWarn
	GraftChimeraFormulaIncompleteWarn
	GraftChimeraFormulaProbIncompleteWarn
	GraftChimeraFormulaWarn
	GraftChimeraNamedWarn
	GreekLetterInRank
	HTMLTagsEntitiesWarn
	HybridCharNoSpaceWarn
	HybridFormulaIncompleteWarn
	HybridFormulaProbIncompleteWarn
	HybridFormulaWarn
	HybridNamedWarn
	LowCaseWarn
	NameApproxWarn
	NameComparisonWarn
	RankUncommonWarn
	SpaceNonStandardWarn
	SpanishAndAsSeparator
	SpeciesNumericWarn
	SubgenusAbbrWarn
	SuperspeciesWarn
	UTF8ConvBadWarn
	UninomialComboWarn
	WhiteSpaceTrailWarn
	YearCharWarn
	YearDotWarn
	YearOrigMisplacedWarn
	YearPageWarn
	YearParensWarn
	YearQuestionWarn
	YearRangeWarn
	YearSqBracketsWarn
)

var warningMap = map[Warning]string{
	TailWarn:                              "Unparsed tail",
	ApostrOtherWarn:                       "Not an ASCII apostrophe",
	AuthAmbiguousFiliusWarn:               "Ambiguous f. (filius or forma)",
	AuthDoubleParensWarn:                  "Authorship in double parentheses",
	AuthEmendWarn:                         "Emend authors are not required",
	AuthEmendWithoutDotWarn:               "`emend` without a period",
	AuthExWarn:                            "Ex authors are not required (ICZN only)",
	AuthExWithDotWarn:                     "`ex` ends with a period",
	AuthMissingOneParensWarn:              "Authorship is missing one parenthesis",
	AuthQuestionWarn:                      "Author as a question mark",
	AuthShortWarn:                         "Author is too short",
	AuthUnknownWarn:                       "Author is unknown",
	AuthUpperCaseWarn:                     "Author in upper case",
	BacteriaMaybeWarn:                     "The genus is a homonym of a bacterial genus",
	BotanyAuthorNotSubgenWarn:             "Possible ICN author instead of subgenus",
	CandidatusName:                        "Bacterial `Candidatus` name",
	CanonicalApostropheWarn:               "Apostrophe is not allowed in canonical",
	CapWordQuestionWarn:                   "Uninomial word with question mark",
	CharBadWarn:                           "Non-standard characters in canonical",
	ContainsIgnoredAnnotation:             "Ignored annotation `mihi`",
	CultivarEpithetWarn:                   "Cultivar epithet",
	DashOtherWarn:                         "Atypical hyphen character",
	DotEpithetWarn:                        "Period character is not allowed in canonical",
	GenusAbbrWarn:                         "Abbreviated uninomial word",
	GenusUpperCharAfterDash:               "Apparent genus with capital character after hyphen",
	GraftChimeraCharNoSpaceWarn:           "Graft-chimera char is not separated by space",
	GraftChimeraFormulaIncompleteWarn:     "Incomplete graft-chimera formula",
	GraftChimeraFormulaProbIncompleteWarn: "Probably incomplete graft-chimera formula",
	GraftChimeraFormulaWarn:               "Graft-chimera formula",
	GraftChimeraNamedWarn:                 "Named graft-chimera",
	GreekLetterInRank:                     "Deprecated Greek letter enumeration in rank",
	HTMLTagsEntitiesWarn:                  "HTML tags or entities in the name",
	HybridCharNoSpaceWarn:                 "Hybrid char is not separated by space",
	HybridFormulaIncompleteWarn:           "Incomplete hybrid formula",
	HybridFormulaProbIncompleteWarn:       "Probably incomplete hybrid formula",
	HybridFormulaWarn:                     "Hybrid formula",
	HybridNamedWarn:                       "Named hybrid",
	LowCaseWarn:                           "Name starts with low-case character",
	NameApproxWarn:                        "Name is approximate",
	NameComparisonWarn:                    "Name comparison",
	RankUncommonWarn:                      "Uncommon rank",
	SpaceNonStandardWarn:                  "Non-standard space characters",
	SpanishAndAsSeparator:                 "Spanish 'y' is used instead of '&'",
	SpeciesNumericWarn:                    "Numeric prefix",
	SubgenusAbbrWarn:                      "Abbreviated subgenus",
	SuperspeciesWarn:                      "Ambiguity: subgenus or superspecies found",
	UTF8ConvBadWarn:                       "Incorrect conversion to UTF-8",
	UninomialComboWarn:                    "Combination of two uninomials",
	WhiteSpaceTrailWarn:                   "Trailing whitespace",
	YearCharWarn:                          "Year with latin character",
	YearDotWarn:                           "Year with period",
	YearOrigMisplacedWarn:                 "Misplaced basionym year",
	YearPageWarn:                          "Year with page info",
	YearParensWarn:                        "Year with parentheses",
	YearQuestionWarn:                      "Year with question mark",
	YearRangeWarn:                         "Years range",
	YearSqBracketsWarn:                    "Year with square brackets",
}

var warningStrMap = func() map[string]Warning {
	res := make(map[string]Warning)
	for k, v := range warningMap {
		res[v] = k
	}
	return res
}()

// WarningQualityMap assigns quality of parsing for each warning type.
var WarningQualityMap = map[Warning]int{
	TailWarn:                              4,
	ApostrOtherWarn:                       3,
	AuthAmbiguousFiliusWarn:               2,
	AuthDoubleParensWarn:                  4,
	AuthEmendWarn:                         2,
	AuthEmendWithoutDotWarn:               3,
	AuthExWarn:                            2,
	AuthExWithDotWarn:                     3,
	AuthMissingOneParensWarn:              4,
	AuthQuestionWarn:                      4,
	AuthShortWarn:                         3,
	AuthUnknownWarn:                       2,
	AuthUpperCaseWarn:                     2,
	BacteriaMaybeWarn:                     1,
	BotanyAuthorNotSubgenWarn:             2,
	CandidatusName:                        2,
	CanonicalApostropheWarn:               3,
	CapWordQuestionWarn:                   4,
	CharBadWarn:                           2,
	ContainsIgnoredAnnotation:             3,
	CultivarEpithetWarn:                   2,
	DashOtherWarn:                         2,
	DotEpithetWarn:                        3,
	GenusAbbrWarn:                         4,
	GenusUpperCharAfterDash:               2,
	GraftChimeraCharNoSpaceWarn:           3,
	GraftChimeraFormulaIncompleteWarn:     4,
	GraftChimeraFormulaProbIncompleteWarn: 2,
	GraftChimeraFormulaWarn:               2,
	GraftChimeraNamedWarn:                 2,
	GreekLetterInRank:                     2,
	HTMLTagsEntitiesWarn:                  3,
	HybridCharNoSpaceWarn:                 3,
	HybridFormulaIncompleteWarn:           4,
	HybridFormulaProbIncompleteWarn:       2,
	HybridFormulaWarn:                     2,
	HybridNamedWarn:                       2,
	LowCaseWarn:                           4,
	NameApproxWarn:                        4,
	NameComparisonWarn:                    4,
	RankUncommonWarn:                      3,
	SpaceNonStandardWarn:                  2,
	SpanishAndAsSeparator:                 2,
	SpeciesNumericWarn:                    3,
	SubgenusAbbrWarn:                      2,
	SuperspeciesWarn:                      2,
	UTF8ConvBadWarn:                       4,
	UninomialComboWarn:                    2,
	WhiteSpaceTrailWarn:                   2,
	YearCharWarn:                          2,
	YearDotWarn:                           2,
	YearOrigMisplacedWarn:                 2,
	YearPageWarn:                          2,
	YearParensWarn:                        2,
	YearQuestionWarn:                      2,
	YearRangeWarn:                         3,
	YearSqBracketsWarn:                    3,
}

// QualityWarning is and object that contains the warning and its
// corresponding quality.
type QualityWarning struct {
	Quality int     `json:"quality"`
	Warning Warning `json:"warning"`
}

// String implements fmt.Stringer interface.
func (w Warning) String() string {
	return warningMap[w]
}

// Quality returns parsing quality number that corresponds to a
// particular warning.
func (w Warning) Quality() int {
	return WarningQualityMap[w]
}

// NewQualityWarning creates new QualityWarning object.
func (w Warning) NewQualityWarning() QualityWarning {
	return QualityWarning{
		Quality: w.Quality(),
		Warning: w,
	}
}

// Map converts slice of warnings to a slice of QualityWarning structures.
func Map(ws []Warning) []QualityWarning {
	res := make([]QualityWarning, len(ws))
	for i, v := range ws {
		res[i] = v.NewQualityWarning()
	}
	return res
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int is null.
func (w Warning) MarshalJSON() ([]byte, error) {
	return []byte("\"" + w.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaller.
func (w *Warning) UnmarshalJSON(bs []byte) error {
	var err error
	var ok bool
	// strings.Trim seems to be ~10 time faster here than
	// json-iter Unmarshal
	s := strings.Trim(string(bs), `"`)
	*w, ok = warningStrMap[s]
	if !ok {
		err = errors.New("cannot decode Warning")
	}
	return err
}
