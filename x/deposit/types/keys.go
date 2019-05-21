// nolint: gochecknoglobals
package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	DepositKeyPrefix = []byte{0x01}
)

func DepositKey(address csdkTypes.AccAddress) []byte {
	return append(DepositKeyPrefix, address.Bytes()...)
}

const (
	ModuleName = "deposit"
	StoreKey   = ModuleName
)
