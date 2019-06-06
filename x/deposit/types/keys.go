// nolint: gochecknoglobals
package types

import (
	csdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DepositKeyPrefix = []byte{0x01}
)

func DepositKey(address csdk.AccAddress) []byte {
	return append(DepositKeyPrefix, address.Bytes()...)
}

const (
	ModuleName   = "deposit"
	StoreKey     = ModuleName
	QuerierRoute = ModuleName
)
