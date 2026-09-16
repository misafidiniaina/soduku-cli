#!/bin/sh

set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
PREFIX=${PREFIX:-"$HOME/.local"}
BINDIR=${BINDIR:-"$PREFIX/bin"}
BINARY="$BINDIR/cligames"
TEMP_BINARY=$(mktemp "${TMPDIR:-/tmp}/cligames.XXXXXX")

cleanup() {
	rm -f "$TEMP_BINARY"
}
trap cleanup EXIT INT TERM

printf 'Building cligames...\n'
cd "$ROOT_DIR"
go build -o "$TEMP_BINARY" ./cmd/cligames

mkdir -p "$BINDIR"
install -m 755 "$TEMP_BINARY" "$BINARY"

printf 'Installed cligames to %s\n' "$BINARY"
case ":${PATH}:" in
	*":$BINDIR:"*) ;;
	*) printf 'Add %s to your PATH to run cligames from anywhere.\n' "$BINDIR" ;;
esac