#!/bin/bash
set -e

PROJECT_ROOT="/Users/rahulsingh/Documents/GitHub/coral-eams-client/Client"

echo "Building macOS binary..."
GOOS=darwin GOARCH=amd64 go build -o coral-eams-client-mac $PROJECT_ROOT/cmd/coral-eams-client

echo "Building Linux binary..."
GOOS=linux GOARCH=amd64 go build -o coral-eams-client-linux $PROJECT_ROOT/cmd/coral-eams-client

echo "Building Windows binary..."
GOOS=windows GOARCH=amd64 go build -o coral-eams-client-windows.exe $PROJECT_ROOT/cmd/coral-eams-client

echo "All binaries built successfully!"
