package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "provider"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ""
)

var (
	ProviderKeyPrefix = []byte{0x00}
)

func ProviderKey(address hub.ProvAddress) []byte {
	return append(ProviderKeyPrefix, address.Bytes()...)
}
