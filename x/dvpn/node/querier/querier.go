package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryNode:
		return queryNode(ctx, req, k)
	case types.QueryNodes:
		return queryNodes(ctx, req, k)
	case types.QueryNodesOfProvider:
		return queryNodesOfProvider(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
