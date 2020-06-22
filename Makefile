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

e2e-tests: ## Integration tests
	@ginkgo -cover -failFast -progress --reportPassed tests/e2e

unit-tests: ## Run tests
	@go test -race -v -cover interview-accountapi/cmd/...

build: ## Tidy Up
	@go mod tidy

lint: ## Run linters
	@docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:$(GOLANG_CI_LINT) golangci-lint run -v

validate: ## Validate files with pre-commit hooks
	@pre-commit run --all-files

.PHONE: docs help hooks validate run build
