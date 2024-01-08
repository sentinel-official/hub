package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/types"
)

func (k *Keeper) StatusChangeDelay(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, types.KeyStatusChangeDelay, &duration)
	return
}

func (k *Keeper) ProofVerificationEnabled(ctx sdk.Context) (yes bool) {
	k.params.Get(ctx, types.KeyProofVerificationEnabled, &yes)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.StatusChangeDelay(ctx),
		k.ProofVerificationEnabled(ctx),
	)
}
