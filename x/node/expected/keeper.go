package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
	GetProviders(ctx sdk.Context, skip, limit int) provider.Providers
}

type PlanKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) plan.Plans
	DeleteNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress)
}
