package grammar

// WordType designates semantic meaning of a word.
type WordType int

const (
	UnknownType WordType = iota
	AuthorWordType
	AuthorWordFiliusType
	GenusType
	InfraSpEpithetType
	RankType
	RankUniType
	SpEpithetType
	SubGenusType
	SuperSpType
	UninomialType
	YearApproximateType
	YearType
)
