package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/nomcode"
	"github.com/spf13/cobra"
)

func batchSizeFlag(cmd *cobra.Command) {
	bs, err := cmd.Flags().GetInt("batch_size")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if bs > 0 {
		opts = append(opts, gnparser.OptBatchSize(bs))
	}
}

func codeFlag(cmd *cobra.Command) {
	s, _ := cmd.Flags().GetString("nomenclatural-code")
	if s == "" {
		return
	}
	code := nomcode.New(s)
	if code == nomcode.Unknown && s != "any" {
		slog.Warn("Cannot determine nomenclatural-code from input", "input", s)
	}
	opts = append(opts, gnparser.OptCode(code))
}

func formatFlag(cmd *cobra.Command) {
	s, _ := cmd.Flags().GetString("format")
	if s != "" {
		frmt, err := gnfmt.NewFormat(s)
		if err != nil {
			slog.Warn("Unknown format input, using default: CSV", "inut", s)
			frmt = gnfmt.CSV
		}
		opts = append(opts, gnparser.OptFormat(frmt))
	}
}

func jobsNumFlag(cmd *cobra.Command) {
	jn, err := cmd.Flags().GetInt("jobs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if jn > 0 {
		opts = append(opts, gnparser.OptJobsNum(jn))
	}
}

func ignoreHTMLTagsFlag(cmd *cobra.Command) {
	ignoreTags, err := cmd.Flags().GetBool("ignore_tags")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if ignoreTags {
		opts = append(opts, gnparser.OptIgnoreHTMLTags(true))
	}
}

func portFlag(cmd *cobra.Command) int {
	webPort, err := cmd.Flags().GetInt("port")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if webPort > 0 {
		opts = append(opts, gnparser.OptPort(webPort))
	}
	return webPort
}

func versionFlag(cmd *cobra.Command) bool {
	version, _ := cmd.Flags().GetBool("version")
	if version {
		fmt.Printf("\nversion: %s\n\nbuild:   %s\n\n",
			gnparser.Version, gnparser.Build)
		return true
	}
	return false
}

func withCapitalizeFlag(cmd *cobra.Command) {
	b, err := cmd.Flags().GetBool("capitalize")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if b {
		opts = append(opts, gnparser.OptWithCapitaliation(true))
	}
}

func withDetailsFlag(cmd *cobra.Command) {
	withDet, err := cmd.Flags().GetBool("details")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if withDet {
		opts = append(opts, gnparser.OptWithDetails(true))
	}
}

func withEnableCultivarsFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("cultivar")
	if b {
		opts = append(opts, gnparser.OptCode(nomcode.Cultivar))
	}
}

func withNoOrderFlag(cmd *cobra.Command) {
	withOrd, err := cmd.Flags().GetBool("unordered")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if withOrd {
		opts = append(opts, gnparser.OptWithNoOrder(true))
	}
}

func withPreserveDiaeresesFlag(cmd *cobra.Command) {
	b, err := cmd.Flags().GetBool("diaereses")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if b {
		opts = append(opts, gnparser.OptWithPreserveDiaereses(true))
	}
}

func spGrCutFlag(cmd *cobra.Command) {
	b, err := cmd.Flags().GetBool("species-group-cut")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if b {
		opts = append(opts, gnparser.OptWithSpeciesGroupCut(true))
	}

}

func withStreamFlag(cmd *cobra.Command) {
	withDet, err := cmd.Flags().GetBool("stream")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if withDet {
		opts = append(opts, gnparser.OptWithStream(true))
	}
}

func withWebLogsFlag(cmd *cobra.Command) bool {
	withLogs, err := cmd.Flags().GetBool("web-logs")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return withLogs
}
