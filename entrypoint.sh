#!/bin/sh
set -e

ARCH=$(uname -m)
case "$ARCH" in
    x86_64 | amd64)
        exec ./notebase "$@"
        ;;
    aarch64 | arm64)
        exec ./notebase-arm64 "$@"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac
