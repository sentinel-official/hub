// DO NOT COVER

package expected

import (
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
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, name string, coins sdk.Coins) error
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
}

type ProviderKeeper interface {
	RevenueShare(ctx sdk.Context) sdk.Dec
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, addr hubtypes.NodeAddress) (nodetypes.Node, bool)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SessionKeeper interface {
	GetSessionsForAccount(ctx sdk.Context, addr sdk.AccAddress) sessiontypes.Sessions
}
