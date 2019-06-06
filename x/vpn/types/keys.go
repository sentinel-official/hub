package types

import (
	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	ModuleName           = "vpn"
	StoreKeySession      = "vpnSession"
	StoreKeyNode         = "vpnNode"
	StoreKeySubscription = "vpnSubscription"
	QuerierRoute         = ModuleName
	RouterKey            = ModuleName

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeRegistered = "DE-REGISTERED"
)

// nolint: gochecknoglobals
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
)

func NodeKey(id sdk.ID) []byte {
	return append(NodeKeyPrefix, csdk.Uint64ToBigEndian(id.Uint64())...)
}

func NodesCountOfAddressKey(address csdk.AccAddress) []byte {
	return append(NodesCountOfAddressKeyPrefix, address.Bytes()...)
}

func NodeIDByAddressKey(address csdk.AccAddress, i uint64) []byte {
	return append(NodeIDByAddressKeyPrefix,
		append(address.Bytes(), csdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionKey(id sdk.ID) []byte {
	return append(SubscriptionKeyPrefix, csdk.Uint64ToBigEndian(id.Uint64())...)
}

func SubscriptionsCountOfNodeKey(id sdk.ID) []byte {
	return append(SubscriptionsCountOfNodeKeyPrefix, csdk.Uint64ToBigEndian(id.Uint64())...)
}

func SubscriptionIDByNodeIDKey(id sdk.ID, i uint64) []byte {
	return append(SubscriptionIDByNodeIDKeyPrefix,
		append(csdk.Uint64ToBigEndian(id.Uint64()), csdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionsCountOfAddressKey(address csdk.AccAddress) []byte {
	return append(SubscriptionsCountOfAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionIDByAddressKey(address csdk.AccAddress, i uint64) []byte {
	return append(SubscriptionIDByAddressKeyPrefix,
		append(address.Bytes(), csdk.Uint64ToBigEndian(i)...)...)
}

func SessionKey(id sdk.ID) []byte {
	return append(SessionKeyPrefix, csdk.Uint64ToBigEndian(id.Uint64())...)
}

func SessionsCountOfSubscriptionKey(id sdk.ID) []byte {
	return append(SessionsCountOfSubscriptionKeyPrefix, csdk.Uint64ToBigEndian(id.Uint64())...)
}

func SessionIDBySubscriptionIDKey(id sdk.ID, i uint64) []byte {
	return append(SessionIDBySubscriptionIDKeyPrefix,
		append(csdk.Uint64ToBigEndian(id.Uint64()), csdk.Uint64ToBigEndian(i)...)...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return csdk.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return csdk.Uint64ToBigEndian(uint64(height))
}
