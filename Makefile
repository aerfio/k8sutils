CURRENT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
GOBIN = $(CURRENT_DIR)/bin

.PHONY: all
all: generate fmt test lint

.PHONY: generate
generate:
	go run ./internal/cmd/gvk-dirs-generator

.PHONY: test
test:
	go test ./... -race -v

bin/golangci-lint-install.sh:
	mkdir -p "bin"
	curl -fL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" -o "$@"
	chmod +x "$@"

GOLANGCI_LINT_VERSION = v2.4.0
GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint-${GOLANGCI_LINT_VERSION}
${GOLANGCI_LINT}: bin/golangci-lint-install.sh
	./bin/golangci-lint-install.sh -b $(abspath bin) ${GOLANGCI_LINT_VERSION}
	mv ./bin/golangci-lint ${GOLANGCI_LINT}

lint: ${GOLANGCI_LINT}
	$(GOLANGCI_LINT) run --config ${PWD}/.golangci.yml

fmt: ${GOLANGCI_LINT}
	$(GOLANGCI_LINT) fmt
