package cmd

import (
	"os"

	"github.com/99designs/keyring"
	"github.com/spf13/cobra"
)

func openKeyring(cmd *cobra.Command) (keyring.Keyring, error) {
	backendFlag := cmd.Flag("backend").Value.String()
	filePath := cmd.Flag("file").Value.String()
	serviceName := cmd.Flag("service").Value.String()

	var allowedBackends []keyring.BackendType

	if backendFlag != "" {
		allowedBackends = []keyring.BackendType{keyring.BackendType(backendFlag)}
	} else {
		allowedBackends = keyring.AvailableBackends()
	}

	config := keyring.Config{
		ServiceName:     serviceName,
		AllowedBackends: allowedBackends,
	}

	if filePath != "" {
		config.FileDir = filePath
	}

	if len(allowedBackends) > 0 && allowedBackends[0] == keyring.KeychainBackend {
		config.KeychainName = getmacOSKeychainPath()
	}

	passwordEnv := os.Getenv("GOKEYRING_PASSWORD")
	if passwordEnv != "" {
		config.FilePasswordFunc = keyring.FixedStringPrompt(passwordEnv)
		config.KeychainPasswordFunc = func(prompt string) (string, error) {
			return passwordEnv, nil
		}
	} else {
		config.FilePasswordFunc = keyring.TerminalPrompt
		config.KeychainPasswordFunc = keyring.TerminalPrompt
	}

	return keyring.Open(config)
}
