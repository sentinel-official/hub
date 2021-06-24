PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TENDERMINT_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

BUILD_TAGS := $(strip netgo,ledger)
LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
    -X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TENDERMINT_VERSION)

.PHONY: benchmark
benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" \
		-o ./build/sentinelhub-${VERSION}-darwin-amd64 ./cmd/sentinelhub
	GOOS=linux GOARCH=amd64 go build -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" \
		-o ./build/sentinelhub-${VERSION}-linux-amd64 ./cmd/sentinelhub
	GOOS=windows GOARCH=amd64 go build -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" \
		-o ./build/sentinelhub-${VERSION}-windows-amd64.exe ./cmd/sentinelhub

.PHONY: clean
clean:
	rm -rf ./bin ./vendor

.PHONY: install
install: mod-vendor
	go install -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" ./cmd/sentinelhub

.PHONY: go-lint
go-lint:
	@golangci-lint run --fix

.PHONY: mod-vendor
mod-vendor: tools
	@go mod vendor
	@modvendor -copy="**/*.proto" -include=github.com/cosmos/cosmos-sdk/proto,github.com/cosmos/cosmos-sdk/third_party/proto

.PHONY: proto-gen
proto-gen:
	@scripts/proto-gen.sh

.PHONY: proto-lint
proto-lint:
	@find proto -name *.proto -exec clang-format-12 -i {} \;

.PHONY: test
test:
	@go test -mod=readonly -v -cover ${PACKAGES}

.PHONY: tools
tools:
	@go install github.com/bufbuild/buf/cmd/buf@v0.37.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	@go install github.com/goware/modvendor@v0.3.0
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
