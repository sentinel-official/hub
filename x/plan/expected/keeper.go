package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type ProviderKeeper interface {
	GetProviders(ctx sdk.Context, skip, limit int64) providertypes.Providers
	HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool)
	GetNodes(ctx sdk.Context, skip, limit int64) nodetypes.Nodes
	GetNodesForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int64) nodetypes.Nodes
}
