package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	EventModuleName = EventModule{Name: ModuleName}
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
	v := append(PlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func ActivePlanKey(id uint64) []byte {
	v := append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func InactivePlanKey(id uint64) []byte {
	v := append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func GetActivePlanForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	v := append(ActivePlanForProviderKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func ActivePlanForProviderKey(address hubtypes.ProvAddress, id uint64) []byte {
	v := append(GetActivePlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetInactivePlanForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	v := append(InactivePlanForProviderKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func InactivePlanForProviderKey(address hubtypes.ProvAddress, id uint64) []byte {
	v := append(GetInactivePlanForProviderKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	v := append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func NodeForPlanKey(id uint64, address hubtypes.NodeAddress) []byte {
	v := append(GetNodeForPlanKeyPrefix(id), address.Bytes()...)
	if len(v) != 1+8+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8+sdk.AddrLen))
	}

	return v
}

func CountForNodeByProviderKey(p hubtypes.ProvAddress, n hubtypes.NodeAddress) []byte {
	v := append(CountForNodeByProviderKeyPrefix, append(p.Bytes(), n.Bytes()...)...)
	if len(v) != 1+2*sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+2*sdk.AddrLen))
	}

	return v
}

func IDFromStatusPlanKey(key []byte) uint64 {
	if len(key) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+8))
	}

	return sdk.BigEndianToUint64(key[1:])
}

func IDFromStatusPlanForProviderKey(key []byte) uint64 {
	if len(key) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+sdk.AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func AddressFromNodeForPlanKey(key []byte) hubtypes.NodeAddress {
	if len(key) != 1+8+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+8+sdk.AddrLen))
	}

	return key[1+8:]
}
