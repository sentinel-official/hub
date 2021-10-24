package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName        = "swap"
	QuerierRoute      = ModuleName
	DefaultParamspace = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
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
