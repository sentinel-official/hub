package subscription

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, item := range state.Subscriptions {
		k.SetSubscription(ctx, item.Subscription)

		for _, quota := range item.Quotas {
			k.SetQuota(ctx, item.Subscription.ID, quota)
			k.SetSubscriptionForAddress(ctx, quota.Address, item.Subscription.ID)
		}

		if item.Subscription.ID == 0 {
			k.SetSubscriptionForNode(ctx, item.Subscription.Node, item.Subscription.ID)
		} else {
			k.SetSubscriptionForPlan(ctx, item.Subscription.Plan, item.Subscription.ID)
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
			Quotas:       k.GetQuotas(ctx, item.ID),
		})
	}

	return types.NewGenesisState(subscriptions, k.GetParams(ctx))
}

func ValidateGenesis(state types.GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, item := range state.Subscriptions {
		if err := item.Subscription.Validate(); err != nil {
			return err
		}
	}

	subscriptions := make(map[uint64]bool)
	for _, item := range state.Subscriptions {
		id := item.Subscription.ID
		if subscriptions[id] {
			return fmt.Errorf("duplicate subscription id %d", id)
		}

		subscriptions[id] = true
	}

	for _, item := range state.Subscriptions {
		quotas := make(map[string]bool)
		for _, quota := range item.Quotas {
			address := quota.Address.String()
			if quotas[address] {
				return fmt.Errorf("duplicate quota for subscription %d", item.Subscription.ID)
			}

			quotas[address] = true
		}
	}

	return nil
}
