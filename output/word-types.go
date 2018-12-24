package output

import grm "gitlab.com/gogna/gnparser/grammar"

var wordTypeMap = map[grm.WordType]string{
	grm.UnknownType:         "word",
	grm.AbbrGenusType:       "genus",
	grm.ApproximateYearType: "approximateYear",
	grm.AuthorWordType:      "authorWord",
	grm.GenusType:           "genus",
	grm.SpEpithetType:       "specificEpithet",
	grm.SubGenusType:        "subgenus",
	grm.RankUniType:         "rank",
	grm.TwoLetterGenusType:  "genus",
	grm.UninomialType:       "uninomial",
	grm.YearType:            "year",
}
