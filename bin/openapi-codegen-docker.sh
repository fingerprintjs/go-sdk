#!/bin/bash

set -e

run_generator() {
  rm -rf sdk/
  mkdir -p sdk/
  cp .openapi-generator-ignore sdk/
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.19.0 generate \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -g go \
      -o /local/sdk \
      -t /local/template \
      -p packageName=fingerprint \
      --git-user-id=fingerprintjs \
      --git-repo-id=go-sdk \
      --global-property=skipGoMod=true \
      --global-property apiTests=false,modelTests=false \
      --additional-properties disallowAdditionalPropertiesIfNotPresent=false
}

run_generator

gofmt -w .
