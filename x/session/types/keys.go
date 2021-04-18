package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
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
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	CountKey                         = []byte{0x00}
	ChannelKeyPrefix                 = []byte{0x10}
	SessionKeyPrefix                 = []byte{0x11}
	SessionForSubscriptionKeyPrefix  = []byte{0x20}
	SessionForNodeKeyPrefix          = []byte{0x21}
	SessionForAddressKeyPrefix       = []byte{0x22}
	ActiveSessionAtKeyPrefix         = []byte{0x30}
	ActiveSessionForAddressKeyPrefix = []byte{0x31}
)

func GetChannelKeyPrefix(address sdk.AccAddress) []byte {
	return append(ChannelKeyPrefix, address.Bytes()...)
}

func ChannelKey(address sdk.AccAddress, subscription uint64, node hub.NodeAddress) []byte {
	return append(GetChannelKeyPrefix(address),
		append(sdk.Uint64ToBigEndian(subscription), node.Bytes()...)...)
}

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

func GetActiveSessionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	return append(ActiveSessionForAddressKeyPrefix, address.Bytes()...)
}

func ActiveSessionForAddressKey(address sdk.AccAddress, subscription uint64, node hub.NodeAddress) []byte {
	return append(GetActiveSessionForAddressKeyPrefix(address),
		append(sdk.Uint64ToBigEndian(subscription), node.Bytes()...)...)
}

func GetActiveSessionAtKeyPrefix(at time.Time) []byte {
	return append(ActiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func ActiveSessionAtKey(at time.Time, id uint64) []byte {
	return append(GetActiveSessionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func IDFromSessionForSubscriptionKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key[1+8:])
}

func IDFromSessionForNodeKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key[1+sdk.AddrLen:])
}

func IDFromSessionForAddressKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key[1+sdk.AddrLen:])
}

func IDFromActiveSessionAtKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key[1+29:])
}
