package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "session"
	QuerierRoute = ModuleName
	AddrLen      = 20
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
	CountKey                           = []byte{0x00}
	SessionKeyPrefix                   = []byte{0x11}
	InactiveSessionForAddressKeyPrefix = []byte{0x30}
	ActiveSessionForAddressKeyPrefix   = []byte{0x31}
	InactiveSessionAtKeyPrefix         = []byte{0x40}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(InactiveSessionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func InactiveSessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	return append(GetInactiveSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetActiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(ActiveSessionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
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
	if len(key) != 1+2*8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+2*8))
	}

	return sdk.BigEndianToUint64(key[1+8:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	if len(key) != 1+AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromStatusSessionForAddressKey(key []byte) uint64 {
	if len(key) != 1+AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromActiveSessionAtKey(key []byte) uint64 {
	if len(key) != 1+29+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+29+8))
	}

	return sdk.BigEndianToUint64(key[1+29:])
}
