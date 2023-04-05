package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "custommint"
	StoreKey   = ModuleName
)

var (
	InflationKeyPrefix = []byte{0x01}
)

func InflationKey(t time.Time) []byte {
	return append(InflationKeyPrefix, sdk.FormatTimeBytes(t)...)
}
