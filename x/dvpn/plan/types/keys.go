package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "plan"
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

func PlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append([]byte{0x02}, address.Bytes()...)
}

func PlanForProviderKey(address hub.ProvAddress, id uint64) []byte {
	return append(PlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func PlanForNodeKeyPrefix(address hub.NodeAddress) []byte {
	return append([]byte{0x03}, address.Bytes()...)
}

func PlanForNodeKey(address hub.NodeAddress, id uint64) []byte {
	return append(PlanForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKeyPrefix(id uint64) []byte {
	return append([]byte{0x04}, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, address hub.NodeAddress) []byte {
	return append(NodeForPlanKeyPrefix(id), address.Bytes()...)
}
