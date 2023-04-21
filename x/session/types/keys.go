package types

import (
	"fmt"
	hubtypes "github.com/sentinel-official/hub/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "session"
)

var (
	TypeMsgStartRequest         = ModuleName + ":start"
	TypeMsgUpdateDetailsRequest = ModuleName + ":update_details"
	TypeMsgEndRequest           = ModuleName + ":end"
)

var (
	CountKey = []byte{0x00}

	SessionKeyPrefix                = []byte{0x10}
	SessionForAccountKeyPrefix      = []byte{0x11}
	SessionForNodeKeyPrefix         = []byte{0x12}
	SessionForSubscriptionKeyPrefix = []byte{0x13}

	InactiveSessionAtKeyPrefix = []byte{0x20}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForAccountKeyPrefix(addr sdk.AccAddress) []byte {
	return append(SessionForAccountKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func SessionForAccountKey(addr sdk.AccAddress, id uint64) []byte {
	return append(GetSessionForAccountKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForNodeKeyPrefix(addr hubtypes.NodeAddress) []byte {
	return append(SessionForNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func SessionForNodeKey(addr hubtypes.NodeAddress, id uint64) []byte {
	return append(GetSessionForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForSubscriptionKeyPrefix(id uint64) []byte {
	return append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func SessionForSubscriptionKey(subscriptionID, sessionID uint64) []byte {
	return append(GetSessionForSubscriptionKeyPrefix(subscriptionID), sdk.Uint64ToBigEndian(sessionID)...)
}

func GetInactiveSessionAtKeyPrefix(at time.Time) []byte {
	return append(InactiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveSessionAtKey(at time.Time, id uint64) []byte {
	return append(GetInactiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func IDFromSessionForAccountKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	// prefix (1 byte) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[1])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return sdk.BigEndianToUint64(key[2+addrLen:])
}

func IDFromSessionForSubscriptionKey(key []byte) uint64 {
	// prefix (1 byte) | subscriptionID (8 bytes) | sessionID (8 bytes)

	if len(key) != 17 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 17))
	}

	return sdk.BigEndianToUint64(key[9:])
}

func IDFromStatusSessionAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | session (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
