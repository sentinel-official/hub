package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/dvpn/plan/types"
)

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
}

type PlanKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) plan.Plans
	DeleteNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress)
}
