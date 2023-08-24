package utils

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/types"
)

func AmountForBytes(gigabytePrice, bytes sdkmath.Int) sdkmath.Int {
	bytePrice := gigabytePrice.ToDec().QuoInt(types.Gigabyte)
	return bytes.ToDec().Mul(bytePrice).Ceil().TruncateInt()
}

func GetProportionOfCoin(coin sdk.Coin, share sdk.Dec) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		coin.Amount.ToDec().Mul(share).RoundInt(),
	)
}
