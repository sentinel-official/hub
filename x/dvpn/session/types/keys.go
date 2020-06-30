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
	SessionsCountKey       = []byte{0x00}
	SessionKeyPrefix       = []byte{0x01}
	ActiveSessionKeyPrefix = []byte{0x02}
)

func SessionKey(id uint64) []byte {
	return append(SessionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func ActiveSessionKey(s uint64, n hub.NodeAddress, a sdk.AccAddress) []byte {
	return append(ActiveSessionKeyPrefix,
		append(sdk.Uint64ToBigEndian(s),
			append(n.Bytes(), a.Bytes()...)...)...)
}
