package output

import "fmt"

// ParseResult structure contains parsing output, its place in the
// slice, and an unexpected error, if it happened durin the parsing.
type ParseResult struct {
	Idx    int
	Parsed Parsed
	Error  error
}

func (pr ParseResult) Index() int {
	return pr.Idx
}

func (pr ParseResult) Unpack(v interface{}) error {
	if pr.Error != nil {
		return pr.Error
	}
	switch p := v.(type) {
	case *Parsed:
		*p = pr.Parsed
		return nil
	default:
		return fmt.Errorf("cannot use %T as Parsed", v)
	}
}
