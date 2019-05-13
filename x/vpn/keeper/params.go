package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	DefaultParamspace = types.ModuleName
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&types.Params{})
}

func (k Keeper) FreeNodesCount(ctx csdkTypes.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyFreeNodesCount, &res)
	return
}

func (k Keeper) FreeSessionsCount(ctx csdkTypes.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyFreeSessionsCount, &res)
	return
}

func (k Keeper) Deposit(ctx csdkTypes.Context) (res csdkTypes.Coin) {
	k.paramStore.Get(ctx, types.KeyDeposit, &res)
	return
}

func (k Keeper) FreeSessionBandwidth(ctx csdkTypes.Context) (res sdkTypes.Bandwidth) {
	k.paramStore.Get(ctx, types.KeyFreeSessionBandwidth, &res)
	return
}

func (k Keeper) NodeInactiveInterval(ctx csdkTypes.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeyNodeInactiveInterval, &res)
	return
}

func (k Keeper) SessionEndInterval(ctx csdkTypes.Context) (res int64) {
	k.paramStore.Get(ctx, types.KeySessionEndInterval, &res)
	return
}

func (k Keeper) GetParams(ctx csdkTypes.Context) types.Params {
	return types.NewParams(
		k.FreeNodesCount(ctx),
		k.FreeSessionsCount(ctx),
		k.Deposit(ctx),
		k.FreeSessionBandwidth(ctx),
		k.NodeInactiveInterval(ctx),
		k.SessionEndInterval(ctx),
	)
}

func (k Keeper) SetParams(ctx csdkTypes.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}
