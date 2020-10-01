package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	switch path[0] {
	case types.QueryNode:
		return queryNode(ctx, req, k)
	case types.QueryNodes:
		return queryNodes(ctx, req, k)
	case types.QueryNodesForProvider:
		return queryNodesForProvider(ctx, req, k)
	default:
		return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
	}
}
