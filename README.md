# aka

A simple alias manager that lets you add, list, and apply command aliases.

## Installation

### Using curl (recommended)

```sh
curl -fsSL https://raw.githubusercontent.com/fdorantesm/aka/refs/heads/main/install.sh | bash
```

### Using npm

```sh
npm install -g aka-cli
```

### Using yarn

```sh
yarn global add aka-cli
```

### Using pnpm

```sh
pnpm add -g aka-cli
```

### Using bun

```sh
bun add -g aka-cli
```

### Using deno

```sh
deno install -A -n aka npm:aka-cli
```

### Using npmx.dev

```sh
npmx aka-cli
```

## Usage

### Help

```bash
aka [command]
```

### Add alias

```bash
aka add ll "ls -la"
```

### List

List all aliases:

```bash
aka list
```

Filter aliases with glob patterns:

```bash
aka list '*dev*'    # Aliases containing "dev"
aka list 'aws*'     # Aliases starting with "aws"
aka list '*-qa'     # Aliases ending with "-qa"
```

### Echo alias command

Print the command for a specific alias:

```bash
aka echo ll
# Output: ls -la
```

Copy to clipboard:

```bash
# macOS
aka echo ll | pbcopy

# Linux
aka echo ll | xclip -selection clipboard
```

### Update alias

Update an existing alias interactively:

```bash
aka update ll
# Current command: ls -la
# Enter new command: ls -lah --color=auto
```

### Rename alias

Rename an existing alias:

```bash
aka rename gs gst
```

### Remove alias

```bash
aka remove ll
```

### Export aliases

```bash
aka export aliases.json
```

### Import aliases

```bash
aka import aliases.json
```

## Contributing

Contributions are welcome! Please see [DEPLOYMENT.md](docs/DEPLOYMENT.md) for information about the release and deployment process.

## License

MIT
