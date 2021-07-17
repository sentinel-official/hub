package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RandomCoins(r *rand.Rand, coins sdk.Coins) sdk.Coins {
	if len(coins) == 0 {
		return nil
	}

	items := make(sdk.Coins, 0, r.Intn(len(coins)))
	for _, coin := range coins {
		items = append(
			items,
			sdk.NewInt64Coin(
				coin.Denom,
				r.Int63n(coin.Amount.Int64()),
			),
		)
	}

	return items
}
