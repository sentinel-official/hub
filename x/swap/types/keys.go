package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "swap"
)

var (
	ParamsSubspace = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgSwapRequest = ModuleName + ":swap"
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
