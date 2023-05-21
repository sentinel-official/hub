// DO NOT COVER

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
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool)
	HasNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) bool
	RevenueShare(ctx sdk.Context) sdk.Dec
}

type SubscriptionKeeper interface {
	GetAllocation(ctx sdk.Context, id uint64, address sdk.AccAddress) (subscriptiontypes.Allocation, bool)
	SetAllocation(ctx sdk.Context, id uint64, alloc subscriptiontypes.Allocation)
	GetSubscription(ctx sdk.Context, id uint64) (subscriptiontypes.Subscription, bool)
}
