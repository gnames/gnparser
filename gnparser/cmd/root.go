// Copyright © 2019 Dmitry Mozzherin <dmozzherin@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnlib/sys"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/config"
	"github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/io/web"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configText = `# Format sets the output format for CLI and Web
interfaces. There are 3 possible settings: 'csv', 'compact', 'pretty'.
# Format csv

# JobsNum sets the level of parallelism used during parsing of a stream
# of name-strings.
# JobsNum 4

# BatchSize determines maximum number of name-strings sent simultaneously
# for parsing. When it is important to have no delay in parsing, set the
# BatchSize to 1.
# BatchSize 50000

# IgnoreHTMLTags can be set to true if it is desirable to not try to remove from
# a few HTML tags often present in names-strings that were planned to be
# presented via an HTML page.
# IgnoreHTMLTags false

# WithDetails can be set to true when a simplified output is not sufficient
# for obtaining a required information.
# WithDetails false

# Port is a port for the gnames service
# Port: 8080
`
)

var (
	opts      []config.Option
	batchSize int
)

// config purpose is to achieve automatic import of data from the
// configuration file, if it exists.
type cfgData struct {
	Format         string
	JobsNum        int
	BatchSize      int
	IgnoreHTMLTags bool
	WithDetails    bool
	Port           int
}

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

		formatFlag(cmd)
		jobsNumFlag(cmd)
		ignoreHTMLTagsFlag(cmd)
		withDetailsFlag(cmd)
		batchSizeFlag(cmd)
		port := portFlag(cmd)
		cfg := config.NewConfig(opts...)
		batchSize = cfg.BatchSize

		if port != 0 {
			gnp := gnparser.NewGNParser(cfg)
			gnps := web.NewGNParserService(gnp, port)
			web.Run(gnps)
			os.Exit(0)
		}

		if len(args) == 0 {
			processStdin(cmd, cfg)
			os.Exit(0)
		}
		data := getInput(cmd, args)
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
	rootCmd.PersistentFlags().BoolP("version", "V", false,
		"shows build version and date, ignores other flags.")

	formatHelp := "sets output format. Can be one of:\n  " +
		"'csv', 'compact', 'pretty'"
	rootCmd.Flags().StringP("format", "f", "", formatHelp)

	rootCmd.Flags().IntP("jobs", "j", 0,
		"nubmer of threads to run. CPU's threads number is the default.")

	rootCmd.Flags().IntP("batch_size", "b", 0,
		"maximum number of names in a batch send for processing.")

	rootCmd.Flags().BoolP("ignore_tags", "i", false,
		"ignore HTML entities and tags when parsing.")

	rootCmd.Flags().BoolP("details", "d", false, "provides more details")

	rootCmd.Flags().IntP("port", "p", 0,
		"starts web site and REST server on the port.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configFile := "gnparser"
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Cannot find home directory: %s.", err)
	}
	home = filepath.Join(home, ".config")

	// Search config in home directory with name ".gnames" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(configFile)

	// Set environment variables to override
	// config file settings
	viper.BindEnv("Format", "GNPARSER_FORMAT")
	viper.BindEnv("JobsNum", "GNPARSER_JOBS_NUM")
	viper.BindEnv("IgnoreHTMLTags", "GNPARSER_IGNORE_HTML_TAGS")
	viper.BindEnv("WithDetails", "GNPARSER_WITH_DETAILS")
	viper.BindEnv("Port", "GNPARSER_PORT")

	viper.AutomaticEnv() // read in environment variables that match

	configPath := filepath.Join(home, fmt.Sprintf("%s.yaml", configFile))
	touchConfigFile(configPath, configFile)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s.", viper.ConfigFileUsed())
	}
	getOpts()
}

func getOpts() []config.Option {
	cfg := &cfgData{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Cannot deserialize config data: %s.", err)
	}
	if cfg.Format != "" {
		opts = append(opts, config.OptFormat(cfg.Format))
	}
	if cfg.JobsNum != 0 {
		opts = append(opts, config.OptJobsNum(cfg.JobsNum))
	}
	if cfg.BatchSize > 0 {
		opts = append(opts, config.OptBatchSize(cfg.BatchSize))
	}
	if cfg.IgnoreHTMLTags != false {
		opts = append(opts, config.OptIgnoreHTMLTags(cfg.IgnoreHTMLTags))
	}
	if cfg.WithDetails != false {
		opts = append(opts, config.OptWithDetails(cfg.WithDetails))
	}
	if cfg.Port != 0 {
		opts = append(opts, config.OptPort(cfg.Port))
	}

	return opts
}

