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

```bash
aka list
```

### Export aliases

```bash
aka export aliases.json
```

### Import aliases

```bash
aka import aliases.json
```
