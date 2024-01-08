package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types"
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

func (k *Keeper) AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Add(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Subtract(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, fromAddr, toAddr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, fromAddr, toModule, sdk.NewCoins(coin))
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool) {
	return k.plan.GetPlan(ctx, id)
}

func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool) {
	return k.session.GetSession(ctx, id)
}

func (k *Keeper) SubscriptionInactivePendingHook(ctx sdk.Context, id uint64) error {
	return k.session.SubscriptionInactivePendingHook(ctx, id)
}
