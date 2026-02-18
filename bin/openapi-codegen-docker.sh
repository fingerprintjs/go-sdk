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
}

prepare_directory
run_generator
gofmt -w .
