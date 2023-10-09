CURRENT_DIR = $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
GOBIN = $(CURRENT_DIR)/bin

.PHONY: all
all: run fmt

.PHONY: run
run:
	go run .

GOFUMPT_VERSION ?= v0.5.0
GOFUMPT = ${GOBIN}/gofumpt-${GOFUMPT_VERSION}
${GOFUMPT}:
	GOBIN=${GOBIN} go install mvdan.cc/gofumpt@${GOFUMPT_VERSION}
	mv bin/gofumpt ${GOFUMPT}

.PHONY: fmt
fmt: ${GOFUMPT}
	${GOFUMPT} -w -extra .

