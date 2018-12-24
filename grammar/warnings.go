package grammar

type Warning int

const (
	TailWarn Warning = iota
	AbbrGenusWarn
	ApproxWarning
	AuthDoubleParensWarn
	AuthExWarn
	AuthMisformedYearWarn
	AuthQuestionWarn
	AuthShortWarn
	AuthUnknownWarn
	BadCharsWarn
	CapWordQuestionWarn
	CompWarning
	ExWithDotWarn
	RareRankWarn
	SuperSpWarn
	UninomialComboWarn
	YearCharWarn
	YearDotWarn
	YearPageWarn
	YearParensWarn
	YearQuestionWarn
	YearRangeWarn
	YearSqBraketsWarn
)
