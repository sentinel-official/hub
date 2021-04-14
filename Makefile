PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TM_CORE_SEMVER := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

BUILD_TAGS := $(strip netgo,ledger)
LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
    -X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_CORE_SEMVER)
BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: tools mod-vendor install
.PHONY: all

benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}
.PHONY: benchmark

install:
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinelhub
.PHONY: install

go-lint:
	@golangci-lint run --fix
.PHONY: go-lint

mod-vendor:
	@go mod vendor
	@modvendor -copy="**/*.proto" -include=github.com/cosmos/cosmos-sdk/proto,github.com/cosmos/cosmos-sdk/third_party/proto
.PHONY: mod-vendor

proto-gen:
	@scripts/proto-gen.sh
.PHONY: proto-gen

proto-lint:
	@find proto -name *.proto -exec clang-format-12 -i {} \;
.PHONY: proto-lint

test:
	@go test -mod=readonly -v -cover ${PACKAGES}
.PHONY: test

tools:
	@go install github.com/bufbuild/buf/cmd/buf@v0.37.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	@go install github.com/goware/modvendor@v0.3.0
.PHONY: tools
