#!/bin/bash

set -e

run_generator() {
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -g go \
      -o /local \
      -t /local/template \
      -p packageName=fingerprint \
      --git-user-id=fingerprintjs \
      --git-repo-id=fingerprint-server-go-sdk
}

run_generator

gofmt -w .
