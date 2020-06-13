SHELL = /bin/bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docs:
	@docker run -v $$PWD/:/docs pandoc/latex -f markdown /docs/README.md -o /docs/build/output/README.pdf

hooks: ## Commit hooks setup
	@pre-commit install
	@pre-commit gc
	@pre-commit autoupdate

validate: ## Validate with pre-commit hooks
	@pre-commit run --all-files

services: ## Run services
	@docker-compose up

run: ## Run it
	@go run main.go

build: ## Build & Tidy
	@go mod tidy
	@go build

.PHONE: docs help hooks validate run build
