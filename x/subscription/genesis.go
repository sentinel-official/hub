package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	inactiveDuration := k.InactiveDuration(ctx)
	for _, item := range state.Subscriptions {
		k.SetSubscription(ctx, item.Subscription)

		for _, quota := range item.Quotas {
			address := quota.GetAddress()
			k.SetQuota(ctx, item.Subscription.Id, quota)

			if item.Subscription.Status.Equal(hubtypes.StatusActive) {
				k.SetActiveSubscriptionForAddress(ctx, address, item.Subscription.Id)
			} else {
				k.SetInactiveSubscriptionForAddress(ctx, address, item.Subscription.Id)
			}
		}

		if item.Subscription.Status.Equal(hubtypes.StatusActive) {
			if item.Subscription.Plan != 0 {
				k.SetInactiveSubscriptionAt(ctx, item.Subscription.Expiry, item.Subscription.Id)
			}
		}
		if item.Subscription.Status.Equal(hubtypes.StatusInactivePending) {
			k.SetInactiveSubscriptionAt(ctx, item.Subscription.StatusAt.Add(inactiveDuration), item.Subscription.Id)
		}
	}

	count := uint64(0)
	for _, item := range state.Subscriptions {
		if item.Subscription.Id > count {
			count = item.Subscription.Id
		}
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var (
		subscriptions = k.GetSubscriptions(ctx)
		items         = make(types.GenesisSubscriptions, 0, len(subscriptions))
	)

	for _, item := range subscriptions {
		items = append(
			items,
			types.GenesisSubscription{
				Subscription: item,
				Quotas:       k.GetQuotas(ctx, item.Id, 0, 0),
			},
		)
	}

	return types.NewGenesisState(items, k.GetParams(ctx))
}
