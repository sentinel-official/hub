package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	providerKeeper "github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	deposit.InitGenesis(ctx, k.Deposit, state.Deposits)
	providerKeeper.InitGenesis(ctx, k.Provider, state.Providers)
	node.InitGenesis(ctx, k.Node, state.Nodes)
	plan.InitGenesis(ctx, k.Plan, state.Plans)
	subscription.InitGenesis(ctx, k.Subscription, state.Subscriptions)
	session.InitGenesis(ctx, k.Session, state.Sessions)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Deposits:      deposit.ExportGenesis(ctx, k.Deposit),
		Providers:     providerKeeper.ExportGenesis(ctx, k.Provider),
		Nodes:         node.ExportGenesis(ctx, k.Node),
		Plans:         plan.ExportGenesis(ctx, k.Plan),
		Subscriptions: subscription.ExportGenesis(ctx, k.Subscription),
		Sessions:      session.ExportGenesis(ctx, k.Session),
	}
}

func ValidateGenesis(state types.GenesisState) error {
	if err := deposit.ValidateGenesis(state.Deposits); err != nil {
		return err
	}
	if err := providerKeeper.ValidateGenesis(state.Providers); err != nil {
		return err
	}
	if err := node.ValidateGenesis(state.Nodes); err != nil {
		return err
	}
	if err := plan.ValidateGenesis(state.Plans); err != nil {
		return err
	}
	if err := subscription.ValidateGenesis(state.Subscriptions); err != nil {
		return err
	}
	if err := session.ValidateGenesis(state.Sessions); err != nil {
		return err
	}

	return nil
}
