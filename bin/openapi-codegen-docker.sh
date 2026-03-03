#!/bin/bash

set -e

SDK_DIR="internal/openapi"
PACKAGE_VERSION=$(node -p "require('./package.json').version")

prepare_directory(){
  rm -rf $SDK_DIR
  rm -rf docs
  mkdir -p $SDK_DIR
  cp .openapi-generator-ignore $SDK_DIR
}

cleanup() {
  LIST_FILE=".openapi-generator/FILES"
  
  if [[ ! -f "$LIST_FILE" ]]; then
      echo "Error: File '$LIST_FILE' not found."
      exit 1
  fi
  
  while IFS= read -r filepath || [[ -n "$filepath" ]]; do
      # Skip empty lines
      [[ -z "$filepath" ]] && continue
  
      # Check if the target file exists before removing
      if [[ -e "$filepath" ]]; then
          rm "$filepath"
          if [[ $? -eq 0 ]]; then
              echo "Successfully removed: $filepath"
          else
              echo "Failed to remove: $filepath"
          fi
      else
          echo "Skip: '$filepath' does not exist."
      fi
      
    done < "$LIST_FILE"
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

move_docs() {
  # Move generated docs to the root directory
  if [ -d "internal/openapi/docs" ]; then
      mkdir -p docs
      mv internal/openapi/docs/* docs/
      rm -rf internal/openapi/docs
      git add docs/
  fi
  
  # Move generated README to the root directory
  if [ -f "internal/openapi/README.md" ]; then
      mv internal/openapi/README.md README.md
  fi
}

prepare_directory
run_generator
move_docs
gofmt -w .