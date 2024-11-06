package preparser

import "log/slog"

func New() *PreParser {
	res := &PreParser{}
	res.Init()
	return res
}

type PreString struct {
	tailIndex int
}

// ParseString returns index of the Tail
func (ppr *PreParser) NewString(s string) {
	ppr.tailIndex = -1
	ppr.Buffer = s
	ppr.Reset()
}

func (ppr *PreParser) TailIndex(s string) int {
	ppr.NewString(s)
	if err := ppr.Parse(); err != nil {
		slog.Error("Preparsing failed", "error", err, "string", s)
		return -1
	}
	ppr.Execute()
	if ppr.tailIndex >= 0 {
		rs := []rune(s)
		head := rs[0:ppr.tailIndex]
		return len([]byte(string(head)))
	}
	return ppr.tailIndex
}

// Debug takes a string, parses it, and prints its AST.
func (ppr *PreParser) Debug(q string) error {
	ppr.NewString(q)
	err := ppr.Parse()
	if err != nil {
		return err
	}
	ppr.PrettyPrintSyntaxTree(q)
	return nil
}
