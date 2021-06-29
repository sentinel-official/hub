package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, address)
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, address)
}
