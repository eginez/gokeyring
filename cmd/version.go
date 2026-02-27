package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print gokeyring version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if Version == "" {
			Version = "dev"
		}
		if Commit == "" {
			Commit = "none"
		}
		if Date == "" {
			Date = "unknown"
		}
		fmt.Printf("gokeyring version %s (%s) built on %s\n", Version, Commit, Date)
	},
}

var (
	Version string
	Commit  string
	Date    string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
