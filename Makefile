include contrib/tools/Makefile

PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
    -X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)
BUILD_TAGS := $(strip netgo,ledger)
BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test benchmark

install: mod_vendor
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinelhub

test:
	@go test -mod=readonly -v -cover ${PACKAGES}

benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}

mod_vendor:
	@go mod vendor
	@modvendor -copy="**/*.proto" -include=github.com/cosmos/cosmos-sdk/proto,github.com/cosmos/cosmos-sdk/third_party/proto

.PHONY: all install test benchmark mod_vendor
