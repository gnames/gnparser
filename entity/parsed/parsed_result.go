package parsed

import "fmt"

// ParsedWithIdx structure contains parsing output, its place in the
// slice, and an unexpected error, if it happened during the parsing.
type ParsedWithIdx struct {
	Idx    int
	Parsed Parsed
	Error  error
}

func (pr ParsedWithIdx) Index() int {
	return pr.Idx
}

func (pr ParsedWithIdx) Unpack(v interface{}) error {
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
