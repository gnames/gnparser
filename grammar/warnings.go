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
