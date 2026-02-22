#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

cleanup() {
  rm -rf rootKey workKey tools
}
trap cleanup EXIT

printf "[1/5] build...\n"
make build

printf "[2/5] generate root key...\n"
./tools key-gen --type root --force

printf "[3/5] generate work key...\n"
./tools key-gen --type work --name work.key --force

printf "[4/5] encrypt/decrypt...\n"
plain="hello-world-$(date +%s)"
cipher=$(./tools encrypt --work-key work.key "$plain")
if [[ -z "$cipher" ]]; then
  echo "encrypt output is empty" >&2
  exit 1
fi

decoded=$(./tools decrypt --work-key work.key "$cipher")
if [[ "$decoded" != "$plain" ]]; then
  echo "decrypt mismatch: expected '$plain', got '$decoded'" >&2
  exit 1
fi

printf "[5/5] ok\n"
