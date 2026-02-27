//go:build darwin

package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock macOS keychain",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getmacOSKeychainPath()

		unlockCmd := exec.Command("security", "unlock-keychain", path)
		unlockCmd.Stdin = os.Stdin
		unlockCmd.Stdout = os.Stdout
		unlockCmd.Stderr = os.Stderr

		return unlockCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(unlockCmd)
}
