CURRENT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
GOBIN = $(CURRENT_DIR)/bin

.PHONY: all
all: generate fmt test lint

.PHONY: generate
generate:
	go run ./internal/cmd/gvk-dirs-generator

GOFUMPT_VERSION ?= v0.6.0
GOFUMPT = ${GOBIN}/gofumpt-${GOFUMPT_VERSION}
${GOFUMPT}:
	GOBIN=${GOBIN} go install mvdan.cc/gofumpt@${GOFUMPT_VERSION}
	mv bin/gofumpt ${GOFUMPT}

GCI_VERSION = v0.13.3
GCI = bin/gci-${GCI_VERSION}
${GCI}:
	GOBIN=${GOBIN} go install github.com/daixiang0/gci@${GCI_VERSION}
	mv bin/gci ${GCI}

.PHONY: fmt
fmt: ${GOFUMPT} ${GCI}
	$(GOFUMPT) -w -extra .
	$(GCI) write . --skip-generated -s standard -s default -s localmodule -s blank -s dot --custom-order --skip-vendor

.PHONY: test
test:
	go test ./... -race -v

bin/golangci-lint-install.sh:
	mkdir -p "bin"
	curl -fL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" -o "$@"
	chmod +x "$@"

GOLANGCI_LINT_VERSION = v1.59.1
GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint-${GOLANGCI_LINT_VERSION}
${GOLANGCI_LINT}: bin/golangci-lint-install.sh
	./bin/golangci-lint-install.sh -b $(abspath bin) ${GOLANGCI_LINT_VERSION}
	mv ./bin/golangci-lint ${GOLANGCI_LINT}

lint: ${GOLANGCI_LINT}
	$(GOLANGCI_LINT) run --config ${PWD}/.golangci.yml
