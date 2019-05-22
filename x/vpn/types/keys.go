package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
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

func NodeKey(i uint64) []byte {
	return append(NodeKeyPrefix, csdkTypes.Uint64ToBigEndian(i)...)
}

func NodesCountOfAddressKey(address csdkTypes.AccAddress) []byte {
	return append(NodesCountOfAddressKeyPrefix, address.Bytes()...)
}

func NodeIDByAddressKey(address csdkTypes.AccAddress, i uint64) []byte {
	return append(NodeIDByAddressKeyPrefix,
		append(address.Bytes(), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func SubscriptionKey(i uint64) []byte {
	return append(SubscriptionKeyPrefix, csdkTypes.Uint64ToBigEndian(i)...)
}

func SubscriptionIDByNodeIDKey(i, j uint64) []byte {
	return append(SubscriptionIDByNodeIDKeyPrefix,
		append(csdkTypes.Uint64ToBigEndian(i), csdkTypes.Uint64ToBigEndian(j)...)...)
}

func SubscriptionsCountOfAddressKey(address csdkTypes.AccAddress) []byte {
	return append(SubscriptionsCountOfAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionIDByAddressKey(address csdkTypes.AccAddress, i uint64) []byte {
	return append(SubscriptionIDByAddressKeyPrefix,
		append(address.Bytes(), csdkTypes.Uint64ToBigEndian(i)...)...)
}

func SessionKey(i uint64) []byte {
	return append(SessionKeyPrefix, csdkTypes.Uint64ToBigEndian(i)...)
}

func SessionIDBySubscriptionIDKey(i, j uint64) []byte {
	return append(SessionIDBySubscriptionIDKeyPrefix,
		append(csdkTypes.Uint64ToBigEndian(i), csdkTypes.Uint64ToBigEndian(j)...)...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}
