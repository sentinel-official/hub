package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "node"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ""
)

var (
	NodeKeyPrefix = []byte{0x00}
)

func NodeKey(address hub.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func NodeAddressForProviderKeyPrefix(pa hub.ProvAddress) []byte {
	return append([]byte{0x01}, pa.Bytes()...)
}

func NodeAddressForProviderKey(pa hub.ProvAddress, na hub.NodeAddress) []byte {
	return append(NodeAddressForProviderKeyPrefix(pa), na.Bytes()...)
}
