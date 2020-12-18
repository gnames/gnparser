package quality

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

var warningValueMap = map[Warning]Value{
	TailWarn:                        MajorProblems,
	ApostrOtherWarn:                 MediumProblems,
	AuthAmbiguousFiliusWarn:         SmallProblems,
	AuthDoubleParensWarn:            MajorProblems,
	AuthExWarn:                      SmallProblems,
	AuthExWithDotWarn:               MediumProblems,
	AuthEmendWarn:                   SmallProblems,
	AuthEmendWithoutDotWarn:         MediumProblems,
	AuthMissingOneParensWarn:        MajorProblems,
	AuthQuestionWarn:                MajorProblems,
	AuthShortWarn:                   MediumProblems,
	AuthUnknownWarn:                 SmallProblems,
	AuthUpperCaseWarn:               SmallProblems,
	BacteriaMaybeWarn:               Clean,
	BotanyAuthorNotSubgenWarn:       SmallProblems,
	CanonicalApostropheWarn:         MediumProblems,
	CapWordQuestionWarn:             MajorProblems,
	CharBadWarn:                     SmallProblems,
	GenusAbbrWarn:                   MajorProblems,
	GenusUpperCharAfterDash:         SmallProblems,
	GreekLetterInRank:               SmallProblems,
	HTMLTagsEntitiesWarn:            MediumProblems,
	HybridCharNoSpaceWarn:           MediumProblems,
	HybridFormulaWarn:               SmallProblems,
	HybridFormulaIncompleteWarn:     MajorProblems,
	HybridFormulaProbIncompleteWarn: SmallProblems,
	HybridNamedWarn:                 SmallProblems,
	NameApproxWarn:                  MajorProblems,
	NameComparisonWarn:              MajorProblems,
	RankUncommonWarn:                MediumProblems,
	SpaceMultipleWarn:               SmallProblems,
	SpaceNonStandardWarn:            SmallProblems,
	SpanishAndAsSeparator:           SmallProblems,
	SpeciesNumericWarn:              MediumProblems,
	SuperSpeciesWarn:                SmallProblems,
	UTF8ConvBadWarn:                 MajorProblems,
	UninomialComboWarn:              SmallProblems,
	WhiteSpaceTrailWarn:             SmallProblems,
	YearCharWarn:                    SmallProblems,
	YearDotWarn:                     SmallProblems,
	YearOrigMisplacedWarn:           SmallProblems,
	YearPageWarn:                    SmallProblems,
	YearParensWarn:                  SmallProblems,
	YearQuestionWarn:                SmallProblems,
	YearRangeWarn:                   MediumProblems,
	YearSqBraketsWarn:               MediumProblems,
}

type QualityWarning struct {
	Value   Value   `json:"value"`
	Warning Warning `json:"warning"`
}

func (w Warning) String() string {
	return warningMap[w]
}

func (w Warning) Value() Value {
	return warningValueMap[w]
}

func (w Warning) QualityWarning() QualityWarning {
	return QualityWarning{
		Value:   w.Value(),
		Warning: w,
	}
}

func Map(ws []Warning) []QualityWarning {
	res := make([]QualityWarning, len(ws))
	for i, v := range ws {
		res[i] = v.QualityWarning()
	}
	return res
}
