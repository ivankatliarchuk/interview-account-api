SHELL = /bin/bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
dc=docker-compose

CURRENT_DIR ?= $(shell realpath ..)
GOLANG_CI_LINT ?= v1.27.0

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docs:
	@docker run -v $$PWD/:/docs pandoc/latex -f markdown /docs/README.md -o /docs/build/output/README.pdf

services: ## Run services
	@docker-compose up

run-int-tests: ## Integration tests
	@ginkgo -cover -failFast -progress --reportPassed integration

run-tests: ## Run tests
	@go test -race -v -cover interview-accountapi/cmd/...

build: ## Build & Tidy
	@go mod tidy
	@go build

lint: ## Run linters
	@docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:$(GOLANG_CI_LINT) golangci-lint run -v

.PHONE: docs help hooks validate run build
