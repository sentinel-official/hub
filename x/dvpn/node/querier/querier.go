package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryKeyNode:
		return queryNode(ctx, req, k)
	case types.QueryKeyNodes:
		return queryNodes(ctx, req, k)
	case types.QueryKeyNodesOfProvider:
		return queryNodesOfProvider(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
