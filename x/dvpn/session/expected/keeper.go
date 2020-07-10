package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

type PlanKeeper interface {
	HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool
}

type SubscriptionKeeper interface {
	SetSubscription(ctx sdk.Context, subscription subscription.Subscription)
	GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool)

	HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool
	HasSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) bool
}
