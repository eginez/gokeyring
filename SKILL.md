---
name: "gokeyring"
description: "A CLI keyring tool for securely storing and retrieving passwords using system keyrings (macOS Keychain, Windows Credential Manager, Secret Service, etc.)"
license: "MIT"
compatibility: "opencode"
---

# gokeyring Skill

A CLI keyring tool for securely storing and retrieving passwords using system keyrings (macOS Keychain, Windows Credential Manager, Secret Service, etc.)

## Installation

```bash
make install
```

This installs `gokeyring` to `~/bin/gokeyring`.

Note: After running `make install`, ensure `~/bin` is in your PATH by adding `export PATH="$HOME/bin:$PATH"` to your shell profile (e.g., ~/.bashrc, ~/.zshrc).

## Basic Usage

### Store a secret
```bash
gokeyring set <key> [value]
```

If value is omitted, you will be prompted to enter it securely.

### Retrieve a secret
```bash
gokeyring get <key>
```

### List all keys
```bash
gokeyring list
```

### Delete a secret
```bash
gokeyring delete <key>
```

### Check status
```bash
gokeyring status
```

### Show version
```bash
gokeyring version
```

## Global Flags

| Flag | Description |
|------|-------------|
| `--backend` | Specify the backend to use (keychain for macOS, wincred for Windows, secret-service, file) |
| `--file` | Specify the file path for file-based backend |
| `--service` | Specify the service name for the keyring |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GOKEYRING_BACKEND` | Set the default backend |
| `GOKEYRING_PASSWORD` | Set the password for the keyring |
