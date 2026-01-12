package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/nameidx"
	"github.com/gnames/gnparser/ent/parsed"
)

func getNames(
	ctx context.Context,
	f io.Reader,
) <-chan nameidx.NameIdx {
	chIn := make(chan nameidx.NameIdx)
	sc := bufio.NewScanner(f)

	go func() {
		defer close(chIn)
		var count int
		for sc.Scan() {
			nameString := sc.Text()
			select {
			case <-ctx.Done():
				return
			case chIn <- nameidx.NameIdx{Index: count, NameString: nameString}:
			}
			count++
		}
	}()
	if err := sc.Err(); err != nil {
		slog.Error("Cannot read data", "error", err)
	}
	return chIn
}

func parseStream(
	gnp gnparser.GNparser,
	f io.Reader,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chIn := getNames(ctx, f)
	chOut := make(chan parsed.Parsed)
	var wg sync.WaitGroup
	wg.Add(1)

	go gnp.ParseNameStream(ctx, chIn, chOut)

	// process parsing results
	go func() {
		defer cancel()
		defer wg.Done()
		start := time.Now()

		header := parsed.HeaderCSV(gnp.Format(), gnp.WithDetails())
		if header != "" {
			fmt.Println(header)
		}

		var count int
		for {
			count++
			if count%50_000 == 0 {
				progressLog(start, count)
			}
			select {
			case <-ctx.Done():
				return
			case v, ok := <-chOut:
				if !ok {
					return
				}
				fmt.Println(v.Output(gnp.Format(), gnp.FlatOutput()))
			}
		}
	}()
	wg.Wait()
}
