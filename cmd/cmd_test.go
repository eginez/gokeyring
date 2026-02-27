package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetmacOSKeychainPath(t *testing.T) {
	path := getmacOSKeychainPath()
	expectedSuffix := "gokeyring.keychain"

	if len(path) < len(expectedSuffix) {
		t.Fatalf("Path too short: %s", path)
	}

	if path[len(path)-len(expectedSuffix):] != expectedSuffix {
		t.Fatalf("Expected path to end with %s, got %s", expectedSuffix, path)
	}
}

func TestRootCommand(t *testing.T) {
	cmd := rootCmd
	if cmd.Use != "gokeyring" {
		t.Fatalf("Expected use 'gokeyring', got '%s'", cmd.Use)
	}
}

func TestGetCommand(t *testing.T) {
	cmd := getCmd
	if cmd.Use != "get <key>" {
		t.Fatalf("Expected use 'get <key>', got '%s'", cmd.Use)
	}
}

func TestSetCommand(t *testing.T) {
	cmd := setCmd
	if cmd.Use != "set <key> [value]" {
		t.Fatalf("Expected use 'set <key> [value]', got '%s'", cmd.Use)
	}
}

func TestDeleteCommand(t *testing.T) {
	cmd := deleteCmd
	if cmd.Use != "delete <key>" {
		t.Fatalf("Expected use 'delete <key>', got '%s'", cmd.Use)
	}
}

func TestListCommand(t *testing.T) {
	cmd := listCmd
	if cmd.Use != "list" {
		t.Fatalf("Expected use 'list', got '%s'", cmd.Use)
	}
}

func TestStatusCommand(t *testing.T) {
	cmd := statusCmd
	if cmd.Use != "status" {
		t.Fatalf("Expected use 'status', got '%s'", cmd.Use)
	}
}

func TestOpenKeyringWithBackend(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("backend", "", "")
	cmd.Flags().String("file", "", "")
	cmd.Flags().String("service", "gokeyring", "")

	ring, err := openKeyring(cmd)
	if err != nil {
		t.Logf("Expected failure on some platforms: %v", err)
	} else if ring != nil {
		t.Logf("Opened keyring successfully")
	}
}

func TestGetmacOSKeychainPathEnv(t *testing.T) {
	os.Setenv("GOKEYRING_PASSWORD", "testpassword")
	path := getmacOSKeychainPath()
	if len(path) == 0 {
		t.Fatalf("Path should not be empty")
	}
	os.Setenv("GOKEYRING_PASSWORD", "")
}
