include contrib/tools/Makefile

PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

BUILD_TAGS := netgo,ledger
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel-hub \
    -X github.com/cosmos/cosmos-sdk/version.ServerName=sentinel-hubd \
    -X github.com/cosmos/cosmos-sdk/version.ClientName=sentinel-hub-cli \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS}
BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test benchmark

build: mod_verify
ifeq (${OS},Windows_NT)
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd.exe ./cmd/sentinel-hubd
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-cli.exe ./cmd/sentinel-hub-cli
else
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd ./cmd/sentinel-hubd
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-cli ./cmd/sentinel-hub-cli
endif

install: mod_verify
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hubd
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hub-cli

test:
	@go test -mod=readonly -v -cover ${PACKAGES}

benchmark:
	@go test -mod=readonly -v -bench ${PACKAGES}

simulate_short:
	@go test -mod=readonly -v -timeout=1h -run TestFullAppSimulation \
		-Enabled=true -Seed=4 -NumBlocks=50 -BlockSize=50 -Commit=true

simulate_long:
	@go test -mod=readonly -v -timeout=1h -run TestFullAppSimulation \
		-Enabled=true -Seed=4 -NumBlocks=2500 -BlockSize=50 -Commit=true

simulate_multi_long:
	@runsim -Jobs=4 -SimAppPkg=. 500 1 TestFullAppSimulation

mod_verify:
	@echo "Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all build install test benchmark \
simulate_short simulate_multi mod_verify
