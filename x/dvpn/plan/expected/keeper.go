package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

type ProviderKeeper interface {
	GetProviders(ctx sdk.Context) provider.Providers
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
	GetNodes(ctx sdk.Context) node.Nodes
}
