package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func RandomSubscription(r *rand.Rand, subscriptions types.Subscriptions) types.Subscription {
	if len(subscriptions) == 0 {
		return types.Subscription{
			ID: 1,
		}
	}

	return subscriptions[r.Intn(
		len(subscriptions),
	)]
}

func RandomQuota(r *rand.Rand, quotas types.Quotas) types.Quota {
	if len(quotas) == 0 {
		return types.Quota{
			Address:   sdk.AccAddress("address"),
			Consumed:  hub.NewBandwidthFromInt64(0, 0),
			Allocated: hub.NewBandwidthFromInt64(1, 1),
		}
	}

	return quotas[r.Intn(
		len(quotas),
	)]
}
