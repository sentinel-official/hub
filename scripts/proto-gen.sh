#!/usr/bin/env bash

set -eo pipefail

directories=$(find proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for directory in ${directories}; do
  files=$(find "${directory}" -maxdepth 1 -name '*.proto')
  for file in ${files}; do
    protoc \
      --proto_path="proto" \
      --proto_path="vendor/github.com/cosmos/cosmos-sdk/proto" \
      --proto_path="vendor/github.com/cosmos/cosmos-sdk/third_party/proto" \
      --gocosmos_out="plugins=interfacetype+grpc,Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:${GOPATH}/src" \
      "${file}"

    protoc \
      --proto_path="proto" \
      --proto_path="vendor/github.com/cosmos/cosmos-sdk/proto" \
      --proto_path="vendor/github.com/cosmos/cosmos-sdk/third_party/proto" \
      --grpc-gateway_out="logtostderr=true:${GOPATH}/src" \
      "${file}"
  done
done
