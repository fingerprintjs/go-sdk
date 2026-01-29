#!/usr/bin/env bash
pnpm exec changeset version && bash ./bin/openapi-codegen-docker.sh --no-gofmt
