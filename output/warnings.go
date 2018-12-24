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
		Message: "Unparseable tail",
	},
	grm.BadCharsWarn: Warning{
		Quality: 2,
		Message: "Non-standard characters in canonical",
	},
	grm.YearCharWarn: Warning{
		Quality: 2,
		Message: "Year with latin character",
	},
	grm.YearParensWarn: Warning{
		Quality: 2,
		Message: "Year with parentheses",
	},
	grm.YearQuestionWarn: Warning{
		Quality: 2,
		Message: "Year with question mark",
	},
	grm.UninomialComboWarn: Warning{
		Quality: 2,
		Message: "Combination of two uninomials",
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
