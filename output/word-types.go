package output

import grm "gitlab.com/gogna/gnparser/grammar"

var wordTypeMap = map[grm.WordType]string{
	grm.UnknownType:          "word",
	grm.AuthorWordType:       "authorWord",
	grm.AuthorWordFiliusType: "authorWordFilius",
	grm.GenusType:            "genus",
	grm.InfraSpEpithetType:   "infraspecificEpithet",
	grm.RankType:             "rank",
	grm.RankUniType:          "rank",
	grm.SpEpithetType:        "specificEpithet",
	grm.SubGenusType:         "infragenericEpithet",
	grm.UninomialType:        "uninomial",
	grm.YearApproximateType:  "approximateYear",
	grm.YearType:             "year",
}
