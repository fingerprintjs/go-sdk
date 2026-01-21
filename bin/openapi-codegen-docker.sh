#!/bin/bash

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/res/fingerprint-server-api-v4.yaml \
    -g go \
    -o /local \
    -t /local/template \
    -p packageName=sdk

gofmt -w .
