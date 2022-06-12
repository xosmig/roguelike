[![Go Report Card](https://goreportcard.com/badge/github.com/xosmig/roguelike)](https://goreportcard.com/report/github.com/xosmig/roguelike)
[![Build and test](https://github.com/xosmig/roguelike/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/xosmig/roguelike/actions/workflows/build-and-test.yml)
[![codecov](https://codecov.io/gh/xosmig/roguelike/branch/master/graph/badge.svg)](https://codecov.io/gh/xosmig/roguelike)

# Roguelike

Simple roguelike game skeleton, written in Go using Termbox-Go library.

## Building and running

* install go: https://golang.org

* checkout the repository: `git clone git@github.com:xosmig/roguelike.git`

* `cd roguelike`

* `go build`

* `./roguelike`

## Working with sources

* You can run unit tests using `go test ./...` or simply `./test.sh`

* You can run `./generate.sh` to regenerate the generated sources (such as mock objects)

## Structure diagram
`diagrams/structure.png`

![](diagrams/structure.png)
