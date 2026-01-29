#!/bin/bash

set -e

SDK_DIR="sdk"

prepare_directory(){
  rm -rf $SDK_DIR
  mkdir -p $SDK_DIR
  cp .openapi-generator-ignore $SDK_DIR
}

run_generator() {
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.19.0 generate \
      -c /local/openapi-config.yml \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -t /local/template
}

override_version_file() {
  local pkg_json="package.json"
  local out_file="${SDK_DIR}/version.go"

  if [[ ! -f "$pkg_json" ]]; then
    echo "Error: $pkg_json not found"
    return 1
  fi

  local version
  version=$(jq -r '.version' "$pkg_json") || {
    echo "Error: Failed to parse version from $pkg_json"
    return 1
  }

  if [[ -z "$version" ]]; then
    echo "Error: version in $pkg_json is empty"
    return 1
  fi

  if sed --version >/dev/null 2>&1; then
    # GNU sed (Linux or GitHub runners)
    sed -i 's/^const Version = "0.0.0"/const Version = "'"$version"'"/' $out_file
  else
    # BSD sed (macOS)
    sed -i '' 's/^const Version = "0.0.0"/const Version = "'"$version"'"/' $out_file
  fi

  echo "Updated Version constant in ${out_file} to ${version}"
}

prepare_directory
run_generator
override_version_file

gofmt -w .
