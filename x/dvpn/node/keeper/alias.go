package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider"
)

func (k Keeper) GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool) {
	return k.provider.GetProvider(ctx, address)
}
