package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	provider.InitGenesis(ctx, k.Provider, state.Providers)
	node.InitGenesis(ctx, k.Node, state.Nodes)
	subscription.InitGenesis(ctx, k.Subscription, state.Subscriptions)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Providers:     provider.ExportGenesis(ctx, k.Provider),
		Nodes:         node.ExportGenesis(ctx, k.Node),
		Subscriptions: subscription.ExportGenesis(ctx, k.Subscription),
	}
}

func ValidateGenesis(state types.GenesisState) error {
	if err := provider.ValidateGenesis(state.Providers); err != nil {
		return err
	}
	if err := node.ValidateGenesis(state.Nodes); err != nil {
		return err
	}
	if err := subscription.ValidateGenesis(state.Subscriptions); err != nil {
		return err
	}

	return nil
}
