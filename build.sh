#!/bin/sh -eux

dep ensure
go build "$@"
