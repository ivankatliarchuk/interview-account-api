#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

# Enable C code, as it is needed for SQLite3 database binary
# Enable go modules
export CGO_ENABLED=1
export GO111MODULE=on
# export GOFLAGS="-mod=vendor"

export SRC_DIRS=cmd pkg
# Collect test targets
TARGETS=$(for d in "$SRC_DIRS"; do echo ./$d/...; done)

echo $TARGETS

go test $TARGETS
