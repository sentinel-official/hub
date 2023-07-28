package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetProportionOfCoin(coin sdk.Coin, share sdk.Dec) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		coin.Amount.ToDec().Mul(share).RoundInt(),
	)
}
