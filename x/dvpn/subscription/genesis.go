package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state.Plans {
		k.SetPlan(ctx, item.Plan)
		k.SetPlanIDForProvider(ctx, item.Plan.Provider, item.Plan.ID)

		for _, address := range item.Nodes {
			k.SetNodeAddressForPlan(ctx, item.Plan.ID, address)
		}

		k.SetPlansCount(ctx, k.GetPlansCount(ctx)+1)
	}

	for _, item := range state.List {
		k.SetSubscription(ctx, item)
		k.SetSubscriptionIDForAddress(ctx, item.Address, item.ID)

		if item.ID > 0 {
			k.SetSubscriptionIDForPlan(ctx, item.Plan, item.ID)
		} else {
			k.SetSubscriptionIDForNode(ctx, item.Node, item.ID)
		}

		k.SetSubscriptionsCount(ctx, k.GetSubscriptionsCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_plans := k.GetPlans(ctx)
	subscriptions := k.GetSubscriptions(ctx)

	plans := make(types.GenesisPlans, 0, len(_plans))
	for _, item := range _plans {
		plan := types.GenesisPlan{
			Plan:  item,
			Nodes: nil,
		}

		nodes := k.GetNodesForPlan(ctx, item.ID)
		for _, node := range nodes {
			plan.Nodes = append(plan.Nodes, node.Address)
		}

		plans = append(plans, plan)
	}

	return types.NewGenesisState(plans, subscriptions)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
