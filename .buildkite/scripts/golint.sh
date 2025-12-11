#!/bin/bash

set -e -o pipefail

export GOCACHE=${HOME}/.go-build
export GOMODCACHE=${HOME}/go/pkg/mod

echo "--- Restoring cache"
buildkite-agent cache restore --ids golang-build,golang-pkg --bucket-url $CACHE_BUCKET_URL

echo "--- Running golangci-lint"
golangci-lint run --verbose --timeout 3m

echo "--- Saving cache"
buildkite-agent cache save --ids golang-build,golang-pkg --bucket-url $CACHE_BUCKET_URL
