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

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeRegistered = "DE-REGISTERED"
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
)

func NodeKey(id hub.ID) []byte {
	return append(NodeKeyPrefix, sdk.Uint64ToBigEndian(id.Uint64())...)
}

func NodesCountOfAddressKey(address sdk.AccAddress) []byte {
	return append(NodesCountOfAddressKeyPrefix, address.Bytes()...)
}

func NodeIDByAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(NodeIDByAddressKeyPrefix,
		append(address.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionKey(id hub.ID) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id.Uint64())...)
}

func SubscriptionsCountOfNodeKey(id hub.ID) []byte {
	return append(SubscriptionsCountOfNodeKeyPrefix, sdk.Uint64ToBigEndian(id.Uint64())...)
}

func SubscriptionIDByNodeIDKey(id hub.ID, i uint64) []byte {
	return append(SubscriptionIDByNodeIDKeyPrefix,
		append(sdk.Uint64ToBigEndian(id.Uint64()), sdk.Uint64ToBigEndian(i)...)...)
}

func SubscriptionsCountOfAddressKey(address sdk.AccAddress) []byte {
	return append(SubscriptionsCountOfAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionIDByAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(SubscriptionIDByAddressKeyPrefix,
		append(address.Bytes(), sdk.Uint64ToBigEndian(i)...)...)
}

func SessionKey(id hub.ID) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id.Uint64())...)
}

func SessionsCountOfSubscriptionKey(id hub.ID) []byte {
	return append(SessionsCountOfSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id.Uint64())...)
}

func SessionIDBySubscriptionIDKey(id hub.ID, i uint64) []byte {
	return append(SessionIDBySubscriptionIDKeyPrefix,
		append(sdk.Uint64ToBigEndian(id.Uint64()), sdk.Uint64ToBigEndian(i)...)...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(height))
}
