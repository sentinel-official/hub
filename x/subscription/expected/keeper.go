package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
)

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) sdk.Error
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plan.Plan, bool)
	GetPlans(ctx sdk.Context, skip, limit int) plan.Plans
}
