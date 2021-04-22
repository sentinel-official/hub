package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) auth.AccountI
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool
	GetProviders(ctx sdk.Context, skip, limit int64) provider.Providers
}

type PlanKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int64) plan.Plans
	DeleteNodeForPlan(ctx sdk.Context, id uint64, address hubtypes.NodeAddress)
}
