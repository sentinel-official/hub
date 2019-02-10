# Sentinel SDK

[![](https://img.shields.io/github/release-pre/ironman0x7b2/sentinel-sdk.svg?style=flat)](https://github.com/ironman0x7b2/sentinel-sdk/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/ironman0x7b2/sentinel-sdk)](https://goreportcard.com/report/github.com/ironman0x7b2/sentinel-sdk)
[![](https://tokei.rs/b1/github/ironman0x7b2/sentinel-sdk)](https://github.com/ironman0x7b2/sentinel-sdk)

```
rm -rf $HOME/.vpnd $HOME/.vpncli && \
vpnd init --chain-id vpn testing && \
vpncli keys add genesis && \
vpnd add-genesis-account $(vpncli keys show genesis -a) 1000000000sent,10000000000sut && \
vpnd gentx --name genesis && \
vpnd collect-gentxs
```
