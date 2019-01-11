package output

import grm "gitlab.com/gogna/gnparser/grammar"

var wordTypeMap = map[grm.WordType]string{
	grm.UnknownType:          "word",
	grm.ComparisonType:       "annotationIdentification",
	grm.ApproxType:           "annotationIdentification",
	grm.AuthorWordType:       "authorWord",
	grm.AuthorWordExType:     "authorWord",
	grm.AuthorWordEmendType:  "authorWord",
	grm.AuthorWordFiliusType: "authorWordFilius",
	grm.GenusType:            "genus",
	grm.HybridCharType:       "hybridChar",
	grm.InfraSpEpithetType:   "infraspecificEpithet",
	grm.RankType:             "rank",
	grm.RankUniType:          "rank",
	grm.SpEpithetType:        "specificEpithet",
	grm.SubGenusType:         "infragenericEpithet",
	grm.UninomialType:        "uninomial",
	grm.YearApproximateType:  "approximateYear",
	grm.YearType:             "year",
}
