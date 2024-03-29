CURRENT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
GOBIN = $(CURRENT_DIR)/bin

.PHONY: all
all: generate fmt test

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
