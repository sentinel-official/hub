package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "session"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	SessionsCountKey = []byte{0x00}
	SessionKeyPrefix = []byte{0x01}
)

func SessionKey(i uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(i)...)
}

func ActiveSessionIDKey(subscription uint64, node hub.NodeAddress, address sdk.AccAddress) []byte {
	return append(sdk.Uint64ToBigEndian(subscription), append(node.Bytes(), address.Bytes()...)...)
}
