package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print gokeyring version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Get git commit hash
		commit, _ := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		if len(commit) > 0 {
			fmt.Printf("gokeyring version %s\n", string(commit))
		} else {
			fmt.Println("gokeyring version dev (not built from git)")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
