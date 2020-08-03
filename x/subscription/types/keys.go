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
	SubscriptionsCountKey           = []byte{0x00}
	SubscriptionKeyPrefix           = []byte{0x01}
	SubscriptionForAddressKeyPrefix = []byte{0x02}
	SubscriptionForPlanKeyPrefix    = []byte{0x03}
	SubscriptionForNodeKeyPrefix    = []byte{0x04}
	QuotaForSubscriptionKeyPrefix   = []byte{0x05}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForAddressByAddressKey(address sdk.AccAddress) []byte {
	return append(SubscriptionForAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(SubscriptionForAddressByAddressKey(address), sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionForPlanByPlanKey(id uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(p, s uint64) []byte {
	return append(SubscriptionForPlanByPlanKey(p), sdk.Uint64ToBigEndian(s)...)
}

func SubscriptionForNodeByNodeKey(address hub.NodeAddress) []byte {
	return append(SubscriptionForNodeKeyPrefix, address.Bytes()...)
}

func SubscriptionForNodeKey(address hub.NodeAddress, id uint64) []byte {
	return append(SubscriptionForNodeByNodeKey(address), sdk.Uint64ToBigEndian(id)...)
}

func QuotaForSubscriptionBySubscriptionKey(id uint64) []byte {
	return append(QuotaForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func QuotaForSubscriptionKey(id uint64, address sdk.AccAddress) []byte {
	return append(QuotaForSubscriptionBySubscriptionKey(id), address.Bytes()...)
}
