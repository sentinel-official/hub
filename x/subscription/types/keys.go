package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
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
	EventModuleName = EventModule{Name: ModuleName}
)

var (
	CountKey                                = []byte{0x00}
	SubscriptionKeyPrefix                   = []byte{0x10}
	SubscriptionForNodeKeyPrefix            = []byte{0x11}
	SubscriptionForPlanKeyPrefix            = []byte{0x12}
	ActiveSubscriptionForAddressKeyPrefix   = []byte{0x20}
	InactiveSubscriptionForAddressKeyPrefix = []byte{0x21}
	InactiveSubscriptionAtKeyPrefix         = []byte{0x30}
	QuotaKeyPrefix                          = []byte{0x40}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForNodeKeyPrefix(address hubtypes.NodeAddress) []byte {
	return append(SubscriptionForNodeKeyPrefix, address.Bytes()...)
}

func SubscriptionForNodeKey(address hubtypes.NodeAddress, id uint64) []byte {
	return append(GetSubscriptionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForPlanKeyPrefix(id uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(p, s uint64) []byte {
	return append(GetSubscriptionForPlanKeyPrefix(p), sdk.Uint64ToBigEndian(s)...)
}

func GetActiveSubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(ActiveSubscriptionForAddressKeyPrefix, address.Bytes()...)
}

func ActiveSubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(GetActiveSubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func GetInactiveSubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(InactiveSubscriptionForAddressKeyPrefix, address.Bytes()...)
}

func InactiveSubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(GetInactiveSubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
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

func QuotaKey(id uint64, address sdk.AccAddress) []byte {
	return append(GetQuotaKeyPrefix(id), address.Bytes()...)
}

func IDFromSubscriptionForNodeKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromSubscriptionForPlanKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+8:])
}

func IDFromStatusSubscriptionForAddressKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromInactiveSubscriptionAtKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+29:])
}
