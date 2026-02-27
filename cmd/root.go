package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gokeyring",
	Short: "A CLI keyring tool for secure password storage",
	Long: `gokeyring is a CLI tool for securely storing and retrieving passwords
using system keyrings (macOS Keychain, Windows Credential Manager, etc.)`,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("backend", "", "Backend to use (keychain, file, secret-service, kwallet, keyctl, pass)")
	rootCmd.PersistentFlags().String("file", "", "Path to secrets file (file backend)")
	rootCmd.PersistentFlags().String("service", "gokeyring", "Service name for keyring")
}
