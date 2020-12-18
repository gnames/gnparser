package quality

type Value int

const (
	Failed Value = iota
	MajorProblems
	MediumProblems
	SmallProblems
	Clean
)

var valMap = map[Value]string{
	Failed:         "Failed",
	MajorProblems:  "MajorProblems",
	MediumProblems: "MediumProblems",
	SmallProblems:  "SmallProblems",
	Clean:          "Clean",
}

func (pq Value) String() string {
	res := valMap[pq]
	return res
}
