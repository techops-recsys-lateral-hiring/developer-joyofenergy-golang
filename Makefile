.PHONY: setup test build dev all clean lint run help
.PHONY: docker-build docker-run docker-kill

BUILD_DIR := bin
TOOLS_DIR := tools

.DEFAULT_GOAL:=help

setup: ## Set up the server.
	@go mod download

test: ## Run the tests.
	@go test -v ./... -cover -coverprofile=coverage.txt -covermode=atomic

build: ## Build the server.
	@go build -v -ldflags "-s -w" -o ./bin/server ./cmd/server

dev: ## Run local server.
	@go run ./cmd/server/main.go

all: clean lint test build run ## Run all tests, then build and run

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

docker-build: ## Build container's Docker.
	@docker build -t server .

docker-run: ## Run container's Docker.
	@docker run --name new-server -p 8080:8080 -it server

docker-kill: ## Kill container's Docker.
	@docker kill new-server
