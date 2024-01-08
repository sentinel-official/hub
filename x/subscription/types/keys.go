package types

import (
	"fmt"
	"time"

	hubtypes "github.com/sentinel-official/hub/v12/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "subscription"
	Day        = 24 * time.Hour
)

var (
	CountKey = []byte{0x00}

	SubscriptionKeyPrefix              = []byte{0x10}
	SubscriptionForInactiveAtKeyPrefix = []byte{0x11}
	SubscriptionForAccountKeyPrefix    = []byte{0x12}
	SubscriptionForNodeKeyPrefix       = []byte{0x13}
	SubscriptionForPlanKeyPrefix       = []byte{0x14}

	AllocationKeyPrefix = []byte{0x20}

	PayoutKeyPrefix                 = []byte{0x30}
	PayoutForNextAtKeyPrefix        = []byte{0x31}
	PayoutForAccountKeyPrefix       = []byte{0x32}
	PayoutForNodeKeyPrefix          = []byte{0x33}
	PayoutForAccountByNodeKeyPrefix = []byte{0x34}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(SubscriptionForAccountKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func SubscriptionForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetSubscriptionForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForNodeKeyPrefix(addr hubtypes.NodeAddress) []byte {
	return append(SubscriptionForNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func SubscriptionForNodeKey(addr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetSubscriptionForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForPlanKeyPrefix(id uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(planID, subscriptionID uint64) []byte {
	return append(GetSubscriptionForPlanKeyPrefix(planID), sdk.Uint64ToBigEndian(subscriptionID)...)
}

func GetSubscriptionForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(SubscriptionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func SubscriptionForInactiveAtKey(at time.Time, id uint64) []byte {
	return append(GetSubscriptionForInactiveAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetAllocationForSubscriptionKeyPrefix(id uint64) []byte {
	return append(AllocationKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func AllocationKey(id uint64, addr sdk.AccAddress) []byte {
	return append(GetAllocationForSubscriptionKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func PayoutKey(id uint64) []byte {
	return append(PayoutKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetPayoutForNextAtKeyPrefix(at time.Time) []byte {
	return append(PayoutForNextAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func PayoutForNextAtKey(at time.Time, id uint64) []byte {
	return append(GetPayoutForNextAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetPayoutForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(PayoutForAccountKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func PayoutForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetPayoutForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetPayoutForNodeKeyPrefix(addr hubtypes.NodeAddress) []byte {
	return append(PayoutForNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func PayoutForNodeKey(addr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetPayoutForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetPayoutForAccountByNodeKeyPrefix(accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress) (key []byte) {
	return append(append(PayoutForAccountByNodeKeyPrefix, address.MustLengthPrefix(accAddr.Bytes())...), address.MustLengthPrefix(nodeAddr.Bytes())...)
}

func PayoutForAccountByNodeKey(accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetPayoutForAccountByNodeKeyPrefix(accAddr, nodeAddr), sdk.Uint64ToBigEndian(id)...)
}

func IDFromSubscriptionForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSubscriptionForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSubscriptionForPlanKey(key []byte) uint64 {
	// prefix (1 byte) | planID (8 bytes) | subscriptionID (8 bytes)

	if len(key) != 17 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 17))
	}

	return sdk.BigEndianToUint64(key[9:])
}

func IDFromSubscriptionForInactiveAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}

func IDFromPayoutForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen(1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromPayoutForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen(1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromPayoutForAccountByNodeKey(key []byte) uint64 {
	// prefix (1 byte) | accAddrLen(1 byte) | accAddr (accAddrLen bytes) | nodeAddrLen(1 byte) | nodeAddr (nodeAddrLen bytes) | id (8 bytes)

	accAddrLen, nodeAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+accAddrLen+nodeAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+accAddrLen+nodeAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+accAddrLen+nodeAddrLen:])
}

func IDFromPayoutForNextAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
