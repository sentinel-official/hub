package utils

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/types"
)

func AmountForBytes(gigabytePrice, bytes sdkmath.Int) sdkmath.Int {
	bytePrice := sdkmath.LegacyNewDecFromInt(gigabytePrice).QuoInt(types.Gigabyte)
	return sdkmath.LegacyNewDecFromInt(bytes).Mul(bytePrice).Ceil().TruncateInt()
}

func GetProportionOfCoin(coin sdk.Coin, share sdkmath.LegacyDec) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		sdkmath.LegacyNewDecFromInt(coin.Amount).Mul(share).RoundInt(),
	)
}
