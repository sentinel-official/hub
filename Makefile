PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TENDERMINT_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

build_tags = netgo ledger
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel \
    -X github.com/cosmos/cosmos-sdk/version.AppName=sentinelhub \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TENDERMINT_VERSION)

ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags)) 

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


.PHONY: benchmark
benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}

.PHONY: clean
clean:
	rm -rf ./vendor

BUILD_TARGETS := install build

.PHONY: install build
$(BUILD_TARGETS):
	go $@ -mod=readonly $(BUILD_FLAGS) -trimpath ./cmd/sentinelhub

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
	@go test -mod=readonly -timeout 15m -v ${PACKAGES}

.PHONT: test-coverage
test-coverage:
	@go test -mod=readonly -timeout 15m -v -covermode=atomic -coverprofile=coverage.txt ${PACKAGES}

.PHONY: tools
tools:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@go install github.com/goware/modvendor@v0.5.0
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
