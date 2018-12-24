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
	BadUTF8ConvWarn
	CapWordQuestionWarn
	CompWarning
	ExWithDotWarn
	UncommonRankWarn
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
