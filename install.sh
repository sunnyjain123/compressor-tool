#!/usr/bin/env bash
set -e

BINARY_NAME=compress-go-tool
INSTALL_DIR=/usr/local/bin

echo "Installing $BINARY_NAME..."

if [ ! -f "$BINARY_NAME" ]; then
  echo "Binary not found. Build first using: go build -o compress"
  exit 1
fi

sudo mkdir -p $INSTALL_DIR
sudo install -m 0755 $BINARY_NAME $INSTALL_DIR/$BINARY_NAME

echo "Installed successfully!"
echo "Try: compress --help"