package output

// ParseResult structure contains parsing output, its place in the
// slice, and an unexpected error, if it happened durin the parsing.
type ParseResult struct {
	Index  int
	Parsed Parsed
	Error  error
}
