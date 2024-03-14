FROM golang:1.22-alpine3.18 as build

COPY . /root/

RUN apk add build-base ca-certificates git linux-headers wget && \
    cd /root/ && \
    ARCH=$(uname -m) && \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $NF}') && \
    wget -q -O /usr/local/lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a && \
    STATIC=true make --jobs=$(nproc) build

FROM alpine:3.19

COPY --from=build /root/build/sentinelhub /usr/local/bin/sentinelhub

ENTRYPOINT ["sentinelhub"]
