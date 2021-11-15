package gnparser

import (
	"log"
	"runtime"

	"github.com/gnames/gnfmt"
)

// Config keeps settings that might affect how parsing is done,
// of change the parsing output.
type Config struct {
	// Format sets the output format for CLI and Web interfaces.
	// There are 3 formats available: 'CSV', 'CompactJSON' and
	// 'PrettyJSON'.
	Format gnfmt.Format

	// JobsNum sets a level of parallelism used during parsing of
	// a stream of name-strings.
	JobsNum int

	// BatchSize sets the maximum number of elements in names-strings slice.
	BatchSize int

	// WithStream changes from parsing a batch by batch, to parsing one name
	// at a time. When WithStream is true, BatchSize setting is ignored.
	WithStream bool

	// IgnoreHTMLTags can be set to true when it is desirable to clean up names
	// from a few HTML tags often present in names-strings that were planned to
	// be presented via an HTML page.
	IgnoreHTMLTags bool

	// WithDetails can be set to true when a simplified output is not sufficient
	// for obtaining a required information.
	WithDetails bool

	// WithNoOrder flag, when true, output and input are in different order.
	WithNoOrder bool

	// WithCapitalization flag, when true, the first letter of a name-string
	// is capitalized, if appropriate.
	WithCapitalization bool

	// WithPreserveDiaereses flag, when true, diaereses will not be transliterated
	WithPreserveDiaereses bool

	// WithCultivars flag, when true, cultivar names will be parsed and
	// modify cardinality, normalized and canonical output.
	WithCultivars bool

	// Port to run wer-service.
	Port int

	// IsTest can be set to true when parsing functionality is used for tests.
	// In such cases the `ParserVersion` field is presented as `test_version`
	// instead of displaying the actual version of `gnparser`.
	IsTest bool

	// Debug sets a "debug" state for parsing. The debug state forces output
	// format to showing parsed ast tree.
	Debug bool
}

// Option is a type that has to be returned by all Option functions. Such
// functions are able to modify the settings of a Config object.
type Option func(*Config)

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

// OptDebugParse returns parsed tree
func OptDebug(b bool) Option {
	return func(cfg *Config) {
		cfg.Debug = b
	}
}

// OptFormat takes a string (one of 'csv', 'compact', 'pretty') to set
// the formatting option for the CLI or Web presentation. If some other
// string is entered, the default, 'CSV' format is set, accompanied by a
// warning.
func OptFormat(s string) Option {
	return func(cfg *Config) {
		f, err := gnfmt.NewFormat(s)
		if err != nil {
			f = gnfmt.CSV
			log.Printf("Set default CSV format due to error: %s.", err)
		}
		cfg.Format = f
	}
}

// OptKeepHTMLTags sets the KeepHTMLTags field. This option is useful if
// names with HTML tags shold not be parsed, or they are absent in input
// data.
func OptIgnoreHTMLTags(b bool) Option {
	return func(cfg *Config) {
		cfg.IgnoreHTMLTags = b
	}
}

// OptIsTest sets a test flag.
func OptIsTest(b bool) Option {
	return func(cfg *Config) {
		cfg.IsTest = b
	}
}

// OptJobsNum sets the JobsNum field.
func OptJobsNum(i int) Option {
	return func(cfg *Config) {
		cfg.JobsNum = i
	}
}

// OptPort sets a port for web-service.
func OptPort(i int) Option {
	return func(cfg *Config) {
		cfg.Port = i
	}
}

// OptWithCapitaliation sets the WithCapitalization field.
func OptWithCapitaliation(b bool) Option {
	return func(cfg *Config) {
		cfg.WithCapitalization = b
	}
}

// OptPreserveDiaereses sets the PreserveDiaereses field.
func OptWithPreserveDiaereses(b bool) Option {
	return func(cfg *Config) {
		cfg.WithPreserveDiaereses = b
	}
}

// OptWithCultivars sets the EnableCultivars field.
func OptWithCultivars(b bool) Option {
	return func(cfg *Config) {
		cfg.WithCultivars = b
	}
}

// OptWithDetails sets the WithDetails field.
func OptWithDetails(b bool) Option {
	return func(cfg *Config) {
		cfg.WithDetails = b
	}
}

// OptWithNoOrder sets the WithNoOrder field.
func OptWithNoOrder(b bool) Option {
	return func(cfg *Config) {
		cfg.WithNoOrder = b
	}
}

// OptWithDetails sets the WithDetails field.
func OptWithStream(b bool) Option {
	return func(cfg *Config) {
		cfg.WithStream = b
	}
}

// NewConfig generates a new Config object. It can take an arbitrary number
// of `Option` functions to modify default configuration settings.
func NewConfig(opts ...Option) Config {
	cfg := Config{
		Format:         gnfmt.CSV,
		JobsNum:        runtime.NumCPU(),
		BatchSize:      50_000,
		IgnoreHTMLTags: false,
		Port:           8080,
	}
	for i := range opts {
		opts[i](&cfg)
	}
	return cfg
}
