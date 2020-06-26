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
	PlansCountKey = []byte{0x00}
	PlanKeyPrefix = []byte{0x01}

	SubscriptionsCountKey = []byte{0x00}
	SubscriptionKeyPrefix = []byte{0x01}
)

func PlanKey(i uint64) []byte {
	return append(PlanKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func PlanIDForProviderKey(address hub.ProvAddress, i uint64) []byte {
	return append(address, sdk.Uint64ToBigEndian(i)...)
}

func NodeAddressForPlanKey(i uint64, address hub.NodeAddress) []byte {
	return append(sdk.Uint64ToBigEndian(i), address.Bytes()...)
}

func SubscriptionKey(i uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionIDForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(address.Bytes(), sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionIDForPlanKey(id, i uint64) []byte {
	return append(sdk.Uint64ToBigEndian(id), sdk.Uint64ToBigEndian(i)...)
}
