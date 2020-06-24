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
