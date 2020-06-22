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
	PlansCountKeyPrefix = []byte{0x00}
	PlanKeyPrefix       = []byte{0x01}
	NodeKeyPrefix       = []byte{0x02}
)

func PlansCountKey(address hub.ProvAddress) []byte {
	return append(PlansCountKeyPrefix, address.Bytes()...)
}

func PlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(PlanKeyPrefix, address.Bytes()...)
}

func PlanKey(address hub.ProvAddress, i uint64) []byte {
	return append(PlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func NodeForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func NodeForPlanKeyPrefix(address hub.ProvAddress, i uint64) []byte {
	return append(NodeForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func NodeKey(pa hub.ProvAddress, i uint64, na hub.NodeAddress) []byte {
	return append(NodeForPlanKeyPrefix(pa, i), na.Bytes()...)
}
