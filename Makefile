PACKAGES=$(shell go list ./... | grep -v '/vendor/')

all: get_tools get_vendor_deps build test

build:
	go build -o bin/sentinel-sdk-cli cmd/sentinel-sdk-cli/main.go && go build -o bin/sentinel-sdkd cmd/sentinel-sdkd/main.go

get_tools:
	go get github.com/golang/dep/cmd/dep

get_vendor_deps:
	@rm -rf vendor/
	@dep ensure -v

test:
	@go test $(PACKAGES)

benchmark:
	@go test -bench=. $(PACKAGES)

.PHONY: all build test benchmark