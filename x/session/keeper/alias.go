package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	nodetypes "github.com/sentinel-official/hub/x/node/types"

	hubtypes "github.com/sentinel-official/hub/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, address hubtypes.NodeAddress) bool {
	return k.plan.HasNodeForPlan(ctx, id, address)
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscriptiontypes.Subscription, bool) {
	return k.subscription.GetSubscription(ctx, id)
}

func (k *Keeper) HasSubscriptionForNode(ctx sdk.Context, address hubtypes.NodeAddress, id uint64) bool {
	return k.subscription.HasSubscriptionForNode(ctx, address, id)
}

func (k *Keeper) SetQuota(ctx sdk.Context, id uint64, quota subscriptiontypes.Quota) {
	k.subscription.SetQuota(ctx, id, quota)
}

func (k *Keeper) GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscriptiontypes.Quota, bool) {
	return k.subscription.GetQuota(ctx, id, address)
}

func (k *Keeper) HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool {
	return k.subscription.HasQuota(ctx, id, address)
}

func (k *Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, to, sdk.NewCoins(coin))
}
