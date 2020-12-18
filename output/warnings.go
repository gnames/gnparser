package output

import (
	"sort"

	grm "github.com/gnames/gnparser/grammar"
	jsoniter "github.com/json-iterator/go"
)

type Warning struct {
	Quality int
	Message string
}

var warningMap = map[grm.Warning]Warning{
	grm.TailWarn: {
		Quality: 3,
		Message: "Unparsed tail",
	},
	grm.ApostrOtherWarn: {
		Quality: 3,
		Message: "Not an ASCII apostrophe",
	},
	grm.AuthAmbiguousFiliusWarn: {
		Quality: 2,
		Message: "Ambiguous f. (filius or forma)",
	},
	grm.AuthDoubleParensWarn: {
		Quality: 3,
		Message: "Authorship in double parentheses",
	},
	grm.AuthExWarn: {
		Quality: 2,
		Message: "Ex authors are not required",
	},
	grm.AuthExWithDotWarn: {
		Quality: 3,
		Message: "`ex` ends with dot",
	},
	grm.AuthEmendWarn: {
		Quality: 2,
		Message: "Emend authors are not required",
	},
	grm.AuthEmendWithoutDotWarn: {
		Quality: 3,
		Message: "`emend` without a period",
	},
	grm.AuthMissingOneParensWarn: {
		Quality: 3,
		Message: "Authorship is missing one parenthesis",
	},
	grm.AuthQuestionWarn: {
		Quality: 3,
		Message: "Author as a question mark",
	},
	grm.AuthShortWarn: {
		Quality: 3,
		Message: "Author is too short",
	},
	grm.AuthUnknownWarn: {
		Quality: 2,
		Message: "Author is unknown",
	},
	grm.AuthUpperCaseWarn: {
		Quality: 2,
		Message: "Author in upper case",
	},
	grm.BacteriaMaybeWarn: {
		Quality: 1,
		Message: "The genus is a homonym of a bacterial genus",
	},
	grm.BotanyAuthorNotSubgenWarn: {
		Quality: 2,
		Message: "Possible ICN author instead of subgenus",
	},
	grm.CanonicalApostropheWarn: {
		Quality: 3,
		Message: "Apostrophe is not allowed in canonical",
	},
	grm.CapWordQuestionWarn: {
		Quality: 3,
		Message: "Uninomial word with question mark",
	},
	grm.CharBadWarn: {
		Quality: 2,
		Message: "Non-standard characters in canonical",
	},
	grm.GenusAbbrWarn: {
		Quality: 3,
		Message: "Abbreviated uninomial word",
	},
	grm.GenusUpperCharAfterDash: {
		Quality: 2,
		Message: "Apparent genus with capital character after hyphen",
	},
	grm.GreekLetterInRank: {
		Quality: 2,
		Message: "Deprecated Greek letter enumeration in rank",
	},
	grm.HTMLTagsEntitiesWarn: {
		Quality: 3,
		Message: "HTML tags or entities in the name",
	},
	grm.HybridCharNoSpaceWarn: {
		Quality: 3,
		Message: "Hybrid char not separated by space",
	},
	grm.HybridFormulaWarn: {
		Quality: 2,
		Message: "Hybrid formula",
	},
	grm.HybridFormulaIncompleteWarn: {
		Quality: 3,
		Message: "Incomplete hybrid formula",
	},
	grm.HybridFormulaProbIncompleteWarn: {
		Quality: 2,
		Message: "Probably incomplete hybrid formula",
	},
	grm.HybridNamedWarn: {
		Quality: 2,
		Message: "Named hybrid",
	},
	grm.NameApproxWarn: {
		Quality: 3,
		Message: "Name is approximate",
	},
	grm.NameComparisonWarn: {
		Quality: 3,
		Message: "Name comparison",
	},
	grm.RankUncommonWarn: {
		Quality: 3,
		Message: "Uncommon rank",
	},
	grm.SpaceMultipleWarn: {
		Quality: 2,
		Message: "Multiple adjacent space characters",
	},
	grm.SpaceNonStandardWarn: {
		Quality: 3,
		Message: "Non-standard space characters",
	},
	grm.SpanishAndAsSeparator: {
		Quality: 2,
		Message: "Spanish 'y' is used instead of '&'",
	},
	grm.SpeciesNumericWarn: {
		Quality: 3,
		Message: "Numeric prefix",
	},
	grm.SuperSpeciesWarn: {
		Quality: 2,
		Message: "Ambiguity: subgenus or superspecies found",
	},
	grm.UTF8ConvBadWarn: {
		Quality: 3,
		Message: "Incorrect conversion to UTF-8",
	},
	grm.UninomialComboWarn: {
		Quality: 2,
		Message: "Combination of two uninomials",
	},
	grm.WhiteSpaceTrailWarn: {
		Quality: 2,
		Message: "Trailing whitespace",
	},
	grm.YearCharWarn: {
		Quality: 2,
		Message: "Year with latin character",
	},
	grm.YearDotWarn: {
		Quality: 2,
		Message: "Year with period",
	},
	grm.YearOrigMisplacedWarn: {
		Quality: 2,
		Message: "Misplaced basionym year",
	},
	grm.YearPageWarn: {
		Quality: 3,
		Message: "Year with page info",
	},
	grm.YearParensWarn: {
		Quality: 2,
		Message: "Year with parentheses",
	},
	grm.YearQuestionWarn: {
		Quality: 2,
		Message: "Year with question mark",
	},
	grm.YearRangeWarn: {
		Quality: 3,
		Message: "Years range",
	},
	grm.YearSqBraketsWarn: {
		Quality: 3,
		Message: "Year with square brakets",
	},
}

func prepareWarnings(ws []grm.Warning) []Warning {
	res := make([]Warning, len(ws))
	for i, v := range ws {
		res[i] = warningMap[v]
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Quality > res[j].Quality {
			return true
		}
		if res[i].Quality < res[j].Quality {
			return false
		}
		return res[i].Message < res[j].Message
	})
	return res
}

func (w *Warning) MarshalJSON() ([]byte, error) {
	arr := []interface{}{w.Quality, w.Message}
	return jsoniter.Marshal(arr)
}

func (w *Warning) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	_ = jsoniter.Unmarshal(bs, &arr)
	w.Quality = int(arr[0].(float64))
	w.Message = arr[1].(string)
	return nil
}
