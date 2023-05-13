package types

import (
	"fmt"
	"time"

	hubtypes "github.com/sentinel-official/hub/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "subscription"
	Year       = 365.25 * 24 * time.Hour
)

var (
	TypeMsgCancelRequest = ModuleName + ":cancel"
	TypeMsgShareRequest  = ModuleName + ":share"
)

var (
	CountKey = []byte{0x00}

	SubscriptionKeyPrefix            = []byte{0x10}
	SubscriptionForExpiryAtKeyPrefix = []byte{0x11}
	SubscriptionForAccountKeyPrefix  = []byte{0x12}
	SubscriptionForNodeKeyPrefix     = []byte{0x13}
	SubscriptionForPlanKeyPrefix     = []byte{0x14}

	QuotaKeyPrefix = []byte{0x20}
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

func GetSubscriptionForExpiryAtKeyPrefix(at time.Time) []byte {
	return append(SubscriptionForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func SubscriptionForExpiryAtKey(at time.Time, id uint64) []byte {
	return append(GetSubscriptionForExpiryAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetQuotaKeyPrefix(id uint64) []byte {
	return append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func QuotaKey(id uint64, addr sdk.AccAddress) []byte {
	return append(GetQuotaKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
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

func IDFromSubscriptionForExpiryAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
