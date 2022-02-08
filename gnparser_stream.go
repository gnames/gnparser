package gnparser

import (
	"context"
	"sync"

	"github.com/gnames/gnparser/ent/nameidx"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/parser"
	"github.com/gnames/organizer"
	"github.com/rs/zerolog/log"
)

// ParseNameStream takes an input channel of input.Name and
// returns back a stream of parsed data following the same order as
// the input.
func (gnp gnparser) ParseNameStream(
	ctx context.Context,
	chIn <-chan nameidx.NameIdx,
	chOut chan<- parsed.Parsed,
) {
	chUnordered := make(chan organizer.Ordered)
	chOrdered := make(chan organizer.Ordered)
	var wgWorker, wgOutput sync.WaitGroup
	jobs := gnp.cfg.JobsNum
	wgWorker.Add(jobs)
	wgOutput.Add(1)

	for i := jobs; i > 0; i-- {
		go gnp.parseStreamWorker(ctx, chIn, chUnordered, &wgWorker)
	}

	if gnp.cfg.WithNoOrder {
		close(chOrdered)
		go sendUnordered(ctx, chUnordered, chOut, &wgOutput)
	} else {
		go organizer.Organize(ctx, chUnordered, chOrdered)
		go sendOrdered(ctx, chOrdered, chOut, &wgOutput)
	}

	wgWorker.Wait()
	close(chUnordered)
	wgOutput.Wait()
}

func (gnp gnparser) parseStreamWorker(
	ctx context.Context,
	chIn <-chan nameidx.NameIdx,
	chOut chan<- organizer.Ordered,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	gnp.parser = parser.New()
	for v := range chIn {
		parseRes := gnp.ParseName(v.NameString)
		select {
		case <-ctx.Done():
			return
		case chOut <- parsed.ParsedWithIdx{Parsed: parseRes, Error: nil, Idx: v.Index}:
		}
	}
}

func sendOrdered(
	ctx context.Context,
	chOrdered <-chan organizer.Ordered,
	chOut chan<- parsed.Parsed,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for v := range chOrdered {
		var p parsed.Parsed
		err := v.Unpack(&p)
		if err != nil {
			log.Fatal().Err(err)
		}
		select {
		case <-ctx.Done():
			return
		case chOut <- p:
		}
	}
	close(chOut)
}

func sendUnordered(
	ctx context.Context,
	chUnordered <-chan organizer.Ordered,
	chOut chan<- parsed.Parsed,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for v := range chUnordered {
		var p parsed.Parsed
		err := v.Unpack(&p)
		if err != nil {
			log.Fatal().Err(err)
		}
		select {
		case <-ctx.Done():
			return
		case chOut <- p:
		}
	}
	close(chOut)
}
