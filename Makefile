.DEFAULT_GOAL := default
VERSION := $(shell git describe --tags | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TENDERMINT_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's/.* //')

comma := ,
whitespace := $() $()

build_tags := $(strip netgo ledger)
ld_flags := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
    -X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TENDERMINT_VERSION)

ifeq ($(STATIC),true)
	build_tags += muslc
	ld_flags += -linkmode=external -extldflags '-Wl,-z,muldefs -static'
endif

BUILD_TAGS = $(subst $(whitespace),$(comma),$(build_tags))
LD_FLAGS = ${ld_flags} -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS}

.PHONY: benchmark
benchmark:
	@go test -mod=readonly -v -bench ./...

.PHONY: build
build:
	go build -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" -trimpath \
		-o ./build/sentinelhub ./cmd/sentinelhub

.PHONY: clean
clean:
	rm -rf ./build ./vendor ./coverage.txt

.PHONE: default
default: build

.PHONY: install
install:
	go install -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" -trimpath ./cmd/sentinelhub

.PHONY: go-lint
go-lint:
	@golangci-lint run --fix

.PHONY: mod-vendor
mod-vendor:
	@go mod vendor
	@modvendor -copy="**/*.proto" -include=github.com/cosmos/cosmos-sdk/proto,github.com/cosmos/cosmos-sdk/third_party/proto

.PHONY: proto-gen
proto-gen:
	@scripts/proto-gen.sh

.PHONY: proto-lint
proto-lint:
	@find proto -name *.proto -exec clang-format -i {} \;

.PHONY: test
test:
	@go test -mod=readonly -timeout 15m -v ./...

.PHONT: test-coverage
test-coverage:
	@go test -mod=readonly -timeout 15m -v -covermode=atomic -coverprofile=coverage.txt ./...

.PHONY: tools
tools:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@go install github.com/goware/modvendor@v0.5.0
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
