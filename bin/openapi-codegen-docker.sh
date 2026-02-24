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
      # Note: 'packageVersion' is not configured via config file because we need to inject its value at runtime
      # Please use the config file for all other configuration
}

# Replaces given acronyms, such as Id, Ip, Url, etc. to match recommended casing in GO
change_casing() {
  INITIALISMS="Id:ID Ip:IP Url:URL Sdk:SDK Vpn:VPN Mitm:MITM"
  # Replace given acronyms, such as Id, Ip, Url, etc. to match recommended casing in GO
  find "$SDK_DIR" \( -name "*.go" -o -name "*.md" \) -print | while read -r file; do
    for pair in $INITIALISMS; do
      from="${pair%%:*}"
      to="${pair##*:}"

      # Standalone identifiers (struct fields, params, vars, Id → ID)
      perl -i -pe "s/(\([^)]*\.md\))|(?<![A-Za-z])\\b$from\\b(?![A-Za-z])/\$1 ? \$1 : \"$to\"/ge" "$file"

      # Leading initialism (IpAddress → IPAddress)
      perl -i -pe "s/(\([^)]*\.md\))|(?<![A-Za-z])$from(?=[A-Z])/\$1 ? \$1 : \"$to\"/ge" "$file"

      # Compound identifiers (RuleId → RuleID)
      perl -i -pe "s/(\([^)]*\.md\))|([A-Za-z])$from(?=[A-Z_\\s\\W]|\$)/\$1 ? \$1 : \"\$2$to\"/ge" "$file"
    done

    # Special-case word normalization
    perl -i -pe 's/(\([^)]*\.md\))|Blocklist/$1 ? $1 : "BlockList"/ge' "$file"
  done
}

prepare_directory
run_generator

# Postprocessing

change_casing
gofmt -w .