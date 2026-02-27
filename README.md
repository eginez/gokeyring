# gokeyring

A CLI keyring tool for securely storing and retrieving passwords using system keyrings.

## Features

- **Multiple backends**: Automatically detects and uses macOS Keychain, Windows Credential Manager, Secret Service, KWallet, KeyCtl, Pass, or encrypted file backend
- **Simple CLI**: Easy-to-use command line interface
- **Cross-platform**: Works on macOS, Linux, and Windows
- **Flexible configuration**: Configure backend, file location, and service name via CLI flags or environment variables

## Installation

### Build from source

```bash
make build
```

### Install

```bash
make install
```

## Usage

```
gokeyring [command]
```

### Commands

| Command | Description |
|---------|-------------|
| `get <key>` | Retrieve a secret from the keyring |
| `set <key> [value]` | Store a secret in the keyring |
| `delete <key>` | Delete a secret from the keyring |
| `list` | List all keys in the keyring |
| `status` | Show keyring configuration |
| `unlock` | Unlock macOS keychain (macOS only) |
| `version` | Print gokeyring version |

### Global Flags

| Flag | Description |
|------|-------------|
| `--backend string` | Backend to use (keychain, file, secret-service, kwallet, keyctl, pass) |
| `--file string` | Path to secrets file/directory (file backend) |
| `--service string` | Service name for keyring (default "gokeyring") |

## Examples

### Store a secret

```bash
# Interactive prompt
gokeyring set mykey

# From command line
gokeyring set mykey "mysecret"

# From stdin
echo "mysecret" | gokeyring set mykey --from-stdin
```

### Retrieve a secret

```bash
gokeyring get mykey
```

### List all keys

```bash
gokeyring list

# With filter
gokeyring list --filter prefix

# JSON output
gokeyring list --json
```

### Delete a secret

```bash
gokeyring delete mykey

# Force delete without prompt
gokeyring delete mykey --force
```

### Configure with file backend

```bash
# Store secrets in a custom directory
gokeyring set --backend file --file ~/my-secrets key value

# With password from environment
GOKEYRING_PASSWORD="mypass" gokeyring set --backend file --file ~/my-secrets key value
```

### View configuration

```bash
gokeyring status
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GOKEYRING_BACKEND` | Force a specific backend |
| `GOKEYRING_PASSWORD` | Password for non-interactive mode (file backend) |

## Testing

```bash
make test
```

## License

MIT
