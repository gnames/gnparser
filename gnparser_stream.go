package gnparser

import (
	"context"
	"log"
	"sync"

	"github.com/gnames/gnlib/organizer"
	"github.com/gnames/gnparser/entity/input"
	"github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/parser"
)

func (gnp gnparser) ParseNameStream(
	ctx context.Context,
	chIn <-chan input.Name,
	chOut chan<- output.Parsed,
) {
	chToOrder := make(chan organizer.Ordered)
	chOrdered := make(chan organizer.Ordered)
	var wgWorker, wgOrd sync.WaitGroup
	jobs := gnp.cfg.JobsNum
	wgWorker.Add(jobs)
	wgOrd.Add(1)

	for i := jobs; i > 0; i-- {
		go gnp.parseStreamWorker(ctx, chIn, chToOrder, &wgWorker)
	}

	go organizer.Organize(ctx, chToOrder, chOrdered)

	go sendOrdered(ctx, chOrdered, chOut, &wgOrd)

	wgWorker.Wait()
	close(chToOrder)
	wgOrd.Wait()
}

func (gnp gnparser) parseStreamWorker(
	ctx context.Context,
	chIn <-chan input.Name,
	chOut chan<- organizer.Ordered,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	gnp.parser = parser.NewParser()
	for v := range chIn {
		parsed := gnp.ParseName(v.NameString)
		select {
		case <-ctx.Done():
			return
		case chOut <- output.ParseResult{Parsed: parsed, Error: nil, Idx: v.Index}:
		}
	}
}

func sendOrdered(
	ctx context.Context,
	chOrdered <-chan organizer.Ordered,
	chOut chan<- output.Parsed,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for v := range chOrdered {
		var p output.Parsed
		err := v.Unpack(&p)
		if err != nil {
			log.Panic(err)
		}
		select {
		case <-ctx.Done():
			return
		case chOut <- p:
		}
	}
	close(chOut)
}
