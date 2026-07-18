#!/bin/bash
set -euo pipefail

VERSION="$1"

# Workspace modules (see go.work) that pin a specific github.com/fingerprintjs/go-sdk/v8
# version. Add a new entry here when a new workspace module starts depending on go-sdk/v8.
MODULES=(
  test
  functional_test
  example
  fingerprinttest
)

for module in "${MODULES[@]}"; do
  echo "Bumping github.com/fingerprintjs/go-sdk/v8 to ${VERSION} in ${module}"
  (
    cd "$module"
    go get "github.com/fingerprintjs/go-sdk/v8@${VERSION}"
    go mod tidy
  )
done
