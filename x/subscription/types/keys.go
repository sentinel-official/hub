package types

import (
	"fmt"
	hubtypes "github.com/sentinel-official/hub/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "subscription"
)

var (
	TypeMsgCancelRequest      = ModuleName + ":cancel"
	TypeMsgShareRequest       = ModuleName + ":share"
	TypeMsgUpdateQuotaRequest = ModuleName + ":update_quota"
)

var (
	CountKey = []byte{0x00}

	SubscriptionKeyPrefix           = []byte{0x10}
	SubscriptionForAddressKeyPrefix = []byte{0x11}
	SubscriptionForNodeKeyPrefix    = []byte{0x12}
	SubscriptionForPlanKeyPrefix    = []byte{0x13}

	InactiveSubscriptionAtKeyPrefix = []byte{0x20}

	QuotaKeyPrefix = []byte{0x30}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForAddressKeyPrefix(addr sdk.AccAddress) []byte {
	return append(SubscriptionForAddressKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func SubscriptionForAddressKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetSubscriptionForAddressKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
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

func GetInactiveSubscriptionAtKeyPrefix(at time.Time) []byte {
	return append(InactiveSubscriptionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveSubscriptionAtKey(at time.Time, id uint64) []byte {
	return append(GetInactiveSubscriptionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetQuotaKeyPrefix(id uint64) []byte {
	return append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func QuotaKey(id uint64, addr sdk.AccAddress) []byte {
	return append(GetQuotaKeyPrefix(id), address.MustLengthPrefix(addr.Bytes())...)
}

func IDFromStatusSubscriptionForAddressKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr | subscription (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromInactiveSubscriptionAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | subscription (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
