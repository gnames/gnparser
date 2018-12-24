package grammar

// WordType designates semantic meaning of a word.
type WordType int

const (
	UnknownType WordType = iota
	AbbrGenusType
	ApproximateYearType
	AuthorWordType
	GenusType
	InfraSpEpithetType
	RankType
	RankUniType
	SpEpithetType
	SubGenusType
	SuperSpType
	TwoLetterGenusType
	UninomialType
	YearType
)
