# aka

A simple alias manager that lets you add, list, and apply command aliases.

## Installation

To install `aka`, run the following command:

```sh
curl -fsSL https://raw.githubusercontent.com/fdorantesm/aka/refs/heads/main/install.sh | bash
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
