package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in the keyring",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		filter := cmd.Flag("filter").Value.String()
		jsonOutput := cmd.Flag("json").Value.String() == "true"

		ring, err := openKeyring(cmd)
		if err != nil {
			return err
		}

		keys, err := ring.Keys()
		if err != nil {
			return err
		}

		var filteredKeys []string
		for _, key := range keys {
			if filter == "" || strings.HasPrefix(key, filter) {
				filteredKeys = append(filteredKeys, key)
			}
		}

		if jsonOutput {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(filteredKeys)
		}

		for _, key := range filteredKeys {
			fmt.Println(key)
		}

		return nil
	},
}

func init() {
	listCmd.Flags().String("filter", "", "Filter keys by prefix")
	listCmd.Flags().String("json", "false", "Output in JSON format")
	rootCmd.AddCommand(listCmd)
}
