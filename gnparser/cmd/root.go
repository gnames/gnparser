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
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"gitlab.com/gogna/gnparser"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gnparser file_or_name",
	Short: "Parses scientific names into their semantic elements.",
	Long: `
Parses scientific names into their semantic elements.

To see version:
gnparser -v

To parse one name:
gnparser "Homo sapiens Linnaeus 1753" [flags]

To parse many names from a file (one name per line):
gnparser names.txt [flags] > parsed_names.txt
 `,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag(cmd)
		f := formatFlag(cmd)
		wn := workersNumFlag(cmd)
		opts := []gnparser.Option{
			gnparser.WorkersNum(wn),
			gnparser.Format(f),
		}
		data := getInput(cmd, args)
		parse(data, opts)
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
	gnp := gnparser.NewGNparser()
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Show build version and date")

	df := gnp.OutputFormat()
	formats := strings.Join(gnparser.AvailableFormats(), ", ")
	formatHelp := fmt.Sprintf("Output format. Can be one of:\n %s", formats)
	rootCmd.Flags().StringP("format", "f", df, formatHelp)

	dj := gnp.WorkersNum()
	rootCmd.Flags().IntP("jobs", "j", dj,
		"Nubmer of threads to run. CPU's threads number is the default.")
}

func versionFlag(cmd *cobra.Command) {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if version {
		fmt.Printf("\nversion: %s\n\nbuild:   %s\n\n",
			gnparser.Version(), gnparser.Build())
		os.Exit(0)
	}
}

func formatFlag(cmd *cobra.Command) string {
	str, err := cmd.Flags().GetString("format")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return str
}

func workersNumFlag(cmd *cobra.Command) int {
	i, err := cmd.Flags().GetInt("jobs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
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
	case 0:
		if !checkStdin() {
			cmd.Help()
			os.Exit(0)
		}
		bs, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Println(err)
		}
		data = string(bs)
	case 1:
		data = args[0]
	case 2:
		data = args[1]
	default:
		cmd.Help()
		os.Exit(0)
	}
	return data
}

func parse(data string, opts []gnparser.Option) {
	gnp := gnparser.NewGNparser(opts...)

	path := string(data)
	if fileExists(path) {
		parseFile(path, opts)
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

func parseFile(path string, opts []gnparser.Option) {
	in := make(chan string)
	out := make(chan *gnparser.ParseResult)
	gnp := gnparser.NewGNparser(opts...)
	var wg sync.WaitGroup
	wg.Add(1)

	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	go gnp.ParseStream(in, out, opts...)
	go processResults(out, &wg)
	sc := bufio.NewScanner(f)
	count := 0
	for sc.Scan() {
		count++
		if count%50000 == 0 {
			log.Printf("Parsing %d-th line\n", count)
		}
		name := sc.Text()
		in <- name
	}
	close(in)
	wg.Wait()
}

func processResults(out <-chan *gnparser.ParseResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range out {
		if r.Error != nil {
			log.Println(r.Error)
		}
		fmt.Println(r.Output)
	}
}

func parseString(gnp gnparser.GNparser, data string) {
	res, err := gnp.ParseAndFormat(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
