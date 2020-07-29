GO_BIN := $(shell go env GOPATH)/bin
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIR := $(shell cd $(shell dirname $(MAKEFILE_PATH)) && pwd)

install_golangci-lint:
	@sh ${MAKEFILE_DIR}/install_golangci-lint.sh -b ${GO_BIN} v1.27.0

install_runsim:
	@go get github.com/cosmos/tools/cmd/runsim

.PHONY: install_golangci install_runsim