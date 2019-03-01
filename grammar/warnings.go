package grammar

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
	CanonicalApostropheWarn
	CapWordQuestionWarn
	CharBadWarn
	GenusAbbrWarn
	GenusUpperCharAfterDash
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
	YearOrigMisplacedWarn
	YearPageWarn
	YearParensWarn
	YearQuestionWarn
	YearRangeWarn
	YearSqBraketsWarn
)
