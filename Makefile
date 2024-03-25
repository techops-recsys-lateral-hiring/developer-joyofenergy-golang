BUILD_DIR := bin
TOOLS_DIR := tools

default: all

all: clean lint test build run

.PHONY: $(BUILD_DIR)/server
bin/server: cmd/server/*.go
	CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./bin/server ./cmd/server/

.PHONY: build
build: bin/server

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	@go mod vendor
	@go mod tidy

.PHONY: run
run: build
	bin/server

.PHONY: lint
lint:
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test
test:
	go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
