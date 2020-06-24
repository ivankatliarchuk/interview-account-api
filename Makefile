SHELL = /bin/bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
dc=docker-compose

IMAGE := interview-acccountapi/cmd
# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
TAG := $(VERSION)
SRC_DIRS := cmd pkg # directories which hold app source (not vendored)
COMMAND ?= /bin/bash

BUILD_IMAGE ?= golang:1.12-alpine

CURRENT_DIR ?= $(shell realpath ..)
GOLANG_CI_LINT ?= v1.27.0

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docs:
	@docker run -v $$PWD/:/docs pandoc/latex -f markdown /docs/README.md -o /docs/build/output/README.pdf

services-up: ## Run services
	@docker-compose up -d

services-down: ## Run services
	@docker-compose down

e2e-tests: ## Integration tests
	@echo "go e2e SDK packages"
	@ginkgo -cover -failFast -progress --reportPassed tests/e2e

unit-tests: ## Run tests
	@echo "go unit-test SDK packages"
	@go test -race -v -cover interview-accountapi/cmd/...

lint: ## Run linters
	@echo "go lint SDK packages"
	@docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:$(GOLANG_CI_LINT) golangci-lint run -v

validate: ## Validate files with pre-commit hooks
	@pre-commit run --all-files

docker-build: ## Build docker container
	@docker build -t $(IMAGE):$(TAG) -t $(IMAGE):latest -f Dockerfile .
	@docker images -q $(IMAGE):$(TAG) > $@

docker-run: ## Run commands in docker container e.g. make docker-run COMMAND='make go-tidy'
docker-run: docker-build
	docker run -it --rm -v $$(pwd):/app -w /app $(IMAGE):$(TAG) $(COMMAND)

.PHONE: docs help hooks validate run lint docker-build docker-run
