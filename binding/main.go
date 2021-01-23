package main

/*
	#include "stdlib.h"
*/
import "C"

import (
	"strings"
	"unsafe"

	"github.com/gnames/gnlib/encode"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/config"
)

// ParseToString function takes a name-string, desired format, a withDetails
// flag as 0|1 integer. It parses the name-string to either JSON, or a CSV
// string, depending on the desired format. Format argument can take values of
// 'csv', 'compact', 'pretty'. If withDetails argument is 0, additional
// parsed details are ommited, if it is 1 -- they are included.
// true.
//export ParseToString
func ParseToString(
	name *C.char,
	f *C.char,
	details C.int,
) *C.char {
	goname := C.GoString(name)
	opts := []config.Option{
		config.OptFormat(C.GoString(f)),
		config.OptWithDetails(int(details) > 0),
	}
	cfg := config.New(opts...)
	gnp := gnparser.New(cfg)
	parsed := gnp.ParseName(goname).Output(gnp.Format())

	return C.CString(parsed)
}

// FreeMemory takes a string pointer and frees its memory.
//export FreeMemory
func FreeMemory(p *C.char) {
	C.free(unsafe.Pointer(p))
}

// ParseAryToString function takes an array of names, parsing format, and a
// withDetails flag as 0|1 integer.  Parsed outputs are sent as a string in
// either CSV or JSON format.  Format argument can take values of 'csv',
// 'compact', or 'pretty'. For withDetails argument 0 means false, 1 means
// true.
//export ParseAryToString
func ParseAryToString(
	in **C.char,
	length C.int,
	f *C.char,
	details C.int,
) *C.char {
	names := make([]string, int(length))

	opts := []config.Option{
		config.OptFormat(C.GoString(f)),
		// config.OptJobsNum(runtime.NumCPU() * 2),
		config.OptWithDetails(int(details) > 0),
	}
	start := unsafe.Pointer(in)
	pointerSize := unsafe.Sizeof(in)

	for i := 0; i < int(length); i++ {
		// Copy each input string into a Go string and add it to the slice.
		pointer := (**C.char)(unsafe.Pointer(uintptr(start) + uintptr(i)*pointerSize))
		name := C.GoString(*pointer)
		names[i] = name
	}

	cfg := config.New(opts...)
	gnp := gnparser.New(cfg)

	var res string
	parsed := gnp.ParseNames(names)
	if gnp.Format() == format.CSV {
		csv := make([]string, length)
		for i := range parsed {
			csv[i] = parsed[i].Output(format.CSV)
		}
		res = strings.Join(csv, "\n")
	} else {
		json, _ := encode.GNjson{}.Encode(parsed)
		res = string(json)
	}
	return C.CString(res)
}

func main() {}
