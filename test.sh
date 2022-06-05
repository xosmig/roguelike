#!/usr/bin/bash
set -eu

go test "$@" ./...
