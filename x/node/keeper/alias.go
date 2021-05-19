package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	hubtypes "github.com/sentinel-official/hub/types"
)

// AccountKeeper Wrapper
func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) FundCommunityPool(ctx sdk.Context, from sdk.AccAddress, coin sdk.Coin) error {
	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), from)
}

func (k *Keeper) HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, address)
}

func (k *Keeper) GetCountForNodeByProvider(ctx sdk.Context, p hubtypes.ProvAddress, n hubtypes.NodeAddress) uint64 {
	return k.plan.GetCountForNodeByProvider(ctx, p, n)
}
