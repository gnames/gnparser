package cmd

import (
	"fmt"
	"log/slog"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser"
	"github.com/spf13/cobra"
)

func batchSizeFlag(cmd *cobra.Command) {
	bs, _ := cmd.Flags().GetInt("batch_size")
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
	jn, _ := cmd.Flags().GetInt("jobs")
	if jn > 0 {
		opts = append(opts, gnparser.OptJobsNum(jn))
	}
}

func ignoreHTMLTagsFlag(cmd *cobra.Command) {
	ignoreTags, _ := cmd.Flags().GetBool("ignore_tags")
	if ignoreTags {
		opts = append(opts, gnparser.OptIgnoreHTMLTags(true))
	}
}

func portFlag(cmd *cobra.Command) int {
	webPort, _ := cmd.Flags().GetInt("port")
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
	b, _ := cmd.Flags().GetBool("capitalize")
	if b {
		opts = append(opts, gnparser.OptWithCapitaliation(true))
	}
}

func withDetailsFlag(cmd *cobra.Command) {
	withDet, _ := cmd.Flags().GetBool("details")
	if withDet {
		opts = append(opts, gnparser.OptWithDetails(true))
	}
}

func withEnableCultivarsFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("cultivar")
	if b {
		opts = append(opts, gnparser.OptCode(nomcode.Cultivars))
	}
}

func withNoOrderFlag(cmd *cobra.Command) {
	withOrd, _ := cmd.Flags().GetBool("unordered")
	if withOrd {
		opts = append(opts, gnparser.OptWithNoOrder(true))
	}
}

func withPreserveDiaeresesFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("diaereses")
	if b {
		opts = append(opts, gnparser.OptWithPreserveDiaereses(true))
	}
}

func withCompactAuthorsFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("compact-authors")
	if b {
		opts = append(opts, gnparser.OptWithCompactAuthors(true))
	}
}

func withFlatOutputFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("flatten-output")
	if b {
		opts = append(opts, gnparser.OptWithFlatOutput(true))
	}
}

func spGrCutFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("species-group-cut")
	if b {
		opts = append(opts, gnparser.OptWithSpeciesGroupCut(true))
	}
}

func withStreamFlag(cmd *cobra.Command) {
	withDet, _ := cmd.Flags().GetBool("stream")
	if withDet {
		opts = append(opts, gnparser.OptWithStream(true))
	}
}
