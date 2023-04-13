package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) SetActiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForProviderKey(provAddr, nodeAddr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasActiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForProviderKey(provAddr, nodeAddr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteActiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeForProviderKey(provAddr, nodeAddr)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForProviderKey(provAddr, nodeAddr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForProviderKey(provAddr, nodeAddr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteInactiveNodeForProvider(ctx sdk.Context, provAddr hubtypes.ProvAddress, nodeAddr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeForProviderKey(provAddr, nodeAddr)
	)

	store.Delete(key)
}
