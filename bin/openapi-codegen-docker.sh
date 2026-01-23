#!/bin/bash

set -e

run_generator() {
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
      -i /local/res/fingerprint-server-api-v4.yaml \
      -g go \
      -o /local \
      -t /local/template \
      -p packageName=fingerprint,gitUserId=fingerprintjs,gitRepoId=fingerprint-server-go-sdk
}

mv README.md temp.README.md # hide manual README
trap 'mv -f temp.README.md README.md' EXIT # restore manual README
run_generator
mv README.md README.generated.md # move generated README

gofmt -w .
