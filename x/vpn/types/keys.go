package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
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
	SubscriptionIDByNodeIDKeyPrefix      = []byte{0x03}
	SubscriptionsCountOfAddressKeyPrefix = []byte{0x04}
	SubscriptionIDByAddressKeyPrefix     = []byte{0x05}

	SessionsCountKey                   = []byte{0x00}
	SessionKeyPrefix                   = []byte{0x01}
	SessionIDBySubscriptionIDKeyPrefix = []byte{0x03}
)

func NodeKey(id sdkTypes.ID) []byte {
	return append(NodeKeyPrefix, csdkTypes.Uint64ToBigEndian(id.UInt64())...)
}

func NodesCountOfAddressKey(address csdkTypes.AccAddress) []byte {
	return append(NodesCountOfAddressKeyPrefix, address.Bytes()...)
}

func NodeIDByAddressKey(address csdkTypes.AccAddress, i uint64) []byte {
	return append(NodeIDByAddressKeyPrefix,
		append(address.Bytes(), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func SubscriptionKey(id sdkTypes.ID) []byte {
	return append(SubscriptionKeyPrefix, csdkTypes.Uint64ToBigEndian(id.UInt64())...)
}

func SubscriptionIDByNodeIDKey(id sdkTypes.ID, i uint64) []byte {
	return append(SubscriptionIDByNodeIDKeyPrefix,
		append(csdkTypes.Uint64ToBigEndian(id.UInt64()), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func SubscriptionsCountOfAddressKey(address csdkTypes.AccAddress) []byte {
	return append(SubscriptionsCountOfAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionIDByAddressKey(address csdkTypes.AccAddress, i uint64) []byte {
	return append(SubscriptionIDByAddressKeyPrefix,
		append(address.Bytes(), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func SessionKey(id sdkTypes.ID) []byte {
	return append(SessionKeyPrefix, csdkTypes.Uint64ToBigEndian(id.UInt64())...)
}

func SessionIDBySubscriptionIDKey(id sdkTypes.ID, i uint64) []byte {
	return append(SessionIDBySubscriptionIDKeyPrefix,
		append(csdkTypes.Uint64ToBigEndian(id.UInt64()), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}
