package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "deposit"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	DepositKeyPrefix = []byte{0x01}
)

func DepositKey(address sdk.AccAddress) []byte {
	return append(DepositKeyPrefix, address.Bytes()...)
}
