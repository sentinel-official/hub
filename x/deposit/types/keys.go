package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "deposit"
	QuerierRoute = ModuleName
	AddrLen = 20
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	DepositKeyPrefix = []byte{0x10}
)

func DepositKey(address sdk.AccAddress) []byte {
	v := append(DepositKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}
