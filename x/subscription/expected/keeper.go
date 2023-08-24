// DO NOT COVER

package expected

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type ProviderKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
}

type NodeKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
	GetNode(ctx sdk.Context, addr hubtypes.NodeAddress) (nodetypes.Node, bool)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SessionKeeper interface {
	GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool)
	SubscriptionInactivePendingHook(ctx sdk.Context, id uint64) error
}
