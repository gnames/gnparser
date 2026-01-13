// Package main provides C-binding functionality to use parser in
// other languages.
package main

/*
  #include "stdlib.h"
*/
import "C"

import (
	"strings"
	"unsafe"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser"
)

// ParseToString function takes a name-string, desired format, a withDetails
// flag as 0|1 integer. It parses the name-string to either JSON, or a CSV
// string, depending on the desired format. Format argument can take values of
// 'csv', 'compact', 'pretty'. If withDetails argument is 0, additional
// parsed details are ommited, if it is 1 -- they are included.
// true.
//
//export ParseToString
func ParseToString(
	name *C.char,
	fmtStr *C.char,
	codeStr *C.char,
	details C.int,
	diaereses C.int,
	compactAuthors C.int,
	flatten C.int,
) *C.char {
	goname := C.GoString(name)
	code := nomcode.New(C.GoString(codeStr))
	frmt, err := gnfmt.NewFormat(C.GoString(fmtStr))
	if err != nil {
		frmt = gnfmt.CSV
	}
	opts := []gnparser.Option{
		gnparser.OptFormat(frmt),
		gnparser.OptWithDetails(int(details) > 0),
		gnparser.OptCode(code),
		gnparser.OptWithPreserveDiaereses(int(diaereses) > 0),
		gnparser.OptWithCompactAuthors(int(compactAuthors) > 0),
		gnparser.OptWithFlatOutput(int(flatten) > 0),
	}
	cfg := gnparser.NewConfig(opts...)
	gnp := gnparser.New(cfg)
	parsed := gnp.ParseName(goname).Output(gnp.Format(), gnp.FlatOutput())

	return C.CString(parsed)
}

// FreeMemory takes a string pointer and frees its memory.
//
//export FreeMemory
func FreeMemory(p *C.char) {
	C.free(unsafe.Pointer(p))
}

// ParseAryToString function takes an array of names, parsing format, and a
// withDetails flag as 0|1 integer.  Parsed outputs are sent as a string in
// either CSV or JSONformat. Format argument can take values of 'csv',
// 'compact', or 'pretty'. For withDetails argument 0 means false, 1 means
// true.
//
//export ParseAryToString
func ParseAryToString(
	in **C.char,
	length C.int,
	fmtStr *C.char,
	codeStr *C.char,
	details C.int,
	diaereses C.int,
	compactAuthors C.int,
	flatten C.int,
) *C.char {
	names := make([]string, int(length))
	code := nomcode.New(C.GoString(codeStr))
	frmt, err := gnfmt.NewFormat(C.GoString(fmtStr))
	if err != nil {
		frmt = gnfmt.CSV
	}

	opts := []gnparser.Option{
		gnparser.OptFormat(frmt),
		gnparser.OptWithDetails(int(details) > 0),
		gnparser.OptCode(code),
		gnparser.OptWithPreserveDiaereses(int(diaereses) > 0),
		gnparser.OptWithCompactAuthors(int(compactAuthors) > 0),
		gnparser.OptWithFlatOutput(int(flatten) > 0),
	}
	start := unsafe.Pointer(in)
	pointerSize := unsafe.Sizeof(in)

	for i := 0; i < int(length); i++ {
		// Copy each input string into a Go string and add it to the slice.
		pointer := (**C.char)(unsafe.Pointer(uintptr(start) + uintptr(i)*pointerSize))
		name := C.GoString(*pointer)
		names[i] = name
	}

	cfg := gnparser.NewConfig(opts...)
	gnp := gnparser.New(cfg)

	var res string
	parsed := gnp.ParseNames(names)
	if gnp.Format() == gnfmt.CSV {
		csv := make([]string, length)
		for i := range parsed {
			csv[i] = parsed[i].Output(gnfmt.CSV, gnp.FlatOutput())
		}
		res = strings.Join(csv, "\n")
	} else {
		json, _ := gnfmt.GNjson{}.Encode(parsed)
		res = string(json)
	}
	return C.CString(res)
}

func main() {}
