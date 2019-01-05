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
	grm.SpaceMultipleWarn: Warning{
		Quality: 2,
		Message: "Multiple adjacent space characters",
	},
	grm.SpaceNonStandardWarn: Warning{
		Quality: 3,
		Message: "Non-standard space characters",
	},
	grm.UTF8ConvBadWarn: Warning{
		Quality: 3,
		Message: "Incorrect conversion to UTF-8",
	},
	grm.WhiteSpaceTrailWarn: Warning{
		Quality: 2,
		Message: "Trailing whitespace",
	},
	grm.TailWarn: Warning{
		Quality: 3,
		Message: "Unparseable tail",
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
	grm.HybridCharNoSpaceWarn: Warning{
		Quality: 3,
		Message: "Hybrid char not separated by space",
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
	grm.SuperSpeciesWarn: Warning{
		Quality: 2,
		Message: "Ambiguity: subgenus or superspecies found",
	},
	grm.UninomialComboWarn: Warning{
		Quality: 2,
		Message: "Combination of two uninomials",
	},
	grm.GenusAbbrWarn: Warning{
		Quality: 3,
		Message: "Abbreviated uninomial word",
	},
	grm.CharBadWarn: Warning{
		Quality: 2,
		Message: "Non-standard characters in canonical",
	},
	grm.CapWordQuestionWarn: Warning{
		Quality: 3,
		Message: "Uninomial word with question mark",
	},
	grm.CanonicalApostropheWarn: Warning{
		Quality: 3,
		Message: "Apostrophe is not allowed in canonical",
	},
	grm.SpeciesNumericWarn: Warning{
		Quality: 3,
		Message: "Numeric prefix",
	},
	grm.AuthExWarn: Warning{
		Quality: 2,
		Message: "Ex authors are not required",
	},
	grm.AuthEmendWarn: Warning{
		Quality: 2,
		Message: "Emend authors are not required",
	},
	grm.YearOrigMisplacedWarn: Warning{
		Quality: 2,
		Message: "Misplaced basionym year",
	},
	grm.AuthDoubleParensWarn: Warning{
		Quality: 3,
		Message: "Authroship in double parentheses",
	},
	grm.AuthExWithDotWarn: Warning{
		Quality: 3,
		Message: "`ex` ends with dot",
	},
	grm.AuthEmendWithoutDotWarn: Warning{
		Quality: 3,
		Message: "`emend` without a period",
	},
	grm.AuthShortWarn: Warning{
		Quality: 3,
		Message: "Author is too short",
	},
	grm.AuthUnknownWarn: Warning{
		Quality: 2,
		Message: "Author is unknown",
	},
	grm.AuthQuestionWarn: Warning{
		Quality: 3,
		Message: "Author as a question mark",
	},
	grm.AuthUpperCaseWarn: Warning{
		Quality: 2,
		Message: "Author in upper case",
	},
	grm.YearRangeWarn: Warning{
		Quality: 3,
		Message: "Years range",
	},
	grm.YearDotWarn: Warning{
		Quality: 2,
		Message: "Year with period",
	},
	grm.YearSqBraketsWarn: Warning{
		Quality: 3,
		Message: "Year with square brakets",
	},
	grm.YearPageWarn: Warning{
		Quality: 3,
		Message: "Year with page info",
	},
	grm.YearParensWarn: Warning{
		Quality: 2,
		Message: "Year with parentheses",
	},
	grm.YearCharWarn: Warning{
		Quality: 2,
		Message: "Year with latin character",
	},
	grm.YearQuestionWarn: Warning{
		Quality: 2,
		Message: "Year with question mark",
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
