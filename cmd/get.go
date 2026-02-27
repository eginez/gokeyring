package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Retrieve a secret from the keyring",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		ring, err := openKeyring(cmd)
		if err != nil {
			return err
		}

		item, err := ring.Get(key)
		if err != nil {
			return err
		}

		fmt.Print(string(item.Data))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
