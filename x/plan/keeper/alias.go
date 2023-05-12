package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) HasNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool {
	return k.node.HasNode(ctx, addr)
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

func (k *Keeper) CreateSubscriptionForPlan(
	ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string,
) (*subscriptiontypes.PlanSubscription, error) {
	return k.subscription.CreateSubscriptionForPlan(ctx, accAddr, id, denom)
}
