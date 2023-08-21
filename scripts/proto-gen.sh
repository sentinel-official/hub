#!/usr/bin/env bash

set -euo pipefail

cd ./proto/ && \
buf generate && \
cd .. && \
cp -r github.com/sentinel-official/hub/* ./ && \
rm -rf github.com/