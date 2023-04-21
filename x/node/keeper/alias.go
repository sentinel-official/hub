package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (k *Keeper) FundCommunityPool(ctx sdk.Context, fromAddr sdk.AccAddress, coin sdk.Coin) error {
	if !coin.IsPositive() {
		return nil
	}

	coins := sdk.NewCoins(coin)
	return k.distribution.FundCommunityPool(ctx, coins, fromAddr)
}

func (k *Keeper) HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}
