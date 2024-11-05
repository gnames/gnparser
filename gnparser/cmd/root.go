// Package cmd creates a command line application for parsing scientific names.
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/io/web"
	"github.com/gnames/gnsys"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// debug is true when output shows Abstract Synthax Tree instead of
// parsed results.
const debug = false

var (
	// opts is a container for configuration options
	opts []gnparser.Option

	// batchSize determines the size of a batch sent to gnparser workers.
	batchSize int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gnparser file_or_name",
	Short: "Parses scientific names into their semantic elements.",
	Long: `
Parses scientific names into their semantic elements.

To see version:
gnparser -V

To parse one name in CSV format
gnparser "Homo sapiens Linnaeus 1758" [flags]
or (the same)
gnparser "Homo sapiens Linnaeus 1758" -f csv [flags]

To parse one name using JSON format:
gnparser "Homo sapiens Linnaeus 1758" -f compact [flags]
or
gnparser "Homo sapiens Linnaeus 1758" -f pretty [flags]

To parse with maximum amount of details:
gnparser "Homo sapiens Linnaeus 1758" -d -f pretty

To parse many names from a file (one name per line):
gnparser names.txt [flags] > parsed_names.txt

To leave HTML tags and entities intact when parsing (faster)
gnparser names.txt -n > parsed_names.txt

To start web service on port 8080 with 5 concurrent jobs:
gnparser -j 5 -p 8080
 `,

	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag(cmd) {
			os.Exit(0)
		}

		if debug {
			opts = append(opts, gnparser.OptDebug(true))
		}

		formatFlag(cmd)
		jobsNumFlag(cmd)
		ignoreHTMLTagsFlag(cmd)
		withDetailsFlag(cmd)
		withStreamFlag(cmd)
		withNoOrderFlag(cmd)
		withCapitalizeFlag(cmd)
		withEnableCultivarsFlag(cmd)
		withPreserveDiaeresesFlag(cmd)
		batchSizeFlag(cmd)
		spGrCutFlag(cmd)
		port := portFlag(cmd)
		cfg := gnparser.NewConfig(opts...)
		batchSize = cfg.BatchSize

		if port != 0 {
			log.Logger = zerolog.New(os.Stderr).With().
				Str("gnApp", "gnparser").Logger()
			webopts := []gnparser.Option{
				gnparser.OptFormat("compact"),
				gnparser.OptWithWebLogs(withWebLogsFlag(cmd)),
				gnparser.OptWebLogsNsqdTCP(webLogsNsqdTCPFlag(cmd)),
			}
			cfg = gnparser.NewConfig(webopts...)
			gnp := gnparser.New(cfg)
			gnps := web.NewGNparserService(gnp, port)
			web.Run(gnps)
			os.Exit(0)
		}

		quiet, _ := cmd.Flags().GetBool("quiet")
		if quiet {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		}

		if len(args) == 0 {
			processStdin(cmd, cfg)
			os.Exit(0)
		}
		data := getInput(cmd, args)

		if debug {
			debugName(data, cfg)
			os.Exit(0)
		}

		parse(data, cfg)
	},
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once to
// the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("batch_size", "b", 0,
		"maximum number of names in a batch send for processing.")

	rootCmd.Flags().BoolP("cultivar", "C", false,
		"include cultivar epithets and graft-chimeras in normalized and canonical outputs")

	rootCmd.Flags().BoolP("capitalize", "c", false,
		"capitalize the first letter of input name-strings")

	rootCmd.Flags().BoolP("diaereses", "D", false,
		"preserve diaereses in names")

	rootCmd.Flags().BoolP("details", "d", false, "provides more details")

	formatHelp := "sets output format. Can be one of:\n" +
		"'csv', 'tsv', 'compact', 'pretty'\n" +
		"default is 'csv'"
	rootCmd.Flags().StringP("format", "f", "", formatHelp)

	rootCmd.Flags().BoolP("ignore_tags", "i", false,
		"ignore HTML entities and tags when parsing.")

	rootCmd.Flags().IntP("jobs", "j", 0,
		"number of threads to run. CPU's threads number is the default.")

	rootCmd.Flags().IntP("port", "p", 0,
		"starts web site and REST server on the port.")

	rootCmd.Flags().BoolP("quiet", "q", false, "do not show progress")

	rootCmd.Flags().BoolP("stream", "s", false,
		"parse one name at a time in a stream instead of a batch parsing")

	rootCmd.Flags().BoolP("unordered", "u", false,
		"output and input are in different order")

	rootCmd.PersistentFlags().BoolP("version", "V", false,
		"shows build version and date, ignores other flags.")

	rootCmd.Flags().BoolP("web-logs", "", false, "enable logs for the web service")

	rootCmd.Flags().
		BoolP("species-group-cut", "", false, "cut autonym/species group names to species for stemmed version")

	rootCmd.Flags().
		StringP("nsqd-tcp", "", "", "an addresss pointing to NSQ TCP service for logs redirection (e.g. 127.0.0.1:4150)")
}

func processStdin(cmd *cobra.Command, cfg gnparser.Config) {
	if !checkStdin() {
		_ = cmd.Help()
		return
	}
	gnp := gnparser.New(cfg)

	if cfg.WithStream {
		parseStream(gnp, os.Stdin)
	} else {
		parseBatch(gnp, os.Stdin)
	}
}

func checkStdin() bool {
	stdInFile := os.Stdin
	stat, err := stdInFile.Stat()
	if err != nil {
		log.Fatal().Err(err)
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func getInput(cmd *cobra.Command, args []string) string {
	var data string
	switch len(args) {
	case 1:
		data = args[0]
	default:
		_ = cmd.Help()
		os.Exit(0)
	}
	return data
}

func debugName(
	data string,
	cfg gnparser.Config,
) {
	gnp := gnparser.New(cfg)
	res := gnp.Debug(data)
	fmt.Println(string(res))
}

func parse(
	data string,
	cfg gnparser.Config,
) {
	gnp := gnparser.New(cfg)

	path := string(data)
	exists, _ := gnsys.FileExists(path)
	if exists {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal().Err(err)
		}
		if cfg.WithStream {
			parseStream(gnp, f)
		} else {
			parseBatch(gnp, f)
		}
		f.Close()
	} else {
		parseString(gnp, data)
	}
}

func parseString(gnp gnparser.GNparser, name string) {
	res := gnp.ParseName(name)
	f := gnp.Format()

	header := parsed.HeaderCSV(f)
	if header != "" {
		fmt.Println(header)
	}

	fmt.Println(res.Output(f))
}

func progressLog(start time.Time, namesNum int) {
	dur := float64(time.Since(start)) / float64(time.Second)
	rate := float64(namesNum) / dur
	rateStr := humanize.Comma(int64(rate))
	log.Info().
		Str("names/sec", rateStr).
		Str("count", humanize.Comma(int64(namesNum))).
		Msg("File parsing")
}
