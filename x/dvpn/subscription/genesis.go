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

		for _, node := range item.Nodes {
			k.SetNodeAddressForPlan(ctx, item.Plan.ID, node)
		}

		k.SetPlansCount(ctx, k.GetPlansCount(ctx)+1)
	}

	for _, item := range state.List {
		k.SetSubscription(ctx, item.Subscription)

		for _, member := range item.Members {
			k.SetSubscriptionIDForAddress(ctx, member, item.Subscription.ID)
			k.SetAddressForSubscriptionID(ctx, item.Subscription.ID, member)
		}

		if item.Subscription.ID > 0 {
			k.SetSubscriptionIDForPlan(ctx, item.Subscription.Plan, item.Subscription.ID)
		} else {
			k.SetSubscriptionIDForNode(ctx, item.Subscription.Node, item.Subscription.ID)
		}

		k.SetSubscriptionsCount(ctx, k.GetSubscriptionsCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_plans := k.GetPlans(ctx)
	_subscriptions := k.GetSubscriptions(ctx)

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

	subscriptions := make(types.GenesisSubscriptions, 0, len(_subscriptions))
	for _, item := range _subscriptions {
		subscriptions = append(subscriptions, types.GenesisSubscription{
			Subscription: item,
			Members:      k.GetAddressesForSubscriptionID(ctx, item.ID),
		})
	}

	return types.NewGenesisState(plans, subscriptions)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
