package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/x/swap/types"
)

func (k *Keeper) SwapEnabled(ctx sdk.Context) (yes bool) {
	k.params.Get(ctx, types.KeySwapEnabled, &yes)
	return
}

func (k *Keeper) SwapDenom(ctx sdk.Context) (denom string) {
	k.params.Get(ctx, types.KeySwapDenom, &denom)
	return
}

func (k *Keeper) ApproveBy(ctx sdk.Context) (address string) {
	k.params.Get(ctx, types.KeyApproveBy, &address)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SwapEnabled(ctx),
		k.SwapDenom(ctx),
		k.ApproveBy(ctx),
	)
}
