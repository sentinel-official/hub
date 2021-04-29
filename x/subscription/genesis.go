package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, item := range state.Subscriptions {
		k.SetSubscription(ctx, item.Subscription)

		if item.Subscription.Id == 0 {
			k.SetSubscriptionForNode(ctx, item.Subscription.GetNode(), item.Subscription.Id)
		} else {
			k.SetSubscriptionForPlan(ctx, item.Subscription.Plan, item.Subscription.Id)
		}

		for _, quota := range item.Quotas {
			k.SetQuota(ctx, item.Subscription.Id, quota)

			if item.Subscription.Status.Equal(hubtypes.StatusInactive) {
				k.SetInactiveSubscriptionForAddress(ctx, quota.GetAddress(), item.Subscription.Id)
			} else {
				k.SetActiveSubscriptionForAddress(ctx, quota.GetAddress(), item.Subscription.Id)
			}
		}

		if item.Subscription.Status.Equal(hubtypes.StatusInactivePending) {
			k.SetInactiveSubscriptionAt(ctx, item.Subscription.StatusAt.Add(k.InactiveDuration(ctx)), item.Subscription.Id)
		}
	}

	k.SetCount(ctx, uint64(len(state.Subscriptions)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	subscriptions := k.GetSubscriptions(ctx, 0, 0)

	items := make(types.GenesisSubscriptions, 0, len(subscriptions))
	for _, item := range subscriptions {
		items = append(items, types.GenesisSubscription{
			Subscription: item,
			Quotas:       k.GetQuotas(ctx, item.Id, 0, 0),
		})
	}

	return types.NewGenesisState(items, k.GetParams(ctx))
}
