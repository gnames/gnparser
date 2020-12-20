package config_test

import (
	"runtime"
	"testing"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.NewConfig()
	deflt := config.Config{
		Format:       format.CSV,
		JobsNum:      runtime.NumCPU(),
		KeepHTMLTags: false,
		WithDetails:  false,
		Port:         8080,
		IsTest:       false,
	}
	assert.Equal(t, cfg, deflt)
}

func TestNewOpts(t *testing.T) {
	opts := opts()
	cnf := config.NewConfig(opts...)
	updt := config.Config{
		Format:       format.CompactJSON,
		JobsNum:      161,
		KeepHTMLTags: true,
		WithDetails:  true,
		Port:         8989,
	}
	assert.Equal(t, cnf, updt)
}

func opts() []config.Option {
	return []config.Option{
		config.OptFormat("compact"),
		config.OptJobsNum(161),
		config.OptKeepHTMLTags(true),
		config.OptWithDetails(true),
		config.OptPort(8989),
	}
}
