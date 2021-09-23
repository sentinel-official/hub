package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func (k *Keeper) GetParams(ctx sdk.Context) minttypes.Params {
	return k.mint.GetParams(ctx)
}

func (k *Keeper) SetParams(ctx sdk.Context, params minttypes.Params) {
	k.mint.SetParams(ctx, params)
}
