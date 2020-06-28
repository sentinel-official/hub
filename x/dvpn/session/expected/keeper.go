package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

type SubscriptionKeeper interface {
	SetSubscription(ctx sdk.Context, subscription subscription.Subscription)
	GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool)

	HasNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool
	HasSubscriptionIDForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool
	HasAddressForSubscriptionID(ctx sdk.Context, id uint64, address sdk.AccAddress) bool
}
