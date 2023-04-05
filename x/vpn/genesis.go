package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	deposit.InitGenesis(ctx, k.Deposit, state.Deposits)
	node.InitGenesis(ctx, k.Node, state.Nodes)
	plan.InitGenesis(ctx, k.Plan, state.Plans)
	provider.InitGenesis(ctx, k.Provider, state.Providers)
	session.InitGenesis(ctx, k.Session, state.Sessions)
	subscription.InitGenesis(ctx, k.Subscription, state.Subscriptions)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Deposits:      deposit.ExportGenesis(ctx, k.Deposit),
		Providers:     provider.ExportGenesis(ctx, k.Provider),
		Nodes:         node.ExportGenesis(ctx, k.Node),
		Plans:         plan.ExportGenesis(ctx, k.Plan),
		Subscriptions: subscription.ExportGenesis(ctx, k.Subscription),
		Sessions:      session.ExportGenesis(ctx, k.Session),
	}
}
