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
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	CountKey = []byte{0x00}

	PlanKeyPrefix = []byte{0x10}

	ActivePlanKeyPrefix   = []byte{0x20}
	InactivePlanKeyPrefix = []byte{0x21}

	ActivePlanForProviderKeyPrefix   = []byte{0x30}
	InactivePlanForProviderKeyPrefix = []byte{0x31}

	NodeForPlanKeyPrefix = []byte{0x40}
)

func PlanKey(i uint64) []byte {
	return append(PlanKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func ActivePlanKey(i uint64) []byte {
	return append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func InactivePlanKey(i uint64) []byte {
	return append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func GetActivePlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(ActivePlanForProviderKeyPrefix, address.Bytes()...)
}

func ActivePlanForProviderKey(address hub.ProvAddress, id uint64) []byte {
	return append(GetActivePlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetInactivePlanForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(InactivePlanForProviderKeyPrefix, address.Bytes()...)
}

func InactivePlanForProviderKey(address hub.ProvAddress, id uint64) []byte {
	return append(GetInactivePlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, address hub.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), address.Bytes()...)
}
