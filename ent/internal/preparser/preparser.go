package preparser

import (
	"log/slog"
	"unicode/utf8"
)

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
		// Convert rune index to byte index without allocating.
		byteIdx := 0
		for range ppr.tailIndex {
			_, sz := utf8.DecodeRuneInString(s[byteIdx:])
			byteIdx += sz
		}
		return byteIdx
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