// touchConfigFile checks if config file exists, and if not, it gets created.
func touchConfigFile(configPath string, configFile string) {
	if sys.FileExists(configPath) {
		return
	}

	log.Printf("Creating config file: %s.", configPath)
	createConfig(configPath, configFile)
}

// createConfig creates config file.
func createConfig(path string, file string) {
	err := sys.MakeDir(filepath.Dir(path))
	if err != nil {
		log.Fatalf("Cannot create dir %s: %s.", path, err)
	}

	err = ioutil.WriteFile(path, []byte(configText), 0644)
	if err != nil {
		log.Fatalf("Cannot write to file %s: %s.", path, err)
	}
}
func versionFlag(cmd *cobra.Command) bool {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatal(err)
	}
	if version {
		fmt.Printf("\nversion: %s\n\nbuild:   %s\n\n",
			gnparser.Version, gnparser.Build)
		return true
	}
	return false
}

func formatFlag(cmd *cobra.Command) {
	f, err := cmd.Flags().GetString("format")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if f != "" {
		opts = append(opts, config.OptFormat(f))
	}
}

func jobsNumFlag(cmd *cobra.Command) {
	jn, err := cmd.Flags().GetInt("jobs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if jn > 0 {
		opts = append(opts, config.OptJobsNum(jn))
	}
}

func ignoreHTMLTagsFlag(cmd *cobra.Command) {
	ignoreTags, err := cmd.Flags().GetBool("ignore_tags")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if ignoreTags {
		opts = append(opts, config.OptIgnoreHTMLTags(true))
	}
}

func withDetailsFlag(cmd *cobra.Command) {
	withDet, err := cmd.Flags().GetBool("details")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if withDet {
		opts = append(opts, config.OptWithDetails(true))
	}
}

func batchSizeFlag(cmd *cobra.Command) {
	bs, err := cmd.Flags().GetInt("batch_size")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if bs > 0 {
		opts = append(opts, config.OptBatchSize(bs))
	}
}

func portFlag(cmd *cobra.Command) int {
	webPort, err := cmd.Flags().GetInt("port")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if webPort > 0 {
		opts = append(opts, config.OptPort(webPort))
	}
	return webPort
}

func processStdin(cmd *cobra.Command, cfg config.Config) {
	if !checkStdin() {
		_ = cmd.Help()
		return
	}
	gnp := gnparser.NewGNParser(cfg)
	parseFile(gnp, os.Stdin)
}

func checkStdin() bool {
	stdInFile := os.Stdin
	stat, err := stdInFile.Stat()
	if err != nil {
		log.Panic(err)
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

func parse(
	data string,
	cfg config.Config,
) {
	gnp := gnparser.NewGNParser(cfg)

	path := string(data)
	if fileExists(path) {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		parseFile(gnp, f)
		f.Close()
	} else {
		parseString(gnp, data)
	}
}

func fileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

func parseFile(
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
		if count == batchSize {
			i++
			log.Printf("Parsing %d-th line\n", count*i)
			chOut <- gnp.ParseNames(batch)
			batch = make([]string, batchSize)
			count = 0
		}
		batch[count] = sc.Text()
		count++
	}
	chOut <- gnp.ParseNames(batch[:count])
	close(chOut)
	wg.Wait()
}

func processResults(
	out <-chan []output.Parsed,
	wg *sync.WaitGroup,
	f format.Format,
) {
	defer wg.Done()
	if f == format.CSV {
		fmt.Println(output.CSVHeader())
	}
	for pr := range out {
		for i := range pr {
			fmt.Println(pr[i].Output(f))
		}
	}
}

func parseString(gnp gnparser.GNParser, name string) {
	res := gnp.ParseName(name)
	f := gnp.Format()
	if f == format.CSV {
		fmt.Println(output.CSVHeader())
	}
	fmt.Println(res.Output(f))
}
