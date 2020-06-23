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
