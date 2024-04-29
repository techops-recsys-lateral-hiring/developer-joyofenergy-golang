.PHONY: setup test all clean lint build run help

BUILD_DIR := bin
TOOLS_DIR := tools

.DEFAULT_GOAL:=help

setup: ## Setup server.
	@go mod download

test: ## Run test.
	@go test -v ./... -covermode=atomic

all: clean lint test build run ## Run all tests, then build and run

.PHONY: $(BUILD_DIR)/server
$(BUILD_DIR)/server:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/server ./cmd/server

build: $(BUILD_DIR)/server ## Build the binary

clean: ## Clean up, i.e. remove build artifacts
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod tidy

run: build ## Run the binary
	$(BUILD_DIR)/server

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

lint: $(TOOLS_DIR)/golangci-lint/golangci-lint ## Run linters
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
