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
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.19.0 generate \
      -c /local/openapi-config.yml \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -t /local/template \
      --additional-properties=packageVersion=$PACKAGE_VERSION
      # Note: 'packageVersion' is not configured via config file because we need to inject its value at runtime
      # Please use the config file for all other configuration
}

prepare_directory
run_generator

if [[ $1 != "--no-gofmt" ]]; then
  gofmt -w .
fi
