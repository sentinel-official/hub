package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, address)
}

func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	k.node.SetNodeForPlan(ctx, id, addr)
}

func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	k.node.DeleteNodeForPlan(ctx, id, addr)
}

func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) nodetypes.Nodes {
	return k.node.GetNodesForPlan(ctx, id)
}
