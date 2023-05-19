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
	NodeKeyPrefix            = []byte{0x10}
	ActiveNodeKeyPrefix      = append(NodeKeyPrefix, 0x01)
	InactiveNodeKeyPrefix    = append(NodeKeyPrefix, 0x02)
	NodeForExpiryAtKeyPrefix = []byte{0x11}
	NodeForPlanKeyPrefix     = []byte{0x12}

	LeaseKeyPrefix                  = []byte{0x30}
	LeaseForDistributionAtKeyPrefix = []byte{0x31}
	LeaseForAccountKeyPrefix        = []byte{0x32}
	LeaseForNodeKeyPrefix           = []byte{0x33}
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

func NodeForPlanKey(id uint64, addr hubtypes.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func LeaseKey(id uint64) []byte {
	return append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForDistributionAtKeyPrefix(at time.Time) []byte {
	return append(LeaseForDistributionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeaseForDistributionAtKey(at time.Time, id uint64) []byte {
	return append(GetLeaseForDistributionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(LeaseForAccountKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func LeaseForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetLeaseForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForNodeKeyPrefix(addr hubtypes.NodeAddress) []byte {
	return append(LeaseForNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func LeaseForNodeKey(addr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetLeaseForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetNodeForExpiryAtKeyPrefix(at time.Time) []byte {
	return append(NodeForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func NodeForExpiryAtKey(at time.Time, addr hubtypes.NodeAddress) []byte {
	return append(GetNodeForExpiryAtKeyPrefix(at), address.MustLengthPrefix(addr.Bytes())...)
}

func AddressFromNodeForPlanKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | id (8 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[9])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return key[10:]
}

func AddressFromNodeForExpiryAtKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | at (29 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[30])
	if len(key) != 31+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 31+addrLen))
	}

	return key[31:]
}

func IDFromLeaseForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen(1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromLeaseForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen(1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromLeaseForDistributionAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
