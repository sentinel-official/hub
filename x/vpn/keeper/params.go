package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	DefaultParamspace = types.ModuleName
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&types.Params{})
}

func (k Keeper) FreeNodesCount(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyFreeNodesCount, &res)
	return
}

func (k Keeper) Deposit(ctx sdk.Context) (res sdk.Coin) {
	k.paramStore.Get(ctx, types.KeyDeposit, &res)
	return
}

func (k Keeper) SessionInactiveInterval(ctx sdk.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeySessionInactiveInterval, &res)
	return
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.FreeNodesCount(ctx),
		k.Deposit(ctx),
		k.SessionInactiveInterval(ctx),
	)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}
