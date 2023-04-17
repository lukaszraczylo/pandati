.DEFAULT_GOAL := help

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m \t%s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: all
all: lint test ## Run all checks on codebase

.PHONY: test
test: ## Run tests on codebase
	go test -race -cover ./...

.PHONY: lint
lint: ## Run linters against codebase
	go fmt ./...
	golangci-lint run ./...

.PHONY: update
update: ## Update dependencies
	go get -u ./...
	go mod tidy -compat=1.20