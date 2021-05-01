package types

import (
	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "provider"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	EventModuleName = EventModule{Name: ModuleName}
)

var (
	ProviderKeyPrefix = []byte{0x10}
)

func ProviderKey(address hubtypes.ProvAddress) []byte {
	return append(ProviderKeyPrefix, address.Bytes()...)
}
