package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "vpn"
	QuerierRoute = ModuleName
	RouterKey    = ModuleName
	
	StoreKeySession      = "vpn_session"
	StoreKeyNode         = "vpn_node"
	StoreKeySubscription = "vpn_subscription"
	StoreKeyResolver     = "resolver_node"
	
	StatusRegistered   = "REGISTERED"
	StatusDeRegistered = "DE-REGISTERED"
	
	StatusActive   = "ACTIVE"
	StatusInactive = "INACTIVE"
)

var (
	NodesCountKey                = []byte{0x00}
	NodeKeyPrefix                = []byte{0x01}
	NodesCountOfAddressKeyPrefix = []byte{0x02}
	NodeIDByAddressKeyPrefix     = []byte{0x03}
	
	SubscriptionsCountKey                = []byte{0x00}
	SubscriptionKeyPrefix                = []byte{0x01}
	SubscriptionsCountOfNodeKeyPrefix    = []byte{0x02}
	SubscriptionIDByNodeIDKeyPrefix      = []byte{0x03}
	SubscriptionsCountOfAddressKeyPrefix = []byte{0x04}
	SubscriptionIDByAddressKeyPrefix     = []byte{0x05}
	
	SessionsCountKey                     = []byte{0x00}
	SessionKeyPrefix                     = []byte{0x01}
	SessionsCountOfSubscriptionKeyPrefix = []byte{0x02}
	SessionIDBySubscriptionIDKeyPrefix   = []byte{0x03}
	
	ResolverCountKey                = []byte{0x00}
	ResolverCountOfAddressKeyPrefix = []byte{0x01}
	ResolverKeyPrefix               = []byte{0x02}
	ResolverIDByAddressPrefix       = []byte{0x03}
	NodesOfResolverKeyPrefix        = []byte{0x04}
	ResolversOfNodeKeyPrefix        = []byte{0x05}
	
	FreeClientKey              = []byte{0x00}
	FreeNodesOfClientKeyPrefix = []byte{0x01}
	FreeClientOfNodeKeyPrefix  = []byte{0x02}
)

func NodeKey(id hub.NodeID) []byte {
	return append(NodeKeyPrefix, id.Bytes()...)
}

func NodesCountOfAddressKey(address sdk.AccAddress) []byte {
	return append(NodesCountOfAddressKeyPrefix, address.Bytes()...)
}

func NodeIDByAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(NodeIDByAddressKeyPrefix,
		append(address.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionKey(id hub.SubscriptionID) []byte {
	return append(SubscriptionKeyPrefix, id.Bytes()...)
}

func SubscriptionsCountOfNodeKey(id hub.NodeID) []byte {
	return append(SubscriptionsCountOfNodeKeyPrefix, id.Bytes()...)
}

func SubscriptionIDByNodeIDKey(id hub.NodeID, i uint64) []byte {
	return append(SubscriptionIDByNodeIDKeyPrefix,
		append(id.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionsCountOfAddressKey(address sdk.AccAddress) []byte {
	return append(SubscriptionsCountOfAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionIDByAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(SubscriptionIDByAddressKeyPrefix,
		append(address.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func SessionKey(id hub.SessionID) []byte {
	return append(SessionKeyPrefix, id.Bytes()...)
}

func SessionsCountOfSubscriptionKey(id hub.SubscriptionID) []byte {
	return append(SessionsCountOfSubscriptionKeyPrefix, id.Bytes()...)
}

func SessionIDBySubscriptionIDKey(id hub.SubscriptionID, i uint64) []byte {
	return append(SessionIDBySubscriptionIDKeyPrefix,
		append(id.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(height))
}

func FreeNodesOfClientKey(client sdk.AccAddress, nodeID hub.NodeID) []byte {
	return append(FreeNodesOfClientKeyPrefix, append(client.Bytes(), nodeID.Bytes()...)...)
}

func FreeClientOfNodeKey(nodeID hub.NodeID, client sdk.AccAddress) []byte {
	return append(FreeClientOfNodeKeyPrefix, append(nodeID.Bytes(), client.Bytes()...)...)
}

func ResolverKey(resolverID hub.ResolverID) []byte {
	return append(ResolverKeyPrefix, resolverID.Bytes()...)
}

func ResolversCountOfAddressKey(address sdk.AccAddress) []byte {
	return append(ResolverCountOfAddressKeyPrefix, address.Bytes()...)
}

func ResolverIDByAddressKey(address sdk.AccAddress, count uint64) []byte {
	return append(
		append(ResolverIDByAddressPrefix, address.Bytes()...),
		sdk.Uint64ToBigEndian(count)...)
}

func NodeOfResolverKey(resolverID hub.ResolverID, nodeID hub.NodeID) []byte {
	return append(NodesOfResolverKeyPrefix, append(resolverID.Bytes(), nodeID.Bytes()...)...)
}

func ResolverOfNodeKey(nodeID hub.NodeID, resolverID hub.ResolverID) []byte {
	return append(ResolversOfNodeKeyPrefix, append(nodeID.Bytes(), resolverID.Bytes()...)...)
}

func GetFreeClientKey(nodeID hub.NodeID) []byte {
	return append(FreeClientKey, nodeID.Bytes()...)
}
