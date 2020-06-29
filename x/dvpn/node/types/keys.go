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

func NodeForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append([]byte{0x01}, address.Bytes()...)
}

func NodeForProviderKey(p hub.ProvAddress, n hub.NodeAddress) []byte {
	return append(NodeForProviderKeyPrefix(p), n.Bytes()...)
}
