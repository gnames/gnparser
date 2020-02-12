package main

/*
	#include "stdlib.h"
*/
import "C"

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"unsafe"

	"gitlab.com/gogna/gnparser"
)

// ParseToString function takes a name-string, desired format, and parses
// the name-string to either JSON, or pipe-separated values, depending on
// the desired format. Format can take values of 'simple', 'compact', 'pretty'.
//export ParseToString
func ParseToString(name *C.char, format *C.char) *C.char {
	goname := C.GoString(name)
	opts := []gnparser.Option{gnparser.OptFormat(C.GoString(format))}
	gnp := gnparser.NewGNparser(opts...)
	parsed, err := gnp.ParseAndFormat(goname)
	if err != nil {
		fmt.Println(err)
		return C.CString("")
	}
	return C.CString(parsed)
}

// ParseAryToStrings function takes an array of names, parsing format and a
// reference to an output: an empty array of strings to return the the data
// back. It populates the output array with raw strings of either JSON or
// pipe-separated parsed values (depending on a given format). Format can take
// values of 'simple', 'compact', or 'pretty'.
//export ParseAryToStrings
func ParseAryToStrings(in **C.char, length C.int, format *C.char, out ***C.char) {
	names := make([]string, int(length))
	inCh := make(chan string)
	outCh := make(chan *gnparser.ParseResult)
	resMap := make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(1)

	opts := []gnparser.Option{
		gnparser.OptFormat(C.GoString(format)),
	}
	jobs := runtime.NumCPU()
	go gnparser.ParseStream(jobs, inCh, outCh, opts...)

	go func() {
		defer wg.Done()
		for parsed := range outCh {
			resMap[parsed.Input] = parsed.Output
		}
	}()

	start := unsafe.Pointer(in)
	pointerSize := unsafe.Sizeof(in)

	for i := 0; i < int(length); i++ {
		// Copy each input string into a Go string and add it to the slice.
		pointer := (**C.char)(unsafe.Pointer(uintptr(start) + uintptr(i)*pointerSize))
		name := C.GoString(*pointer)
		inCh <- name
		names[i] = name
	}

	close(inCh)
	wg.Wait()

	outArray := (C.malloc(C.ulong(length) * C.ulong(pointerSize)))
	*out = (**C.char)(outArray)

	for i := 0; i < int(length); i++ {
		pointer := (**C.char)(unsafe.Pointer(uintptr(outArray) + uintptr(i)*pointerSize))
		if parsed, ok := resMap[names[i]]; ok {
			*pointer = C.CString(parsed)
		} else {
			log.Printf("Cannot find result for %s", names[i])
			*pointer = C.CString("[]")
		}
	}
}

func main() {}
