package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
)

func (k *Keeper) SendCoin(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.SendCoinsFromAccountToModule(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Add(ctx, address, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Subtract(ctx, address, sdk.NewCoins(coin))
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool) {
	return k.plan.GetPlan(ctx, id)
}
