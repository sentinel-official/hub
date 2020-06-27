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

var (
	SubscriptionsCountKey = []byte{0x00}
	SubscriptionKeyPrefix = []byte{0x01}
)

func SubscriptionKey(i uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionIDForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append([]byte{0x02}, address.Bytes()...)
}

func SubscriptionIDForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(SubscriptionIDForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionIDForPlanKeyPrefix(id uint64) []byte {
	return append([]byte{0x03}, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionIDForPlanKey(id, i uint64) []byte {
	return append(SubscriptionIDForPlanKeyPrefix(id), sdk.Uint64ToBigEndian(i)...)
}

func SubscriptionIDForNodeKeyPrefix(address hub.NodeAddress) []byte {
	return append([]byte{0x04}, address.Bytes()...)
}

func SubscriptionIDForNodeKey(address hub.NodeAddress, i uint64) []byte {
	return append(SubscriptionIDForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func AddressForSubscriptionIDKeyPrefix(id uint64) []byte {
	return append([]byte{0x05}, sdk.Uint64ToBigEndian(id)...)
}

func AddressForSubscriptionIDKey(id uint64, address sdk.AccAddress) []byte {
	return append(AddressForSubscriptionIDKeyPrefix(id), address.Bytes()...)
}
