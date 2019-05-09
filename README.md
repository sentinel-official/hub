# Sentinel SDK

[![](https://img.shields.io/github/release-pre/ironman0x7b2/sentinel-sdk.svg?style=flat)](https://github.com/ironman0x7b2/sentinel-sdk/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ironman0x7b2/sentinel-sdk)](https://goreportcard.com/report/github.com/ironman0x7b2/sentinel-sdk)
[![](https://tokei.rs/b1/github/ironman0x7b2/sentinel-sdk)](https://github.com/ironman0x7b2/sentinel-sdk)

```
rm -rf $HOME/.hubd $HOME/.hubcli && \
hubd init --chain-id hub testing && \
hubcli keys add genesis && \
hubd add-genesis-account $(hubcli keys show genesis -a) 2000000000000000usent && \
hubd gentx --name genesis && \
hubd collect-gentxs
```
