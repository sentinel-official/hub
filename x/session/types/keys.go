package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
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
	SessionForExpiryAtKeyPrefix     = []byte{0x11}
	SessionForAccountKeyPrefix      = []byte{0x12}
	SessionForNodeKeyPrefix         = []byte{0x13}
	SessionForSubscriptionKeyPrefix = []byte{0x14}
	SessionForQuotaKeyPrefix        = []byte{0x15}
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

func GetSessionForQuotaKeyPrefix(id uint64, addr sdk.AccAddress) []byte {
	return append(SessionForQuotaKeyPrefix, append(sdk.Uint64ToBigEndian(id), address.MustLengthPrefix(addr)...)...)
}

func SessionForQuotaKey(subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) []byte {
	return append(GetSessionForQuotaKeyPrefix(subscriptionID, addr), sdk.Uint64ToBigEndian(sessionID)...)
}

func GetSessionForExpiryAtKeyPrefix(at time.Time) []byte {
	return append(SessionForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func SessionForExpiryAtKey(at time.Time, id uint64) []byte {
	return append(GetSessionForExpiryAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
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

func IDFromSessionForQuotaKey(key []byte) uint64 {
	// prefix (1 byte) | subscriptionID (8 bytes) | addrLen (1 byte) | addr (addrLen bytes) | id (8 bytes)

	addrLen := int(key[9])
	if len(key) != 18+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 18+addrLen))
	}

	return sdk.BigEndianToUint64(key[10+addrLen:])
}

func IDFromStatusSessionAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | session (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
