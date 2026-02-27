package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a secret from the keyring",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		ring, err := openKeyring(cmd)
		if err != nil {
			return err
		}

		if !force {
			fmt.Printf("Are you sure you want to delete '%s'? [y/N] ", key)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				return nil
			}
		}

		return ring.Remove(key)
	},
}

var force bool

func init() {
	deleteCmd.Flags().BoolVar(&force, "force", false, "Skip confirmation prompt")
	rootCmd.AddCommand(deleteCmd)
}
