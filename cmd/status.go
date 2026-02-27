package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show keyring configuration",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		backendFlag := cmd.Flag("backend").Value.String()
		filePath := cmd.Flag("file").Value.String()

		fmt.Println("gokeyring Configuration")
		fmt.Println("")

		var backend string
		if backendFlag != "" {
			backend = backendFlag
		} else {
			backend = "auto"
		}
		fmt.Printf("Backend: %s\n", backend)

		var source string
		if backendFlag != "" {
			source = "cli-flag"
		} else if os.Getenv("GOKEYRING_BACKEND") != "" {
			source = "env"
		} else {
			source = "auto-detect"
		}
		fmt.Printf("Backend Source: %s\n", source)

		var path string
		if filePath != "" {
			path = filePath
		} else if runtime.GOOS == "darwin" {
			path = getmacOSKeychainPath()
		} else {
			path = "~/"
		}
		fmt.Printf("Keyring Path: %s\n", path)

		if runtime.GOOS == "darwin" {
			fmt.Println("Keychain Status: N/A (unlock command available)")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
