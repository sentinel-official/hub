package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	EmptyProviderAddress = ProvAddress(make([]byte, sdk.AddrLen))
	EmptyCoins           = sdk.Coins{
		sdk.Coin{
			Denom:  "",
			Amount: sdk.NewInt(0),
		},
	}
)

func AreEmptyCoins(coins sdk.Coins) bool {
	for _, coin := range coins {
		if coin.Denom == "" && coin.Amount.Int64() == 0 {
			continue
		} else {
			return false
		}
	}

	return true
}
