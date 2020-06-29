package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "subscription"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	SubscriptionsCountKey = []byte{0x00}
	SubscriptionKeyPrefix = []byte{0x01}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append([]byte{0x02}, address.Bytes()...)
}

func SubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(SubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionForPlanKeyPrefix(id uint64) []byte {
	return append([]byte{0x03}, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(p, s uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix(p), sdk.Uint64ToBigEndian(s)...)
}

func SubscriptionForNodeKeyPrefix(address hub.NodeAddress) []byte {
	return append([]byte{0x04}, address.Bytes()...)
}

func SubscriptionForNodeKey(address hub.NodeAddress, id uint64) []byte {
	return append(SubscriptionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func MemberForSubscriptionKeyPrefix(id uint64) []byte {
	return append([]byte{0x05}, sdk.Uint64ToBigEndian(id)...)
}

func MemberForSubscriptionKey(id uint64, address sdk.AccAddress) []byte {
	return append(MemberForSubscriptionKeyPrefix(id), address.Bytes()...)
}
