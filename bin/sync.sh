#!/bin/bash
set -euo pipefail

defaultBaseUrl="https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi"
schemaUrl="${1:-$defaultBaseUrl/schemas/fingerprint-server-api-v4-normalized.yaml}"
examplesBaseUrl="${2:-$defaultBaseUrl/examples}"

mkdir -p ./res
curl -fSL -o ./res/fingerprint-server-api-v4.yaml "$schemaUrl"

examples=(
  'events/get_event_200.json'
  'events/search/get_event_search_200.json'
  'errors/400_ip_address_invalid.json'
  'errors/400_request_body_invalid.json'
  'errors/400_visitor_id_invalid.json'
  'errors/403_feature_not_enabled.json'
  'errors/403_secret_api_key_required.json'
  'errors/403_subscription_not_active.json'
  'errors/404_event_not_found.json'
  'errors/404_visitor_not_found.json'
  'errors/409_state_not_ready.json'
  'errors/429_too_many_requests.json'
)

baseDestination="./test/mocks"

for example in "${examples[@]}"; do
  destinationPath="$baseDestination/$example"
  destinationDir="$(dirname "$destinationPath")"
  mkdir -p "$destinationDir"

  exampleUrl="$examplesBaseUrl/$example"
  echo "Downloading $exampleUrl to $destinationPath"
  curl -fSL -o "$destinationPath" "$exampleUrl"
done
