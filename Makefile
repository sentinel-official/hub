PACKAGES := $(shell go list ./...)
VERSION := $(shell git rev-parse --short HEAD)
COMMIT := $(shell git log -1 --format='%H')
GOSUM := $(shell which gosum)

export GO111MODULE=on

BUILD_TAGS := netgo
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w \
	-X github.com/sentinel-official/sentinel-hub/version.Version=${VERSION} \
	-X github.com/sentinel-official/sentinel-hub/version.Commit=${COMMIT} \
	-X github.com/sentinel-official/sentinel-hub/version.BuildTags=${BUILD_TAGS}
ifneq (${GOSUM},)
	ifneq (${wildcard go.sum},)
		LD_FLAGS += -X github.com/sentinel-official/sentinel-hub/version.VendorHash=$(shell ${GOSUM} go.sum)
	endif
endif

BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test

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
