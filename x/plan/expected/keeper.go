package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account
}

type ProviderKeeper interface {
	GetProviders(ctx sdk.Context, skip, limit int) provider.Providers
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
	GetNodesForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) node.Nodes
}
