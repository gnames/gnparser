package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/entity/input"
	"github.com/gnames/gnparser/entity/output"
)

func getNames(
	ctx context.Context,
	f io.Reader,
) <-chan input.Name {
	chIn := make(chan input.Name)
	sc := bufio.NewScanner(f)

	go func() {
		defer close(chIn)
		var count int
		for sc.Scan() {
			nameString := sc.Text()
			select {
			case <-ctx.Done():
				return
			case chIn <- input.Name{Index: count, NameString: nameString}:
			}
			count++
		}
	}()
	if err := sc.Err(); err != nil {
		log.Panic(err)
	}
	return chIn
}

func parseStream(
	gnp gnparser.GNParser,
	f io.Reader,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chIn := getNames(ctx, f)
	chOut := make(chan output.Parsed)
	var wg sync.WaitGroup
	wg.Add(1)

	if gnp.Format() == format.CSV {
		output.HeaderCSV()
	}

	go gnp.ParseNameStream(ctx, chIn, chOut)

	go func() {
		defer cancel()
		defer wg.Done()
		var count int
		for {
			count++
			if count%50_000 == 0 {
				log.Printf("Processing %d-th name", count)
			}
			select {
			case <-ctx.Done():
				return
			case v, ok := <-chOut:
				if !ok {
					return
				}
				fmt.Println(v.Output(gnp.Format()))
			}
		}
	}()
	wg.Wait()
}
