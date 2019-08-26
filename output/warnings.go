package output

import (
	"sort"

	jsoniter "github.com/json-iterator/go"
	grm "gitlab.com/gogna/gnparser/grammar"
)

type Warning struct {
	Quality int
	Message string
}

var warningMap = map[grm.Warning]Warning{
	grm.TailWarn: Warning{
		Quality: 3,
		Message: "Unparsed tail",
	},
	grm.ApostrOtherWarn: Warning{
		Quality: 3,
		Message: "Not an ASCII apostrophe",
	},
	grm.AuthAmbiguousFiliusWarn: Warning{
		Quality: 2,
		Message: "Ambiguous f. (filius or forma)",
	},
	grm.AuthDoubleParensWarn: Warning{
		Quality: 3,
		Message: "Authorship in double parentheses",
	},
	grm.AuthExWarn: Warning{
		Quality: 2,
		Message: "Ex authors are not required",
	},
	grm.AuthExWithDotWarn: Warning{
		Quality: 3,
		Message: "`ex` ends with dot",
	},
	grm.AuthEmendWarn: Warning{
		Quality: 2,
		Message: "Emend authors are not required",
	},
	grm.AuthEmendWithoutDotWarn: Warning{
		Quality: 3,
		Message: "`emend` without a period",
	},
	grm.AuthMissingOneParensWarn: Warning{
		Quality: 3,
		Message: "Authorship is missing one parenthesis",
	},
	grm.AuthQuestionWarn: Warning{
		Quality: 3,
		Message: "Author as a question mark",
	},
	grm.AuthShortWarn: Warning{
		Quality: 3,
		Message: "Author is too short",
	},
	grm.AuthUnknownWarn: Warning{
		Quality: 2,
		Message: "Author is unknown",
	},
	grm.AuthUpperCaseWarn: Warning{
		Quality: 2,
		Message: "Author in upper case",
	},
	grm.BacteriaMaybeWarn: Warning{
		Quality: 1,
		Message: "The genus is a homonym of a bacterial genus",
	},
	grm.BotanyAuthorNotSubgenWarn: Warning{
		Quality: 2,
		Message: "Possible ICN author instead of subgenus",
	},
	grm.CanonicalApostropheWarn: Warning{
		Quality: 3,
		Message: "Apostrophe is not allowed in canonical",
	},
	grm.CapWordQuestionWarn: Warning{
		Quality: 3,
		Message: "Uninomial word with question mark",
	},
	grm.CharBadWarn: Warning{
		Quality: 2,
		Message: "Non-standard characters in canonical",
	},
	grm.GenusAbbrWarn: Warning{
		Quality: 3,
		Message: "Abbreviated uninomial word",
	},
	grm.GenusUpperCharAfterDash: Warning{
		Quality: 2,
		Message: "Apparent genus with capital character after hyphen",
	},
	grm.GreekLetterInRank: Warning{
		Quality: 2,
		Message: "Deprecated Greek letter enumeration in rank",
	},
	grm.HybridCharNoSpaceWarn: Warning{
		Quality: 3,
		Message: "Hybrid char not separated by space",
	},
	grm.HybridFormulaWarn: Warning{
		Quality: 2,
		Message: "Hybrid formula",
	},
	grm.HybridFormulaIncompleteWarn: Warning{
		Quality: 3,
		Message: "Incomplete hybrid formula",
	},
	grm.HybridFormulaProbIncompleteWarn: Warning{
		Quality: 2,
		Message: "Probably incomplete hybrid formula",
	},
	grm.HybridNamedWarn: Warning{
		Quality: 2,
		Message: "Named hybrid",
	},
	grm.NameApproxWarn: Warning{
		Quality: 3,
		Message: "Name is approximate",
	},
	grm.NameComparisonWarn: Warning{
		Quality: 3,
		Message: "Name comparison",
	},
	grm.RankUncommonWarn: Warning{
		Quality: 3,
		Message: "Uncommon rank",
	},
	grm.SpaceMultipleWarn: Warning{
		Quality: 2,
		Message: "Multiple adjacent space characters",
	},
	grm.SpaceNonStandardWarn: Warning{
		Quality: 3,
		Message: "Non-standard space characters",
	},
	grm.SpeciesNumericWarn: Warning{
		Quality: 3,
		Message: "Numeric prefix",
	},
	grm.SuperSpeciesWarn: Warning{
		Quality: 2,
		Message: "Ambiguity: subgenus or superspecies found",
	},
	grm.UTF8ConvBadWarn: Warning{
		Quality: 3,
		Message: "Incorrect conversion to UTF-8",
	},
	grm.UninomialComboWarn: Warning{
		Quality: 2,
		Message: "Combination of two uninomials",
	},
	grm.WhiteSpaceTrailWarn: Warning{
		Quality: 2,
		Message: "Trailing whitespace",
	},
	grm.YearCharWarn: Warning{
		Quality: 2,
		Message: "Year with latin character",
	},
	grm.YearDotWarn: Warning{
		Quality: 2,
		Message: "Year with period",
	},
	grm.YearOrigMisplacedWarn: Warning{
		Quality: 2,
		Message: "Misplaced basionym year",
	},
	grm.YearPageWarn: Warning{
		Quality: 3,
		Message: "Year with page info",
	},
	grm.YearParensWarn: Warning{
		Quality: 2,
		Message: "Year with parentheses",
	},
	grm.YearQuestionWarn: Warning{
		Quality: 2,
		Message: "Year with question mark",
	},
	grm.YearRangeWarn: Warning{
		Quality: 3,
		Message: "Years range",
	},
	grm.YearSqBraketsWarn: Warning{
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
	jsoniter.Unmarshal(bs, &arr)
	// TODO: add error handling here.
	w.Quality = int(arr[0].(float64))
	w.Message = arr[1].(string)
	return nil
}
