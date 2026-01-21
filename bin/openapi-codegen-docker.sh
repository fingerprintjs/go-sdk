#!/bin/bash

mv README.md temp.README.md

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/res/fingerprint-server-api-v4.yaml \
    -g go \
    -o /local \
    -t /local/template \
    -p packageName=sdk

mv README.md README.generated.md
mv temp.README.md README.md

gofmt -w .
