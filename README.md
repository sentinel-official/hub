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

### Install GO: MacOS

[Go 1.16+](https://golang.org/dl/)


### Install 

```bash
git clone https://github.com/sentinel-official/hub/
cd hub
make install
sentinelhubd init chooseanicehandle
```

**Note:** To install a specific version or commit use `git checkout` command

## Additional Documentation

For the additional documentation on the Sentinel Hub, please visit - https://docs.sentinel.co
