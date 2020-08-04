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
	CountKey                 = []byte{0x00}
	PlanKeyPrefix            = []byte{0x01}
	PlanForProviderKeyPrefix = []byte{0x02}

	NodeForPlanKeyPrefix = []byte{0x03}
)

func PlanKey(i uint64) []byte {
	return append(PlanKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func GetPlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(PlanForProviderKeyPrefix, address.Bytes()...)
}

func PlanForProviderKey(address hub.ProvAddress, id uint64) []byte {
	return append(GetPlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, address hub.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), address.Bytes()...)
}
