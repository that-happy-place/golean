PROJECT_ROOT ?= .
GOPATH=$(shell go env GOPATH)
GO_BASE_CMD = GO111MODULE=on GOPROXY="https://proxy.golang.org,direct" go

.PHONY: all
all: polish test

.PHONY: init
init:
	@echo " > Installing required dependencies"
	[[ -e `which brew` ]] || /bin/bash -c "curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh"
	[[ -e `which pre-commit` ]] || brew install pre-commit
	$(GO_BASE_CMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	$(GO_BASE_CMD) install golang.org/x/tools/cmd/goimports@v0.24.0
	$(GO_BASE_CMD) install github.com/go-critic/go-critic/cmd/gocritic@latest
	pre-commit install

.PHONY: polish
polish: tidy fmt vet lint

.PHONY: tidy
tidy:
	@echo " > Running go mod tidy"
	$(GO_BASE_CMD) mod tidy

.PHONY: fmt
fmt:
	@echo " > Formatting code"
	$(GO_BASE_CMD) fmt $(PROJECT_ROOT)/...

.PHONY: lint
lint:
	@echo " > Running linters"
	$(GOPATH)/bin/golangci-lint run --allow-parallel-runners $(PROJECT_ROOT)/...

.PHONY: test
test:
	rm -rf $(PROJECT_ROOT)/coverage.out
	@echo " > Running tests with coverage"
	$(GO_BASE_CMD) test -race -coverprofile=coverage.out -timeout 30m -coverpkg=./... $(PROJECT_ROOT)/...
	$(GO_BASE_CMD) tool cover -func=coverage.out

.PHONY: vet
vet:
	@echo " > Running go vet"
	$(GO_BASE_CMD) vet $(PROJECT_ROOT)/...
