package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName     = "node"
	ParamsSubspace = ModuleName
	QuerierRoute   = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	NodeKeyPrefix            = []byte{0x00}
	NodeForProviderKeyPrefix = []byte{0x01}
	ActiveNodeAtKeyPrefix    = []byte{0x02}
)

func NodeKey(address hub.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func GetNodeForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(NodeForProviderKeyPrefix, address.Bytes()...)
}

func NodeForProviderKey(p hub.ProvAddress, n hub.NodeAddress) []byte {
	return append(GetNodeForProviderKeyPrefix(p), n.Bytes()...)
}

func GetActiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(ActiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func ActiveNodeAtKey(at time.Time, address hub.NodeAddress) []byte {
	return append(GetActiveNodeAtKeyPrefix(at), address.Bytes()...)
}
