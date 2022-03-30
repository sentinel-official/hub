package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
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
	CountKey                           = []byte{0x00}
	SessionKeyPrefix                   = []byte{0x11}
	InactiveSessionForAddressKeyPrefix = []byte{0x30}
	ActiveSessionForAddressKeyPrefix   = []byte{0x31}
	InactiveSessionAtKeyPrefix         = []byte{0x40}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSessionForAddressKeyPrefix(addr sdk.AccAddress) []byte {
	return append(InactiveSessionForAddressKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveSessionForAddressKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetInactiveSessionForAddressKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetActiveSessionForAddressKeyPrefix(addr sdk.AccAddress) []byte {
	return append(ActiveSessionForAddressKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func ActiveSessionForAddressKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetActiveSessionForAddressKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetInactiveSessionAtKeyPrefix(at time.Time) []byte {
	return append(InactiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveSessionAtKey(at time.Time, id uint64) []byte {
	return append(GetInactiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func IDFromStatusSessionForAddressKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr | session (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromStatusSessionAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | session (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
