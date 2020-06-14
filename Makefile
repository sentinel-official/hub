PACKAGES := $(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SIMAPP = ./simapp

BUILD_TAGS := netgo,ledger
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=sentinel-hub \
    -X github.com/cosmos/cosmos-sdk/version.ServerName=sentinel-hub-daemon \
    -X github.com/cosmos/cosmos-sdk/version.ClientName=sentinel-hub-cli \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS}
BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test

build: dep_verify
ifeq (${OS},Windows_NT)
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-daemon.exe ./cmd/sentinel-hub-daemon
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-cli.exe ./cmd/sentinel-hub-cli
else
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-daemon ./cmd/sentinel-hub-daemon
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hub-cli ./cmd/sentinel-hub-cli
endif

install: dep_verify
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hub-daemon
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hub-cli

test:
	@go test -mod=readonly -cover ${PACKAGES}

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 100
SIM_COMMIT ?= true

test_sim_hub_fast:
	@echo "Running hub simulation for numBlocks=100, blockSize=200. This may take awhile!"
	@go test -mod=readonly $(SIMAPP) -run TestFullAppSimulation -Genesis=${HOME}/.sentinel-hub-daemon/config/genesis.json \
	    -Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99  -v -Period=5


test_sim_benchmark:
	@echo "Running hub benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$  $(SIMAPP) -bench=BenchmarkFullAppSimulation  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -Seed 15

benchmark:
	@go test -mod=readonly -bench=. ${PACKAGES}

dep_verify:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all build install test benchmark dep_verify test_sim_hub_fast test_sim_benchmark
