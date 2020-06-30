package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
)

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hub.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
}
