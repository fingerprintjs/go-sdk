#!/bin/bash

set -e

SDK_DIR="v8"
PACKAGE_VERSION=$(node -p "require('./package.json').version")

prepare_directory(){
  rm -rf $SDK_DIR
  mkdir -p $SDK_DIR
  cp .openapi-generator-ignore $SDK_DIR
}

run_generator() {
  docker run --rm -u "$(id -u):$(id -g)" -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.19.0 generate \
      -c /local/openapi-config.yml \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -t /local/template \
      --additional-properties=packageVersion=$PACKAGE_VERSION
      # Note: 'packageVersion' is not configured via config file because we need to inject its value at runtime
      # Please use the config file for all other configuration
}

prepare_directory
run_generator
gofmt -w .

# Replace "Id" with "ID" in all generated .go and .md files, skipping links
find "$SDK_DIR" \( -name "*.go" -o -name "*.md" \) | while read -r file; do
  perl -i -pe 's/(?<![(\[#\/])([a-zA-Z])Id(?=[A-Z_\s\W]|$)/$1ID/g' "$file"
done