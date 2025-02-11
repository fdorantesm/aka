#!/bin/sh
set -e

# ====================================
# aka Installer Script
# ------------------------------------
# This script downloads the aka binary (compiled from Go)
# and installs it into $HOME/bin.
#
# It also creates the configuration directory $HOME/.aka
# to store alias files.
#
# Finally, it appends the following lines to known shell configuration
# files (if they exist) so that your PATH is updated and the alias
# definitions (from `aka apply`) are sourced:
#
#   export PATH="$HOME/bin:$PATH"
#   source <(aka apply)
#
# If none of the known files are found, it will ask you to add them manually.
#
# NOTE: Update the AKA_URL variable below with the actual URL where your
# binary is hosted (for example, for Linux/amd64).
# ====================================

VERSION="v1.0.2"
if [ "$(uname)" = "Darwin" ]; then
  AKA_URL="https://github.com/fdorantesm/aka/releases/download/$VERSION/aka-darwin"
else
  AKA_URL="https://github.com/fdorantesm/aka/releases/download/$VERSION/aka-linux"
fi

echo "Installing aka..."

# Create $HOME/bin if it doesn't exist.
if [ ! -d "$HOME/bin" ]; then
  mkdir -p "$HOME/bin"
fi

# Create $HOME/.aka if it doesn't exist.
if [ ! -d "$HOME/.aka" ]; then
  mkdir -p "$HOME/.aka"
fi

# Download the aka binary using curl or wget.
if command -v curl >/dev/null 2>&1; then
  curl -L "$AKA_URL" -o "$HOME/bin/aka"
elif command -v wget >/dev/null 2>&1; then
  wget -O "$HOME/bin/aka" "$AKA_URL"
else
  echo "Error: curl or wget is required to download aka."
  exit 1
fi

# Make the binary executable.
chmod a+x "$HOME/bin/aka"

echo "aka was installed successfully to $HOME/bin/aka"
echo ""

# List of known configuration files to try updating.
CONFIG_FILES="$HOME/.zshrc $HOME/.bashrc $HOME/.profile $HOME/.bash_profile"
UPDATED=0

for file in $CONFIG_FILES; do
  if [ -f "$file" ]; then
    echo "" >> "$file"
    echo "# Added by aka installer" >> "$file"
    echo 'export PATH="$HOME/bin:$PATH"' >> "$file"
    echo 'source <(aka apply)' >> "$file"
    echo "" >> "$file"
    echo "Updated $file with the required lines."
    UPDATED=1
  fi
done

if [ $UPDATED -eq 0 ]; then
  echo "No known shell configuration file found."
  echo "Please add the following lines manually to your shell profile:"
  echo ""
  echo 'export PATH="$HOME/bin:$PATH"'
  echo 'source <(aka apply)'
  echo ""
fi

echo ""
echo "To complete the installation, restart your terminal or run:"
echo "exec \$SHELL"
