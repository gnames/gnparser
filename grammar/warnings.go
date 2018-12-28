package grammar

type Warning int

const (
	TailWarn Warning = iota
	AuthDoubleParensWarn
	AuthExWarn
	AuthMisformedYearWarn
	AuthQuestionWarn
	AuthShortWarn
	AuthUnknownWarn
	AuthUpperCaseWarn
	CanonicalApostropheWarn
	CapWordQuestionWarn
	CharBadWarn
	ExWithDotWarn
	GenusAbbrWarn
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
	SpeciesNumericWarn
	SuperSpeciesWarn
	UTF8ConvBadWarn
	UninomialComboWarn
	WhiteSpaceTrailWarn
	YearCharWarn
	YearDotWarn
	YearOrigMisformedWarn
	YearPageWarn
	YearParensWarn
	YearQuestionWarn
	YearRangeWarn
	YearSqBraketsWarn
)
