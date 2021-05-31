package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func getRandomSubscriptionIDForPlan(r *rand.Rand, subscriptions types.Subscriptions) uint64 {
	if len(subscriptions) == 0 {
		return uint64(1)
	}

	return subscriptions[r.Intn(len(subscriptions))].Id
}

func getRandomSubscription(r *rand.Rand, subscriptions types.Subscriptions, kind string) types.Subscription {
	if len(subscriptions) == 0 {
		return types.Subscription{Id: 1}
	}

	if kind == "node" {
		for i, s := range subscriptions {
			if s.Plan == 0 {
				subscriptions[i] = subscriptions[len(subscriptions)-1]
				subscriptions = subscriptions[:len(subscriptions)-1]
			}
		}
		return subscriptions[r.Intn(len(subscriptions))]
	}

	for i, s := range subscriptions {
		if s.Plan != 0 {
			subscriptions[i] = subscriptions[len(subscriptions)-1]
			subscriptions = subscriptions[:len(subscriptions)-1]
		}
	}
	return subscriptions[r.Intn(len(subscriptions))]
}

func getRandomBytes(r *rand.Rand) sdk.Int {
	gb := int64(1024 << 20)
	randomFactor := r.Int63n(16) + 2
	return sdk.NewInt(gb * randomFactor)
}
