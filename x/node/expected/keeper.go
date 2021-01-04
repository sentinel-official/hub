package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type AccountKeeper interface {
	IterateAccounts(ctx sdk.Context, process func(exported.Account) (stop bool))
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
	GetProviders(ctx sdk.Context, skip, limit int) provider.Providers
}

type PlanKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) plan.Plans
	DeleteNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress)
}
