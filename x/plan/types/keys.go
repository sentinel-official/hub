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
	TypeMsgAddRequest        = ModuleName + ":add"
	TypeMsgSetStatusRequest  = ModuleName + ":set_status"
	TypeMsgAddNodeRequest    = ModuleName + ":add_node"
	TypeMsgRemoveNodeRequest = ModuleName + ":remove_node"
)

var (
	CountKey                 = []byte{0x00}
	PlanKeyPrefix            = []byte{0x10}
	ActivePlanKeyPrefix      = append(PlanKeyPrefix, 0x01)
	InactivePlanKeyPrefix    = append(PlanKeyPrefix, 0x02)
	PlanForProviderKeyPrefix = []byte{0x20}
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

func IDFromStatusPlanKey(key []byte) uint64 {
	// prefix (1 byte) | plan (8 bytes)

	if len(key) != 9 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 9))
	}

	return sdk.BigEndianToUint64(key[1:])
}

func IDFromStatusPlanForProviderKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr | plan (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func AddressFromNodeForPlanKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | plan (8 bytes) | addrLen (1 byte) | addr

	addrLen := int(key[9])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return key[10:]
}
