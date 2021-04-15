PACKAGE_NAME := github.com/sentinel-official/hub
GOLANG_VERSION := 1.16.3
GOLANG_CROSS_VERSION := v$(GOLANG_VERSION)
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
BUILD_FLAGS := -tags="${BUILD_TAGS}"

benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}
.PHONY: benchmark

install: mod-vendor
	go install -mod=readonly ${BUILD_FLAGS} -ldflags "${LD_FLAGS}" ./cmd/sentinelhub
.PHONY: install

go-lint:
	@golangci-lint run --fix
.PHONY: go-lint

mod-vendor: tools
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

release-dry-run: mod-vendor
	docker run --rm --privileged \
	--env CGO_ENABLED=1 \
	--env BUILD_FLAGS=${BUILD_FLAGS} \
	--env LD_FLAGS="${LD_FLAGS}" \
	--volume `pwd`:/go/src/$(PACKAGE_NAME) \
	--workdir /go/src/$(PACKAGE_NAME) \
	troian/golang-cross:${GOLANG_CROSS_VERSION} \
	--rm-dist --snapshot
.PHONY: release-dry-run

release: mod-vendor
	@if [ -z "${GITHUB_TOKEN}" ]; then \
		echo "\033[91mGITHUB_TOKEN is required for release\033[0m";\
		exit 1;\
	fi
	docker run --rm --privileged \
	--env CGO_ENABLED=1 \
	--env BUILD_FLAGS=${BUILD_FLAGS} \
	--env LD_FLAGS="${LD_FLAGS}" \
	--env GITHUB_TOKEN=${GITHUB_TOKEN} \
	--volume `pwd`:/go/src/$(PACKAGE_NAME) \
	--workdir /go/src/$(PACKAGE_NAME) \
	troian/golang-cross:${GOLANG_CROSS_VERSION} \
	--rm-dist
.PHONY: release
