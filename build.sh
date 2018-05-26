#!/bin/sh -eux

dep ensure
go build -o=roguelike.out "$@" ./cmd/main.go
