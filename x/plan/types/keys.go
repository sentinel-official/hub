package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
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
	TypeMsgAddRequest        = ModuleName + ":add"
	TypeMsgSetStatusRequest  = ModuleName + ":set_status"
	TypeMsgAddNodeRequest    = ModuleName + ":add_node"
	TypeMsgRemoveNodeRequest = ModuleName + ":remove_node"
)

var (
	CountKey                         = []byte{0x00}
	PlanKeyPrefix                    = []byte{0x10}
	ActivePlanKeyPrefix              = []byte{0x20}
	InactivePlanKeyPrefix            = []byte{0x21}
	ActivePlanForProviderKeyPrefix   = []byte{0x30}
	InactivePlanForProviderKeyPrefix = []byte{0x31}
	NodeForPlanKeyPrefix             = []byte{0x40}
	CountForNodeByProviderKeyPrefix  = []byte{0x50}
)

func PlanKey(id uint64) []byte {
	return append(PlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func ActivePlanKey(id uint64) []byte {
	return append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func InactivePlanKey(id uint64) []byte {
	return append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetActivePlanForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(ActivePlanForProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func ActivePlanForProviderKey(addr hubtypes.ProvAddress, id uint64) []byte {
	return append(GetActivePlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetInactivePlanForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(InactivePlanForProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactivePlanForProviderKey(addr hubtypes.ProvAddress, id uint64) []byte {
	return append(GetInactivePlanForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, addr hubtypes.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func CountForNodeByProviderKey(provider hubtypes.ProvAddress, node hubtypes.NodeAddress) []byte {
	v := append(CountForNodeByProviderKeyPrefix, address.MustLengthPrefix(provider.Bytes())...)

	return append(v, address.MustLengthPrefix(node.Bytes())...)
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
