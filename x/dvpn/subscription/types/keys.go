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
	PlansCountForProviderKeyPrefix = []byte{0x00}
	PlanKeyPrefix                  = []byte{0x01}
)

func PlansCountForProviderKey(address hub.ProvAddress) []byte {
	return append(PlansCountForProviderKeyPrefix, address.Bytes()...)
}

func PlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(PlanKeyPrefix, address.Bytes()...)
}

func PlanForProviderKey(address hub.ProvAddress, i uint64) []byte {
	return append(PlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}
