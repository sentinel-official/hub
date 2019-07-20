// nolint: gochecknoglobals
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DepositKeyPrefix = []byte{0x01}
)

func DepositKey(address sdk.AccAddress) []byte {
	return append(DepositKeyPrefix, address.Bytes()...)
}

const (
	ModuleName   = "deposit"
	StoreKey     = ModuleName
	QuerierRoute = ModuleName
)
