package gnparser_test

import (
	"runtime"
	"testing"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := gnparser.NewConfig()
	deflt := gnparser.Config{
		Format:         format.CSV,
		JobsNum:        runtime.NumCPU(),
		BatchSize:      50_000,
		IgnoreHTMLTags: false,
		WithDetails:    false,
		Port:           8080,
		IsTest:         false,
	}
	assert.Equal(t, cfg, deflt)
}

func TestNewOpts(t *testing.T) {
	opts := opts()
	cnf := gnparser.NewConfig(opts...)
	updt := gnparser.Config{
		Format:         format.CompactJSON,
		JobsNum:        161,
		BatchSize:      1,
		IgnoreHTMLTags: true,
		WithDetails:    true,
		Port:           8989,
	}
	assert.Equal(t, cnf, updt)
}

func opts() []gnparser.Option {
	return []gnparser.Option{
		gnparser.OptFormat("compact"),
		gnparser.OptJobsNum(161),
		gnparser.OptBatchSize(1),
		gnparser.OptIgnoreHTMLTags(true),
		gnparser.OptWithDetails(true),
		gnparser.OptPort(8989),
	}
}
