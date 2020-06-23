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
	PlansCountKeyPrefix  = []byte{0x00}
	PlanKeyPrefix        = []byte{0x01}
	NodeAddressKeyPrefix = []byte{0x02}
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

func NodeAddressForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(NodeAddressKeyPrefix, address.Bytes()...)
}

func NodeAddressForPlanKeyPrefix(address hub.ProvAddress, i uint64) []byte {
	return append(NodeAddressForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func NodeAddressKey(pa hub.ProvAddress, i uint64, na hub.NodeAddress) []byte {
	return append(NodeAddressForPlanKeyPrefix(pa, i), na.Bytes()...)
}
