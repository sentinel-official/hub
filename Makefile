PACKAGES=$(shell go list ./... | grep -v '/vendor/')

all: get_tools get_vendor_deps build test

build:
	go build -o bin/sentinel-hub-cli cmd/sentinel-hub-cli/main.go && go build -o bin/sentinel-hubd cmd/sentinel-hubd/main.go
	go build -o bin/sentinel-vpn-cli cmd/sentinel-vpn-cli/main.go && go build -o bin/sentinel-vpnd cmd/sentinel-vpnd/main.go

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