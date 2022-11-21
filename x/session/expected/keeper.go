package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, name string, coins sdk.Coins) error
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error
}

type PlanKeeper interface {
	HasNodeForPlan(ctx sdk.Context, id uint64, address hubtypes.NodeAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool)
	StakingShare(ctx sdk.Context) sdk.Dec
}

type SubscriptionKeeper interface {
	GetSubscription(ctx sdk.Context, id uint64) (subscriptiontypes.Subscription, bool)
	GetActiveSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int64) subscriptiontypes.Subscriptions

	SetQuota(ctx sdk.Context, id uint64, quota subscriptiontypes.Quota)
	HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool
	GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscriptiontypes.Quota, bool)
}
