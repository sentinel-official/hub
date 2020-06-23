package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "node"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	NodeKeyPrefix = []byte{0x00}
)

func NodeKey(address hub.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func NodeAddressForProviderKey(pa hub.ProvAddress, na hub.NodeAddress) []byte {
	return append(pa.Bytes(), na.Bytes()...)
}
