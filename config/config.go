package config

import (
	"log"
	"runtime"

	"github.com/gnames/gnlib/format"
)

// Config keeps settings that might affect how parsing is done,
// of change the parsing output.
type Config struct {
	// Format sets the output format for CLI and Web interfaces.
	// There are 3 formats available: 'CSV', 'CompactJSON' and
	// 'PrettyJSON'.
	Format format.Format

	// JobsNum sets a level of parallelism used during parsing of
	// a stream of name-strings.
	JobsNum int

	// BatchSize sets the maximum number of elements in names-strings slice.
	BatchSize int

	// KeepHTMLTags can be set to true when it is desirable to clean up names from
	// a few HTML tags often present in names-strings that were planned to be
	// presented via an HTML page.
	KeepHTMLTags bool

	// WithDetails can be set to true when a simplified output is not sufficient
	// for obtaining a required information.
	WithDetails bool

	// Port to run wer-service.
	Port int

	// IsTest can be set to true when parsing functionality is used for tests.
	// In such cases the `ParserVersion` field is presented as `test_version`
	// instead of displaying the actual version of `gnparser`.
	IsTest bool
}

// NewConfig generates a new Config object. It can take an arbitrary number
// of `Opt` functions to modify default configuration settings.
func NewConfig(opts ...Option) Config {
	cfg := Config{
		Format:       format.CSV,
		JobsNum:      runtime.NumCPU(),
		BatchSize:    50_000,
		KeepHTMLTags: false,
		Port:         8080,
	}
	for i := range opts {
		opts[i](&cfg)
	}
	return cfg
}

// Option is a type that has to be returned by all Option functions. Such
// functions are able to modify the settings of a Config object.
type Option func(*Config)

// OptFormat takes a string (one of 'csv', 'compact', 'pretty') to set
// the formatting option for the CLI or Web presentation. If some other
// string is entered, the default, 'CSV' format is set, accompanied by a
// warning.
func OptFormat(s string) Option {
	return func(cfg *Config) {
		f, err := format.NewFormat(s)
		if err != nil {
			f = format.CSV
			log.Printf("Set default CSV format due to error: %s.", err)
		}
		cfg.Format = f
	}
}

// OptJobsNum sets the JobsNum field.
func OptJobsNum(i int) Option {
	return func(cfg *Config) {
		cfg.JobsNum = i
	}
}

// OptKeepHTMLTags sets the RemoveHTML field.
func OptKeepHTMLTags(b bool) Option {
	return func(cfg *Config) {
		cfg.KeepHTMLTags = b
	}
}

// OptWithDetails sets the WithDetails field.
func OptWithDetails(b bool) Option {
	return func(cfg *Config) {
		cfg.WithDetails = b
	}
}

// OptBatchSize sets the max number of names in a batch.
func OptBatchSize(i int) Option {
	return func(cfg *Config) {
		if i <= 0 {
			log.Println("Batch size should be a positive number")
			return
		}
		cfg.BatchSize = i
	}
}

// OptPort sets a port for web-service.
func OptPort(i int) Option {
	return func(cfg *Config) {
		cfg.Port = i
	}
}

// OptIsTest sets a test flag.
func OptIsTest(b bool) Option {
	return func(cfg *Config) {
		cfg.IsTest = b
	}
}
