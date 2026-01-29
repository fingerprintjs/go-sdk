#!/usr/bin/env bash
bash ./bin/openapi-codegen-docker.sh && pnpm exec changeset publish
