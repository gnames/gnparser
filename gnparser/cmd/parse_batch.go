package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/entity/output"
)

func parseBatch(
	gnp gnparser.GNParser,
	f io.Reader,
) {
	batch := make([]string, batchSize)
	chOut := make(chan []output.Parsed)
	var wg sync.WaitGroup

	wg.Add(1)
	go processResults(chOut, &wg, gnp.Format())

	sc := bufio.NewScanner(f)
	var i, count int
	for sc.Scan() {
		batch[count] = sc.Text()
		count++
		if count == batchSize {
			i++
			log.Printf("Parsing %d-th line\n", count*i)
			chOut <- gnp.ParseNames(batch)
			batch = make([]string, batchSize)
			count = 0
		}
	}
	chOut <- gnp.ParseNames(batch[:count])
	close(chOut)
	if err := sc.Err(); err != nil {
		log.Panic(err)
	}
	wg.Wait()
}

func processResults(
	out <-chan []output.Parsed,
	wg *sync.WaitGroup,
	f format.Format,
) {
	defer wg.Done()
	if f == format.CSV {
		fmt.Println(output.HeaderCSV())
	}
	for pr := range out {
		for i := range pr {
			fmt.Println(pr[i].Output(f))
		}
	}
}
