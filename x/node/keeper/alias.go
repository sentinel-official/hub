package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) FundCommunityPool(ctx sdk.Context, fromAddr sdk.AccAddress, coin sdk.Coin) error {
	if !coin.IsPositive() {
		return nil
	}

	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), fromAddr)
}

func (k *Keeper) HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) CreateSubscriptionForNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, gigabytes, hours int64, denom string) (*subscriptiontypes.NodeSubscription, error) {
	return k.subscription.CreateSubscriptionForNode(ctx, accAddr, nodeAddr, gigabytes, hours, denom)
}
