package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	provider "github.com/sentinel-official/hub/x/provider/types"
)

type ProviderKeeper interface {
	GetProviders(ctx sdk.Context, skip, limit int) provider.Providers
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int) node.Nodes
}
