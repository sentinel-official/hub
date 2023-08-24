// DO NOT COVER

package expected

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
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
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SubscriptionKeeper interface {
	GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (subscriptiontypes.Allocation, bool)
	SetAllocation(ctx sdk.Context, alloc subscriptiontypes.Allocation)
	GetSubscription(ctx sdk.Context, id uint64) (subscriptiontypes.Subscription, bool)
	GetLatestPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress) (subscriptiontypes.Payout, bool)
	SessionInactiveHook(ctx sdk.Context, subscriptionID uint64, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, bytes sdkmath.Int) error
}
