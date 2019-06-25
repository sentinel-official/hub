# Sentinel Hub

[![](https://img.shields.io/github/release-pre/sentinel-official/hub.svg?style=flat)](https://github.com/sentinel-official/hub/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sentinel-official/hub)](https://goreportcard.com/report/github.com/sentinel-official/hub)
[![](https://tokei.rs/b1/github/sentinel-official/hub)](https://github.com/sentinel-official/hub)

```
rm -rf $HOME/.sentinel-hubd $HOME/.sentinel-hubcli && \
sentinel-hubd init --chain-id hub testing && \
sentinel-hubcli keys add genesis && \
sentinel-hubd add-genesis-account $(sentinel-hubcli keys show genesis -a) 2000000000stake && \
sentinel-hubd gentx --name genesis && \
sentinel-hubd collect-gentxs
```

```
sentinel-hubd start --inv-check-period 1
```
