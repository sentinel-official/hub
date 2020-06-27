package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	deposit.InitGenesis(ctx, k.Deposit, state.Deposits)
	provider.InitGenesis(ctx, k.Provider, state.Providers)
	node.InitGenesis(ctx, k.Node, state.Nodes)
	subscription.InitGenesis(ctx, k.Subscription, state.Subscription)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Deposits:     deposit.ExportGenesis(ctx, k.Deposit),
		Providers:    provider.ExportGenesis(ctx, k.Provider),
		Nodes:        node.ExportGenesis(ctx, k.Node),
		Subscription: subscription.ExportGenesis(ctx, k.Subscription),
	}
}

func ValidateGenesis(state types.GenesisState) error {
	if err := deposit.ValidateGenesis(state.Deposits); err != nil {
		return err
	}
	if err := provider.ValidateGenesis(state.Providers); err != nil {
		return err
	}
	if err := node.ValidateGenesis(state.Nodes); err != nil {
		return err
	}
	if err := subscription.ValidateGenesis(state.Subscription); err != nil {
		return err
	}

	return nil
}
