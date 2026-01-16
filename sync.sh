#!/bin/bash

# Script to download the latest OpenAPI template.
# Run without arguments to download APIv4
# ./sync.sh
# Run with argument v3 to download deprecated APIv3
# ./sync.sh v3

downloadV3(){
  curl -o ./res/fingerprint-server-api.yaml https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/schemas/fingerprint-server-api-compact.yaml
}

downloadV4(){
  curl -o ./res/fingerprint-server-api-v4.yaml https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/schemas/fingerprint-server-api-v4.yaml
}

if [[ "$1" = "v3" ]]; then
  downloadV3
else
  downloadV4
fi
