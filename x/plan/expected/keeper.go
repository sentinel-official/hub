package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) auth.AccountI
}

type ProviderKeeper interface {
	GetProviders(ctx sdk.Context, skip, limit int) provider.Providers
	HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
	GetNodesForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int) node.Nodes
}
