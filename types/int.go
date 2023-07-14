package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func AmountForBytes(gigabytePrice, bytes sdk.Int) sdk.Int {
	bytePrice := gigabytePrice.ToDec().QuoInt(Gigabyte)
	return bytes.ToDec().Mul(bytePrice).Ceil().TruncateInt()
}
