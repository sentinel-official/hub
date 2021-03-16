# Sentinel Hub

[![Tag](https://img.shields.io/github/tag/sentinel-official/hub.svg)](https://github.com/sentinel-official/hub/releases/latest)
[![GoReportCard](https://goreportcard.com/badge/github.com/sentinel-official/hub)](https://goreportcard.com/report/github.com/sentinel-official/hub)
[![Licence](https://img.shields.io/github/license/sentinel-official/hub.svg)](https://github.com/sentinel-official/hub/blob/development/LICENSE)
[![LoC](https://tokei.rs/b1/github/sentinel-official/hub)](https://github.com/sentinel-official/hub)

## Installation
Two steps:

1) Setup a Go enviornment
2) Compile and install Sentinel

## Linux

### Install Go: Ubuntu and Debian 
Ubuntu and debian do not have the latest version of Go in their repositories.

To get the latest version of Go, and configure your GOPATH, do like:

```bash
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install -y git golang-go build-essential
mkdir ~/go
export GOPATH=~/go
export PATH=$PATH:/go/bin
export GOBIN=~/go/bin
```

Users of other distros, like Arch, will find the latest Go in their package manager.  

### Install GO: MacOS

[Go 1.16+](https://golang.org/dl/)


### Install 

```bash
git clone https://github.com/sentinel-official/hub/
cd hub
make install
sentinelhubd init chooseanicehandle --chain-id sentinel-turing-3a
```

### Run
```
sentinelhubd start --p2p.seeds 091715cf98995180a6da44bd28d3c11f8636a962@51.158.189.149:26656,790026684d76c66347941e1c21a904b141014568@3.8.10.143:26656,b34f0b79731365b1cb89b9791dc0e1392ced77c9@206.189.253.224:26656,835b12099f5869ac9160376d60ab58060169a9c6@128.199.31.151:26656 --log_level "main:info,state:info,x/node:info,x/subscription:info,x/session:info,*:error" --invar-check-period 1
```


**Note:** To install a specific version or commit use `git checkout` command

## Additional Documentation

For the additional documentation on the Sentinel Hub, please visit - https://docs.sentinel.co
