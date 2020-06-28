package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func (k Keeper) GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool) {
	return k.provider.GetProvider(ctx, address)
}

func (k Keeper) GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k Keeper) SendCoin(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.bank.SendCoins(ctx, from, to, sdk.NewCoins(coin))
}

func (k Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.Add(ctx, address, sdk.Coins{coin})
}

func (k Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.Subtract(ctx, address, sdk.Coins{coin})
}
