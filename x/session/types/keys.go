package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "session"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgStartRequest  = ModuleName + ":start"
	TypeMsgUpdateRequest = ModuleName + ":update"
	TypeMsgEndRequest    = ModuleName + ":end"
)

var (
	EventModuleName = EventModule{Name: ModuleName}
)

var (
	CountKey                           = []byte{0x00}
	SessionKeyPrefix                   = []byte{0x11}
	SessionForSubscriptionKeyPrefix    = []byte{0x20}
	SessionForNodeKeyPrefix            = []byte{0x21}
	InactiveSessionForAddressKeyPrefix = []byte{0x30}
	ActiveSessionForAddressKeyPrefix   = []byte{0x31}
	InactiveSessionAtKeyPrefix         = []byte{0x40}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForSubscriptionKeyPrefix(id uint64) []byte {
	return append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SessionForSubscriptionKey(subscription, id uint64) []byte {
	return append(GetSessionForSubscriptionKeyPrefix(subscription), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForNodeKeyPrefix(address hubtypes.NodeAddress) []byte {
	return append(SessionForNodeKeyPrefix, address.Bytes()...)
}

func SessionForNodeKey(address hubtypes.NodeAddress, id uint64) []byte {
	return append(GetSessionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(InactiveSessionForAddressKeyPrefix, address.Bytes()...)
}

func InactiveSessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	return append(GetInactiveSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetActiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(ActiveSessionForAddressKeyPrefix, address.Bytes()...)
}

func ActiveSessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	return append(GetActiveSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSessionAtKeyPrefix(at time.Time) []byte {
	return append(InactiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveSessionAtKey(at time.Time, id uint64) []byte {
	return append(GetInactiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func IDFromSessionForSubscriptionKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+8:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromStatusSessionForAddressKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromActiveSessionAtKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1+29:])
}
