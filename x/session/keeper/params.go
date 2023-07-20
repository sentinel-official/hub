package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) ExpiryDuration(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, types.KeyExpiryDuration, &duration)
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
		k.ExpiryDuration(ctx),
		k.ProofVerificationEnabled(ctx),
	)
}
