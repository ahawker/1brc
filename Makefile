#!/usr/bin/env make
.SUFFIXES:
.DEFAULT_GOAL := help

# Swap (or override) for real dataset.
#INPUT ?= datasets/measurements.txt
INPUT  ?= datasets/dev.txt

.PHONY: clean
clean: ## Clean up all previously built binaries.
	@rm -r bin

.PHONY: v1
v1: run-v1 ## "v1" - A real naiive shit house.

.PHONY: run-%
run-%:
	@go build -ldflags "-s -w" -o bin/$* $*/main.go
	@time bin/$* $(INPUT)

.PHONY: help
help:
	@grep -E '^[%a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
