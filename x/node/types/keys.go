package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "node"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	NodeKeyPrefix                    = []byte{0x10}
	ActiveNodeKeyPrefix              = []byte{0x20}
	InactiveNodeKeyPrefix            = []byte{0x21}
	ActiveNodeForProviderKeyPrefix   = []byte{0x30}
	InactiveNodeForProviderKeyPrefix = []byte{0x31}
	ActiveNodeAtKeyPrefix            = []byte{0x40}
	InactiveNodeAtKeyPrefix          = []byte{0x41}
)

func NodeKey(address hubtypes.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func ActiveNodeKey(address hubtypes.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, address.Bytes()...)
}

func InactiveNodeKey(address hubtypes.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, address.Bytes()...)
}

func GetActiveNodeForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	return append(ActiveNodeForProviderKeyPrefix, address.Bytes()...)
}

func ActiveNodeForProviderKey(provider hubtypes.ProvAddress, address hubtypes.NodeAddress) []byte {
	return append(GetActiveNodeForProviderKeyPrefix(provider), address.Bytes()...)
}

func GetInactiveNodeForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	return append(InactiveNodeForProviderKeyPrefix, address.Bytes()...)
}

func InactiveNodeForProviderKey(provider hubtypes.ProvAddress, address hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeForProviderKeyPrefix(provider), address.Bytes()...)
}

func GetActiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(ActiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func ActiveNodeAtKey(at time.Time, address hubtypes.NodeAddress) []byte {
	return append(GetActiveNodeAtKeyPrefix(at), address.Bytes()...)
}

func GetInactiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveNodeAtKey(at time.Time, address hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeAtKeyPrefix(at), address.Bytes()...)
}

func AddressFromStatusNodeKey(key []byte) hubtypes.NodeAddress {
	return key[1:]
}

func AddressFromStatusNodeForProviderKey(key []byte) hubtypes.NodeAddress {
	return key[1+sdk.AddrLen:]
}

func AddressFromStatusNodeAtKey(key []byte) hubtypes.NodeAddress {
	return key[1+29:]
}
