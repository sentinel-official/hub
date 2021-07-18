package types

import (
	"fmt"
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
	v := append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func GetSessionForSubscriptionKeyPrefix(id uint64) []byte {
	v := append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func SessionForSubscriptionKey(subscription, id uint64) []byte {
	v := append(GetSessionForSubscriptionKeyPrefix(subscription), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+2*8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+2*8))
	}

	return v
}

func GetSessionForNodeKeyPrefix(address hubtypes.NodeAddress) []byte {
	v := append(SessionForNodeKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func SessionForNodeKey(address hubtypes.NodeAddress, id uint64) []byte {
	v := append(GetSessionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetInactiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(InactiveSessionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func InactiveSessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	v := append(GetInactiveSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetActiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(ActiveSessionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func ActiveSessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	v := append(GetActiveSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetInactiveSessionAtKeyPrefix(at time.Time) []byte {
	v := append(InactiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
	if len(v) != 1+29 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+29))
	}

	return v
}

func InactiveSessionAtKey(at time.Time, id uint64) []byte {
	v := append(GetInactiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+29+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+29+8))
	}

	return v
}

func IDFromSessionForSubscriptionKey(key []byte) uint64 {
	if len(key) != 1+2*8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+2*8))
	}

	return sdk.BigEndianToUint64(key[1+8:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	if len(key) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+sdk.AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromStatusSessionForAddressKey(key []byte) uint64 {
	if len(key) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+sdk.AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromActiveSessionAtKey(key []byte) uint64 {
	if len(key) != 1+29+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+29+8))
	}

	return sdk.BigEndianToUint64(key[1+29:])
}
