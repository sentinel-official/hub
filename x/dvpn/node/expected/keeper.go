package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/dvpn/plan/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

type ProviderKeeper interface {
	GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool)
}

type PlanKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) plan.Plans
	DeleteNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress)
}
