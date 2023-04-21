package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName = "plan"
)

var (
	TypeMsgCreateRequest       = ModuleName + ":create"
	TypeMsgUpdateStatusRequest = ModuleName + ":update_status"
	TypeMsgLinkNodeRequest     = ModuleName + ":link_node"
	TypeMsgUnlinkNodeRequest   = ModuleName + ":unlink_node"
)

var (
	CountKey = []byte{0x00}

	PlanKeyPrefix            = []byte{0x10}
	ActivePlanKeyPrefix      = append(PlanKeyPrefix, 0x01)
	InactivePlanKeyPrefix    = append(PlanKeyPrefix, 0x02)
	PlanForProviderKeyPrefix = []byte{0x11}
)

func ActivePlanKey(id uint64) []byte {
	return append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func InactivePlanKey(id uint64) []byte {
	return append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetPlanForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(PlanForProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func GetActivePlanForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(GetPlanForProviderKeyPrefix(addr), 0x01)
}

func ActivePlanForProviderKey(addr hubtypes.ProvAddress, id uint64) []byte {
	return append(GetActivePlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetInactivePlanForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(GetPlanForProviderKeyPrefix(addr), 0x02)
}

func InactivePlanForProviderKey(addr hubtypes.ProvAddress, id uint64) []byte {
	return append(GetInactivePlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func IDFromPlanForProviderKey(key []byte) uint64 {
	// prefix (2 bytes) | addrLen (1 byte) | addr | id (8 bytes)

	addrLen := int(key[2])
	if len(key) != 11+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+addrLen))
	}

	return sdk.BigEndianToUint64(key[3+addrLen:])
}
