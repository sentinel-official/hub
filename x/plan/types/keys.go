package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/v1/types"
)

const (
	ModuleName = "plan"
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

func PlanForProviderKey(addr hubtypes.ProvAddress, id uint64) []byte {
	return append(GetPlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func IDFromPlanForProviderKey(key []byte) uint64 {
	// prefix (1 bytes) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}
