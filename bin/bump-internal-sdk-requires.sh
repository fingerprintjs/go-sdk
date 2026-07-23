#!/bin/bash
set -euo pipefail

VERSION="$1"

if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "error: invalid version '${VERSION}' (expected format vX.Y.Z)" >&2
  exit 1
fi

# Workspace modules (see go.work) that pin a specific github.com/fingerprintjs/go-sdk/v8
# version, excluding the root go-sdk module itself.
ROOT_DIR="$(pwd)"
MODULES=()
while IFS= read -r dir; do
  [[ "$dir" == "$ROOT_DIR" ]] && continue
  MODULES+=("$(basename "$dir")")
done < <(go list -m -f '{{.Dir}}')

for module in "${MODULES[@]}"; do
  echo "Bumping github.com/fingerprintjs/go-sdk/v8 to ${VERSION} in ${module}"
  (
    cd "$module"
    go get "github.com/fingerprintjs/go-sdk/v8@${VERSION}"
    go mod tidy
  )
done
