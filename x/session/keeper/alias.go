package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	nodetypes "github.com/sentinel-official/hub/x/node/types"

	hubtypes "github.com/sentinel-official/hub/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, addr)
}

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) bool {
	return k.node.HasNodeForPlan(ctx, id, addr)
}

func (k *Keeper) GetNode(ctx sdk.Context, addr hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, addr)
}

func (k *Keeper) GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (subscriptiontypes.Allocation, bool) {
	return k.subscription.GetAllocation(ctx, id, addr)
}

func (k *Keeper) SetAllocation(ctx sdk.Context, id uint64, allocation subscriptiontypes.Allocation) {
	k.subscription.SetAllocation(ctx, id, allocation)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscriptiontypes.Subscription, bool) {
	return k.subscription.GetSubscription(ctx, id)
}
