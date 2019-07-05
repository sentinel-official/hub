PACKAGES := $(shell go list ./...)
VERSION := $(shell git rev-parse --short HEAD)
COMMIT := $(shell git log -1 --format='%H')
GOSUM := $(shell which gosum)

export GO111MODULE=on

BUILD_TAGS := netgo
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w \
	-X github.com/sentinel-official/hub/version.Version=${VERSION} \
	-X github.com/sentinel-official/hub/version.Commit=${COMMIT} \
	-X github.com/sentinel-official/hub/version.BuildTags=${BUILD_TAGS}
ifneq (${GOSUM},)
	ifneq (${wildcard go.sum},)
		LD_FLAGS += -X github.com/sentinel-official/hub/version.VendorHash=$(shell ${GOSUM} go.sum)
	endif
endif

BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test

test_sim_hub_fast:
	@echo "Running hub simulation. This may take several minutes..."
	@go test -mod=readonly ./app -run TestHubSimulation  -num_blocks=100 -block_size=100 -commit=true -seed=99  -v -timeout 24h

test_sim_hub_import_export:
	@echo "Running hub simulation. This may take several minutes..."
	@go test -mod=readonly ./app -run TestHubSimulationImportExport  -num_blocks=100 -block_size=100 -commit=true -seed=99  -v -timeout 24h

test_sim_hub_after_import:
	@echo "Running hub simulation. This may take several minutes..."
	@go test -mod=readonly ./app -run TestHubSimulationAfterImport  -num_blocks=100 -block_size=100 -commit=true -seed=99  -v -timeout 24h

build: dep_verify
ifeq (${OS},Windows_NT)
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd.exe cmd/sentinel-hubd/main.go
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubcli.exe cmd/sentinel-hubcli/main.go
else
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd cmd/sentinel-hubd/main.go
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubcli cmd/sentinel-hubcli/main.go
endif

install: dep_verify
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hubd
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hubcli

test:
	@go test -mod=readonly -cover ${PACKAGES}

benchmark:
	@go test -mod=readonly -bench=. ${PACKAGES}

dep_verify:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all build install test benchmark, dep_verify
