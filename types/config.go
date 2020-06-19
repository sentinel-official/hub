package types

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Config struct {
	*sdk.Config
	prefixes map[string]string
	sealed   bool
	mtx      sync.Mutex
}

var (
	config = &Config{
		Config: sdk.GetConfig(),
		prefixes: map[string]string{
			"provider_addr": Bech32PrefixProvAddr,
			"node_addr":     Bech32PrefixNodeAddr,
		},
	}
)

func GetConfig() *Config {
	return config
}

func (c *Config) assert() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.sealed {
		panic("config is sealed")
	}
}

func (c *Config) Seal() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.Config.Seal()
	c.sealed = true
}

func (c *Config) SetBech32PrefixForProvider(prefix string) {
	c.assert()
	config.prefixes["provider_addr"] = prefix
}

func (c *Config) SetBech32PrefixForNode(prefix string) {
	c.assert()
	config.prefixes["node_addr"] = prefix
}

func (c *Config) GetBech32ProviderAddrPrefix() string {
	return c.prefixes["provider_addr"]
}

func (c *Config) GetBech32NodeAddrPrefix() string {
	return c.prefixes["node_addr"]
}
