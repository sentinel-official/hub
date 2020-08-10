package simulation

import (
	"math/rand"

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
