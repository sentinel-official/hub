package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, address)
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (node.Node, bool) {
	return k.node.GetNode(ctx, address)
}
