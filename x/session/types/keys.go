package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName     = "session"
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
	SessionKeyPrefix                = []byte{0x01}
	SessionForSubscriptionKeyPrefix = []byte{0x02}
	SessionForNodeKeyPrefix         = []byte{0x03}
	SessionForAddressKeyPrefix      = []byte{0x04}
	OngoingSessionKeyPrefix         = []byte{0x05}
	ActiveSessionAtKeyPrefix        = []byte{0x06}
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

func GetSessionForNodeKeyPrefix(address hub.NodeAddress) []byte {
	return append(SessionForNodeKeyPrefix, address.Bytes()...)
}

func SessionForNodeKey(address hub.NodeAddress, id uint64) []byte {
	return append(GetSessionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(SessionForAddressKeyPrefix, address.Bytes()...)
}

func SessionForAddressKey(address sdk.AccAddress, id uint64) []byte {
	return append(GetSessionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
}

func GetOngoingSessionPrefix(id uint64) []byte {
	return append(OngoingSessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func OngoingSessionKey(id uint64, address sdk.AccAddress) []byte {
	return append(GetOngoingSessionPrefix(id), address.Bytes()...)
}

func GetActiveSessionAtKeyPrefix(at time.Time) []byte {
	return append(ActiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func ActiveSessionAtKey(at time.Time, id uint64) []byte {
	return append(GetActiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}
