PACKAGES = $(shell go list ./... | grep -v '/vendor/')
VERSION = $(shell git rev-parse --short HEAD)
COMMIT := $(shell git log -1 --format='%H')
CAT := $(if $(filter $(OS),Windows_NT),type,cat)
BUILD_TAGS = netgo
BUILD_FLAGS = -tags "${BUILD_TAGS}" -ldflags \
	"-X github.com/ironman0x7b2/sentinel-sdk/version.Version=${VERSION} \
	-X github.com/ironman0x7b2/sentinel-sdk/version.Commit=${COMMIT} \
	-X github.com/ironman0x7b2/sentinel-sdk/version.VendorDirHash=$(shell $(CAT) .vendor_version) \
	-s -w"

all: get_tools get_vendor_deps install test

build:
	go build $(BUILD_FLAGS) -o bin/hub-cli cmd/hub-cli/main.go
	go build $(BUILD_FLAGS) -o bin/hubd cmd/hubd/main.go
	go build $(BUILD_FLAGS) -o bin/vpn-cli cmd/vpn-cli/main.go
	go build $(BUILD_FLAGS) -o bin/vpnd cmd/vpnd/main.go

install:
	go install $(BUILD_FLAGS) ./cmd/hub-cli
	go install $(BUILD_FLAGS) ./cmd/hubd
	go install $(BUILD_FLAGS) ./cmd/vpn-cli
	go install $(BUILD_FLAGS) ./cmd/vpnd

get_tools:
	go get github.com/golang/dep/cmd/dep

get_vendor_deps:
	@rm -rf vendor/ .vendor-new/
	@dep ensure -v
	tar -c vendor/ | sha1sum | cut -d' ' -f1 > ".vendor_version"

test:
	@go test -cover $(PACKAGES)

benchmark:
	@go test -bench=. $(PACKAGES)

.PHONY: all build test benchmark
