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
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"gitlab.com/gogna/gnparser"
	"gitlab.com/gogna/gnparser/preprocess"
	"gitlab.com/gogna/gnparser/rpc"
	"gitlab.com/gogna/gnparser/web"
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

To clean names from html tags and entities
gnparser names.txt -c > cleanded_names.txt

To start gRPC parsing service on port 3355 with a limit
of 10 concurrent jobs per request:
gnparser -j 10 -g 3355

To start web service on port 8080 with 5 concurrent jobs:
gnparser -j 5 -g 8080
 `,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag(cmd)
		wn := workersNumFlag(cmd)

		cleanup := cleanupFlag(cmd)

		grpcPort := grpcFlag(cmd)
		if grpcPort != 0 {
			fmt.Println("Running gnparser as gRPC service:")
			fmt.Printf("port: %d\n", grpcPort)
			fmt.Printf("Max jobs per request: %d\n\n", wn)
			rpc.Run(grpcPort, wn)
			os.Exit(0)
		}

		webPort := webFlag(cmd)
		if webPort != 0 {
			fmt.Println("Running gnparser as a website and REST server:")
			fmt.Printf("port: %d\n", webPort)
			fmt.Printf("jobs: %d\n\n", wn)
			web.Run(webPort, wn)
			os.Exit(0)
		}
		f := formatFlag(cmd)
		opts := []gnparser.Option{
			gnparser.WorkersNum(wn),
			gnparser.Format(f),
		}
		if len(args) == 0 {
			processStdin(cmd, cleanup, wn, opts)
			os.Exit(0)
		}
		data := getInput(cmd, args)
		if cleanup {
			cleanupData(data, wn)
			os.Exit(0)
		}
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
	rootCmd.PersistentFlags().BoolP("version", "v", false, "shows build version and date, ignores other flags.")

	df := gnp.OutputFormat()
	formats := strings.Join(gnparser.AvailableFormats(), ", ")
	formatHelp := fmt.Sprintf("sets output format. Can be one of:\n %s.", formats)
	rootCmd.Flags().StringP("format", "f", df, formatHelp)

	dj := gnp.WorkersNum()
	rootCmd.Flags().IntP("jobs", "j", dj,
		"nubmer of threads to run. CPU's threads number is the default.")

	rootCmd.Flags().BoolP("cleanup", "c", false, "removes HTML entities and tags instead of parsing.")

	rootCmd.Flags().IntP("grpc_port", "g", 0, "starts gRPC server on the port.")

	rootCmd.Flags().IntP("web_port", "w", 0,
		"starts web site and REST server on the port.")
}

func versionFlag(cmd *cobra.Command) {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if version {
		var gnp *gnparser.GNparser
		fmt.Printf("\nversion: %s\n\nbuild:   %s\n\n",
			gnp.Version(), gnp.Build())
		os.Exit(0)
	}
}

func cleanupFlag(cmd *cobra.Command) bool {
	cleanup, err := cmd.Flags().GetBool("cleanup")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cleanup
}

func grpcFlag(cmd *cobra.Command) int {
	grpcPort, err := cmd.Flags().GetInt("grpc_port")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return grpcPort
}

func webFlag(cmd *cobra.Command) int {
	webPort, err := cmd.Flags().GetInt("web_port")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return webPort
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

func processStdin(cmd *cobra.Command, cleanup bool, wn int,
	opts []gnparser.Option) {
	if !checkStdin() {
		cmd.Help()
		return
	}
	if cleanup {
		cleanupFile(os.Stdin, wn)
		return
	}
	parseFile(os.Stdin, opts)
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
		cmd.Help()
		os.Exit(0)
	}
	return data
}

func parse(data string, opts []gnparser.Option) {
	gnp := gnparser.NewGNparser(opts...)

	path := string(data)
	if fileExists(path) {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		parseFile(f, opts)
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

func parseFile(f io.Reader, opts []gnparser.Option) {
	in := make(chan string)
	out := make(chan *gnparser.ParseResult)
	gnp := gnparser.NewGNparser(opts...)
	var wg sync.WaitGroup
	wg.Add(1)

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

func cleanupData(data string, wc int) {
	path := string(data)
	if fileExists(path) {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		cleanupFile(f, wc)
		f.Close()
	} else {
		res := preprocess.StripTags(data)
		fmt.Println(data + "|" + res)
	}
}

func cleanupFile(f io.Reader, wn int) {
	in := make(chan string)
	out := make(chan *preprocess.CleanupResult)
	var wg sync.WaitGroup
	wg.Add(1)

	go preprocess.CleanupStream(in, out, wn)
	go processCleanup(out, &wg)
	sc := bufio.NewScanner(f)
	count := 0
	for sc.Scan() {
		count++
		if count%1000000 == 0 {
			log.Printf("Cleaning %d-th line\n", count)
		}
		name := sc.Text()
		in <- name
	}
	close(in)
	wg.Wait()
}

func processCleanup(out <-chan *preprocess.CleanupResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range out {
		fmt.Printf("%s|%s", r.Input, r.Output)
	}
}
