PACKAGES = $(shell go list ./... | grep -v '/vendor/')
VERSION = $(shell git rev-parse --short HEAD)
BUILD_FLAGS = -ldflags "-X github.com/ironman0x7b2/sentinel-sdk/vendor/github.com/cosmos/cosmos-sdk/version.Version=${VERSION} -s -w"

all: get_tools get_vendor_deps build test

build:
	go build $(BUILD_FLAGS) -o bin/sentinel-hub-cli cmd/sentinel-hub-cli/main.go
	go build $(BUILD_FLAGS) -o bin/sentinel-hubd cmd/sentinel-hubd/main.go
	go build $(BUILD_FLAGS) -o bin/sentinel-vpn-cli cmd/sentinel-vpn-cli/main.go
	go build $(BUILD_FLAGS) -o bin/sentinel-vpnd cmd/sentinel-vpnd/main.go

get_tools:
	go get github.com/golang/dep/cmd/dep

get_vendor_deps:
	@rm -rf vendor/
	@dep ensure -v

test:
	@go test -cover $(PACKAGES)

benchmark:
	@go test -bench=. $(PACKAGES)

.PHONY: all build test benchmark
