package output

import (
	"errors"
	"strings"
)

type Warning int

const (
	TailWarn Warning = iota
	ApostrOtherWarn
	AuthAmbiguousFiliusWarn
	AuthDoubleParensWarn
	AuthExWarn
	AuthExWithDotWarn
	AuthEmendWarn
	AuthEmendWithoutDotWarn
	AuthMissingOneParensWarn
	AuthQuestionWarn
	AuthShortWarn
	AuthUnknownWarn
	AuthUpperCaseWarn
	BacteriaMaybeWarn
	BotanyAuthorNotSubgenWarn
	CanonicalApostropheWarn
	CapWordQuestionWarn
	CharBadWarn
	GenusAbbrWarn
	GenusUpperCharAfterDash
	GreekLetterInRank
	HTMLTagsEntitiesWarn
	HybridCharNoSpaceWarn
	HybridFormulaWarn
	HybridFormulaIncompleteWarn
	HybridFormulaProbIncompleteWarn
	HybridNamedWarn
	NameApproxWarn
	NameComparisonWarn
	RankUncommonWarn
	SpaceMultipleWarn
	SpaceNonStandardWarn
	SpanishAndAsSeparator
	SpeciesNumericWarn
	SuperSpeciesWarn
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
	YearSqBraketsWarn
)

var warningMap = map[Warning]string{
	TailWarn:                        "Unparsed tail",
	ApostrOtherWarn:                 "Not an ASCII apostrophe",
	AuthAmbiguousFiliusWarn:         "Ambiguous f. (filius or forma)",
	AuthDoubleParensWarn:            "Authorship in double parentheses",
	AuthExWarn:                      "Ex authors are not required",
	AuthExWithDotWarn:               "`ex` ends with a dot",
	AuthEmendWarn:                   "Emend authors are not required",
	AuthEmendWithoutDotWarn:         "`emend` without a period",
	AuthMissingOneParensWarn:        "Authorship is missing one parenthesis",
	AuthQuestionWarn:                "Author as a question mark",
	AuthShortWarn:                   "Author is too short",
	AuthUnknownWarn:                 "Author is unknown",
	AuthUpperCaseWarn:               "Author in upper case",
	BacteriaMaybeWarn:               "The genus is a homonym of a bacterial genus",
	BotanyAuthorNotSubgenWarn:       "Possible ICN author instead of subgenus",
	CanonicalApostropheWarn:         "Apostrophe is not allowed in canonical",
	CapWordQuestionWarn:             "Uninomial word with question mark",
	CharBadWarn:                     "Non-standard characters in canonical",
	GenusAbbrWarn:                   "Abbreviated uninomial word",
	GenusUpperCharAfterDash:         "Apparent genus with capital character after hyphen",
	GreekLetterInRank:               "Deprecated Greek letter enumeration in rank",
	HTMLTagsEntitiesWarn:            "HTML tags or entities in the name",
	HybridCharNoSpaceWarn:           "Hybrid char not separated by space",
	HybridFormulaWarn:               "Hybrid formula",
	HybridFormulaIncompleteWarn:     "Incomplete hybrid formula",
	HybridFormulaProbIncompleteWarn: "Probably incomplete hybrid formula",
	HybridNamedWarn:                 "Named hybrid",
	NameApproxWarn:                  "Name is approximate",
	NameComparisonWarn:              "Name comparison",
	RankUncommonWarn:                "Uncommon rank",
	SpaceMultipleWarn:               "Multiple adjacent space characters",
	SpaceNonStandardWarn:            "Non-standard space characters",
	SpanishAndAsSeparator:           "Spanish 'y' is used instead of '&'",
	SpeciesNumericWarn:              "Numeric prefix",
	SuperSpeciesWarn:                "Ambiguity: subgenus or superspecies found",
	UTF8ConvBadWarn:                 "Incorrect conversion to UTF-8",
	UninomialComboWarn:              "Combination of two uninomials",
	WhiteSpaceTrailWarn:             "Trailing whitespace",
	YearCharWarn:                    "Year with latin character",
	YearDotWarn:                     "Year with period",
	YearOrigMisplacedWarn:           "Misplaced basionym year",
	YearPageWarn:                    "Year with page info",
	YearParensWarn:                  "Year with parentheses",
	YearQuestionWarn:                "Year with question mark",
	YearRangeWarn:                   "Years range",
	YearSqBraketsWarn:               "Year with square brakets",
}

var warningStrMap = func() map[string]Warning {
	res := make(map[string]Warning)
	for k, v := range warningMap {
		res[v] = k
	}
	return res
}()

var warningQualityMap = map[Warning]int{
	TailWarn:                        4,
	ApostrOtherWarn:                 3,
	AuthAmbiguousFiliusWarn:         2,
	AuthDoubleParensWarn:            4,
	AuthExWarn:                      2,
	AuthExWithDotWarn:               3,
	AuthEmendWarn:                   2,
	AuthEmendWithoutDotWarn:         3,
	AuthMissingOneParensWarn:        4,
	AuthQuestionWarn:                4,
	AuthShortWarn:                   3,
	AuthUnknownWarn:                 2,
	AuthUpperCaseWarn:               2,
	BacteriaMaybeWarn:               1,
	BotanyAuthorNotSubgenWarn:       2,
	CanonicalApostropheWarn:         3,
	CapWordQuestionWarn:             4,
	CharBadWarn:                     2,
	GenusAbbrWarn:                   4,
	GenusUpperCharAfterDash:         2,
	GreekLetterInRank:               2,
	HTMLTagsEntitiesWarn:            3,
	HybridCharNoSpaceWarn:           3,
	HybridFormulaWarn:               2,
	HybridFormulaIncompleteWarn:     4,
	HybridFormulaProbIncompleteWarn: 2,
	HybridNamedWarn:                 2,
	NameApproxWarn:                  4,
	NameComparisonWarn:              4,
	RankUncommonWarn:                3,
	SpaceMultipleWarn:               2,
	SpaceNonStandardWarn:            2,
	SpanishAndAsSeparator:           2,
	SpeciesNumericWarn:              3,
	SuperSpeciesWarn:                2,
	UTF8ConvBadWarn:                 4,
	UninomialComboWarn:              2,
	WhiteSpaceTrailWarn:             2,
	YearCharWarn:                    2,
	YearDotWarn:                     2,
	YearOrigMisplacedWarn:           2,
	YearPageWarn:                    2,
	YearParensWarn:                  2,
	YearQuestionWarn:                2,
	YearRangeWarn:                   3,
	YearSqBraketsWarn:               3,
}

type QualityWarning struct {
	Quality int     `json:"quality"`
	Warning Warning `json:"warning"`
}

func (w Warning) String() string {
	return warningMap[w]
}

func (w Warning) Quality() int {
	return warningQualityMap[w]
}

func (w Warning) NewQualityWarning() QualityWarning {
	return QualityWarning{
		Quality: w.Quality(),
		Warning: w,
	}
}

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
