// DO NOT COVER

package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool
}

type NodeKeeper interface {
	HasNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool
	SetNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress)
	DeleteNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress)
	GetNodesForPlan(ctx sdk.Context, id uint64) nodetypes.Nodes
}
