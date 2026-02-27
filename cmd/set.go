package cmd

import (
	"io"
	"os"

	"github.com/99designs/keyring"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <key> [value]",
	Short: "Store a secret in the keyring",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		if fromStdin {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			value = string(data)
		}

		ring, err := openKeyring(cmd)
		if err != nil {
			return err
		}

		return ring.Set(keyring.Item{
			Key:  key,
			Data: []byte(value),
		})
	},
}

var fromStdin bool

func init() {
	setCmd.Flags().BoolVar(&fromStdin, "from-stdin", false, "Read value from stdin")
	rootCmd.AddCommand(setCmd)
}
