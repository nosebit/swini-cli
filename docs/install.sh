#!/bin/sh

# GitHub repository information
REPO="nosebit/swini-cli"

# Installation directory
INSTALL_DIR="$HOME/.local/bin"

# Determine the OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Adjust the architecture string for different platforms
if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
elif [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64"
fi

echo "[*] Detected OS: ${OS}"
echo "[*] Detected Architecture: ${ARCH}"

# Get the latest tag name from the GitHub API using awk
VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | \
  awk -F'"' '/"tag_name":/ {print $4}')

if [ -z "$VERSION" ]; then
  echo "[!] Error: Could not retrieve the latest version tag from GitHub."
  exit 1
fi

RAW_VERSION="${VERSION#v}"

echo "[*] Latest version found: ${VERSION}"

# Define the full asset filename based on version, OS, and architecture
FILENAME="swini-${RAW_VERSION}-${OS}-${ARCH}.tar.gz"

# Construct the full download URL
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "[*] Downloading from: ${DOWNLOAD_URL}"

# Create the installation directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Download the asset into a temporary directory
curl -s -L $DOWNLOAD_URL | tar -xz -C $INSTALL_DIR

# Ensure the install directory is in the user's PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "Adding $INSTALL_DIR to your PATH in your shell configuration file..."
  SHELL_CONFIG=""
  if [ -f "$HOME/.zshrc" ]; then
    SHELL_CONFIG="$HOME/.zshrc"
  elif [ -f "$HOME/.bashrc" ]; then
    SHELL_CONFIG="$HOME/.bashrc"
  elif [ -f "$HOME/.bash_profile" ]; then
    SHELL_CONFIG="$HOME/.bash_profile"
  fi

  if [ -n "$SHELL_CONFIG" ]; then
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_CONFIG"
    echo "[!] Please restart your terminal or run 'source $SHELL_CONFIG' to update your PATH."
  else
    echo "[!] Could not find a shell configuration file (.zshrc or .bashrc)."
    echo "[!] Please add the following line to your shell profile to make 'swini-cli' executable:"
    echo "[>] export PATH=\"\$PATH:$INSTALL_DIR\""
  fi
fi

echo "[✔] swini has been successfully installed to ${INSTALL_DIR}."
