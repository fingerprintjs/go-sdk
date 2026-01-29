#!/usr/bin/env bash
bash ./bin/openapi-codegen-docker.sh --no-gofmt && pnpm exec changeset publish
