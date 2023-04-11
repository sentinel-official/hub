package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName = "node"
)

var (
	TypeMsgRegisterRequest  = ModuleName + ":register"
	TypeMsgUpdateRequest    = ModuleName + ":update"
	TypeMsgSetStatusRequest = ModuleName + ":set_status"
)

var (
	NodeKeyPrefix           = []byte{0x10}
	ActiveNodeKeyPrefix     = append(NodeKeyPrefix, 0x11)
	InactiveNodeKeyPrefix   = append(NodeKeyPrefix, 0x12)
	InactiveNodeAtKeyPrefix = []byte{0x21}
)

func ActiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func GetInactiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveNodeAtKey(at time.Time, addr hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeAtKeyPrefix(at), address.MustLengthPrefix(addr.Bytes())...)
}

func AddressFromInactiveNodeAtKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | at (29 bytes) | addrLen (1 byte) | addr

	addrLen := int(key[30])
	if len(key) != 31+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 31+addrLen))
	}

	return key[31:]
}
