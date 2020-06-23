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
	NodeKeyPrefix        = []byte{0x00}
	NodeAddressKeyPrefix = []byte{0x01}
)

func NodeKey(address hub.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func NodeAddressForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(NodeAddressKeyPrefix, address.Bytes()...)
}

func NodeAddressForProviderKey(pa hub.ProvAddress, na hub.NodeAddress) []byte {
	return append(NodeAddressForProviderKeyPrefix(pa), na.Bytes()...)
}
