.DEFAULT_GOAL := help

.PHONY: help fmt vet test race cover build run ci

help: ## Show available targets
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9_.-]+:.*## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS=":.*## "}; {printf "  %-8s %s\n", $$1, $$2}'

fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

test: ## Run tests
	go test ./...

race: ## Run tests with race detector
	go test ./... -race

cover: ## Run tests with coverage profile
	go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out | awk '/^total:/ {print "Coverage:", $$3}'

build: ## Build CLI binary to ./bin/deck
	mkdir -p bin
	go build -o ./bin/deck ./cmd/deck

run: ## Run CLI help
	go run ./cmd/deck --help

ci: ## Run local CI checks (fmt check, vet, test, race)
	@unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not gofmt-formatted:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi
	go vet ./...
	go test ./...
	go test ./... -race
