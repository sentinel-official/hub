package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) Deposit(ctx sdk.Context) (deposit sdk.Coin) {
	k.params.Get(ctx, types.KeyDeposit, &deposit)
	return
}

func (k *Keeper) InactiveDuration(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, types.KeyInactiveDuration, &duration)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Deposit(ctx),
		k.InactiveDuration(ctx),
	)
}
