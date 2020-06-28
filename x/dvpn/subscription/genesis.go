package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
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
	_subscriptions := k.GetSubscriptions(ctx)

	subscriptions := make(types.GenesisSubscriptions, 0, len(_subscriptions))
	for _, item := range _subscriptions {
		subscriptions = append(subscriptions, types.GenesisSubscription{
			Subscription: item,
			Members:      k.GetAddressesForSubscriptionID(ctx, item.ID),
		})
	}

	return types.NewGenesisState(subscriptions)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
