package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) SetActiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForPlanKey(id, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasActiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForPlanKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteActiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForPlanKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForPlanKey(id, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForPlanKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteInactiveNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForPlanKey(id, addr)
	)

	store.Delete(key)
}
