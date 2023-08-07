package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "swap"
	StoreKey   = ModuleName
)

var (
	PrecisionLoss = sdk.NewInt(100)
)

var (
	SwapKeyPrefix = []byte{0x10}
)

func SwapKey(hash EthereumHash) []byte {
	return append(SwapKeyPrefix, hash.Bytes()...)
}
