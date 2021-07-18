package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RandomCoin(r *rand.Rand, coin sdk.Coin) sdk.Coin {
	return sdk.NewInt64Coin(
		coin.Denom,
		r.Int63n(coin.Amount.Int64()),
	)
}

func RandomCoins(r *rand.Rand, coins sdk.Coins) sdk.Coins {
	if len(coins) == 0 {
		return nil
	}

	items := make(sdk.Coins, 0, r.Intn(len(coins)))
	for _, coin := range coins {
		items = append(
			items,
			RandomCoin(r, coin),
		)
	}

	return items
}
