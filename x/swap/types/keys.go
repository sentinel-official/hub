package types

import (
	"fmt"

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
	EventModuleName = EventModule{Name: ModuleName}
)

var (
	PrecisionLoss = sdk.NewInt(100)
)

var (
	SwapKeyPrefix = []byte{0x10}
)

func SwapKey(hash EthereumHash) []byte {
	v := append(SwapKeyPrefix, hash.Bytes()...)
	if len(v) != 1+EthereumHashLength {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+EthereumHashLength))
	}

	return v
}
