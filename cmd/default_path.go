package cmd

import (
	"os"
	"path/filepath"
)

func getmacOSKeychainPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "gokeyring.keychain"
	}
	configDir := filepath.Join(homeDir, ".config", "keyring")
	return filepath.Join(configDir, "gokeyring.keychain")
}
