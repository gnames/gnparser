package cmd

import (
	"fmt"
	"os"

	"github.com/gnames/gnparser"
	"github.com/rs/zerolog/log"
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

func formatFlag(cmd *cobra.Command) {
	f, err := cmd.Flags().GetString("format")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if f != "" {
		opts = append(opts, gnparser.OptFormat(f))
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
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatal().Err(err)
	}
	if version {
		fmt.Printf("\nversion: %s\n\nbuild:   %s\n\n",
			gnparser.Version, gnparser.Build)
		return true
	}
	return false
}

func webLogsNsqdTCPFlag(cmd *cobra.Command) string {
	u, err := cmd.Flags().GetString("nsqd-tcp")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return u
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
	b, err := cmd.Flags().GetBool("cultivar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if b {
		opts = append(opts, gnparser.OptWithCultivars(true))
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
