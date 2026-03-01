#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

cleanup() {
  rm -rf "$KEY_DIR" tools
}
trap cleanup EXIT

KEY_DIR="$ROOT_DIR/.keys-self-check"

printf "[1/5] build...\n"
make build

printf "[2/5] generate root key...\n"
./tools gen-root-key --force --dir "$KEY_DIR"

printf "[3/5] generate work key...\n"
./tools gen-work-key --name work.key --force --dir "$KEY_DIR"

printf "[4/5] encrypt/decrypt...\n"
plain="hello-world-$(date +%s)"
cipher=$(./tools encrypt --work-key work.key --key-dir "$KEY_DIR" "$plain")
if [[ -z "$cipher" ]]; then
  echo "encrypt output is empty" >&2
  exit 1
fi

decoded=$(./tools decrypt --work-key work.key --key-dir "$KEY_DIR" "$cipher")
if [[ "$decoded" != "$plain" ]]; then
  echo "decrypt mismatch: expected '$plain', got '$decoded'" >&2
  exit 1
fi

printf "[5/5] ok\n"
