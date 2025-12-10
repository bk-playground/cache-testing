#!/bin/bash

set -e -o pipefail

export GOCACHE=$(pwd)/.gocache
export GOMODCACHE=$(pwd)/.gomodcache

echo "--- Restoring cache"
buildkite-agent cache restore --ids golang --bucket-url $CACHE_BUCKET_URL

echo "--- Running tests"
go test -coverprofile coverage.out -coverpkg=./... ./...
go run github.com/nikolaydubina/go-cover-treemap -coverprofile coverage.out > cover-tree.svg
echo '<details><summary>Coverage tree map</summary><img src="artifact://cover-tree.svg" alt="Test coverage tree map" width="70%"></details>' | buildkite-agent annotate --style "info"

echo "--- Saving cache"
buildkite-agent cache save --ids golang --bucket-url $CACHE_BUCKET_URL