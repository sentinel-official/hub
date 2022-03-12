package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName   = "subscription"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgSubscribeToNodeRequest = ModuleName + ":subscribe_to_node"
	TypeMsgSubscribeToPlanRequest = ModuleName + ":subscribe_to_plan"
	TypeMsgCancelRequest          = ModuleName + ":cancel"
	TypeMsgAddQuotaRequest        = ModuleName + ":add_quota"
	TypeMsgUpdateQuotaRequest     = ModuleName + ":update_quota"
)

var (
	CountKey                                = []byte{0x00}
	SubscriptionKeyPrefix                   = []byte{0x10}
	ActiveSubscriptionForAddressKeyPrefix   = []byte{0x20}
	InactiveSubscriptionForAddressKeyPrefix = []byte{0x21}
	InactiveSubscriptionAtKeyPrefix         = []byte{0x30}
	QuotaKeyPrefix                          = []byte{0x40}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetActiveSubscriptionForAddressKeyPrefix(addr sdk.AccAddress) []byte {
	return append(ActiveSubscriptionForAddressKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func ActiveSubscriptionForAddressKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetActiveSubscriptionForAddressKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSubscriptionForAddressKeyPrefix(addr sdk.AccAddress) []byte {
	return append(InactiveSubscriptionForAddressKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveSubscriptionForAddressKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetInactiveSubscriptionForAddressKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
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

func IDFromSubscriptionForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr | subscription (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSubscriptionForPlanKey(key []byte) uint64 {
	// prefix (1 byte) | plan (8 bytes) | subscription (8 bytes)

	if len(key) != 17 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 17))
	}

	return sdk.BigEndianToUint64(key[9:])
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
