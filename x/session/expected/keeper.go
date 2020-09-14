package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	subscription "github.com/sentinel-official/hub/x/subscription/types"
)

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error
}

type PlanKeeper interface {
	HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool
}

type SubscriptionKeeper interface {
	GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool)

	HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool

	SetQuota(ctx sdk.Context, id uint64, quota subscription.Quota)
	GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscription.Quota, bool)
	HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool
}
