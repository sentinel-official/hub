package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName     = "subscription"
	ParamsSubspace = ModuleName
	QuerierRoute   = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	CountKey                        = []byte{0x00}
	SubscriptionKeyPrefix           = []byte{0x01}
	SubscriptionForAddressKeyPrefix = []byte{0x02}
	SubscriptionForPlanKeyPrefix    = []byte{0x03}
	SubscriptionForNodeKeyPrefix    = []byte{0x04}
	CancelSubscriptionAtKeyPrefix   = []byte{0x05}

	QuotaKeyPrefix = []byte{0x10}
)

func SubscriptionKey(id uint64) []byte {
	return append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(SubscriptionForAddressKeyPrefix, address.Bytes()...)
}

func SubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	return append(GetSubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
}

func GetSubscriptionForPlanKeyPrefix(id uint64) []byte {
	return append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SubscriptionForPlanKey(p, s uint64) []byte {
	return append(GetSubscriptionForPlanKeyPrefix(p), sdk.Uint64ToBigEndian(s)...)
}

func GetSubscriptionForNodeKeyPrefix(address hub.NodeAddress) []byte {
	return append(SubscriptionForNodeKeyPrefix, address.Bytes()...)
}

func SubscriptionForNodeKey(address hub.NodeAddress, id uint64) []byte {
	return append(GetSubscriptionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetCancelSubscriptionAtKeyPrefix(at time.Time) []byte {
	return append(CancelSubscriptionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func CancelSubscriptionAtKey(at time.Time, id uint64) []byte {
	return append(GetCancelSubscriptionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetQuotaKeyPrefix(id uint64) []byte {
	return append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func QuotaKey(id uint64, address sdk.AccAddress) []byte {
	return append(GetQuotaKeyPrefix(id), address.Bytes()...)
}

func IDFromSubscriptionForAddressKey(key []byte) (i uint64) {
	binary.BigEndian.PutUint64(key[1+sdk.AddrLen:], i)
	return
}

func IDFromSubscriptionForPlanKey(key []byte) (i uint64) {
	binary.BigEndian.PutUint64(key[1+8:], i)
	return
}

func IDFromSubscriptionForNodeKey(key []byte) (i uint64) {
	binary.BigEndian.PutUint64(key[1+sdk.AddrLen:], i)
	return
}

func IDFromCancelSubscriptionAtKey(key []byte) (i uint64) {
	binary.BigEndian.PutUint64(key[1+29:], i)
	return
}
