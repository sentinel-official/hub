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
	TypeMsgRegisterRequest      = ModuleName + ":register"
	TypeMsgUpdateDetailsRequest = ModuleName + ":update_details"
	TypeMsgUpdateStatusRequest  = ModuleName + ":update_status"
)

var (
	NodeKeyPrefix         = []byte{0x10}
	ActiveNodeKeyPrefix   = append(NodeKeyPrefix, 0x01)
	InactiveNodeKeyPrefix = append(NodeKeyPrefix, 0x02)
	NodeForPlanKeyPrefix  = []byte{0x11}

	InactiveNodeAtKeyPrefix = []byte{0x20}

	LeaseKeyPrefix           = []byte{0x30}
	LeaseForAccountKeyPrefix = []byte{0x31}
	LeaseForNodeKeyPrefix    = []byte{0x32}
)

func ActiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetActiveNodeForPlanKeyPrefix(id uint64) []byte {
	return append(GetNodeForPlanKeyPrefix(id), 0x01)
}

func ActiveNodeForPlanKey(id uint64, addr hubtypes.NodeAddress) []byte {
	return append(GetActiveNodeForPlanKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func GetInactiveNodeForPlanKeyPrefix(id uint64) []byte {
	return append(GetNodeForPlanKeyPrefix(id), 0x02)
}

func InactiveNodeForPlanKey(id uint64, addr hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeForPlanKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func LeaseKey(id uint64) []byte {
	return append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(LeaseForAccountKeyPrefix, address.MustLengthPrefix(addr)...)
}

func LeaseForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetLeaseForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForNodeKeyPrefix(addr hubtypes.NodeAddress) []byte {
	return append(LeaseForNodeKeyPrefix, address.MustLengthPrefix(addr)...)
}

func LeaseForNodeKey(addr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetLeaseForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
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
