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
	ProvidersCountKey = []byte{0x00}
	ProviderKeyPrefix = []byte{0x01}
)

func ProviderKey(id hub.ProviderID) []byte {
	return append(ProviderKeyPrefix, id.Bytes()...)
}
