#!/bin/bash

set -e -o pipefail

export GOCACHE=${HOME}/.go-build
export GOMODCACHE=${HOME}/go/pkg/mod

echo "GOCACHE: $GOCACHE"
echo "GOMODCACHE: $GOMODCACHE"

echo "--- Restoring cache"
buildkite-agent cache restore --ids golang_build --ids golang_pkg --bucket-url $CACHE_BUCKET_URL

echo "--- Running tests"
go test -coverprofile coverage.out -coverpkg=./... ./...
go run github.com/nikolaydubina/go-cover-treemap -coverprofile coverage.out > cover-tree.svg
echo '<details><summary>Coverage tree map</summary><img src="artifact://cover-tree.svg" alt="Test coverage tree map" width="70%"></details>' | buildkite-agent annotate --style "info"

echo "--- Saving cache"
buildkite-agent cache save --ids golang_build --ids golang_pkg --bucket-url $CACHE_BUCKET_URL