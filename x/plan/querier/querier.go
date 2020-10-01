package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	switch path[0] {
	case types.QueryPlan:
		return queryPlan(ctx, req, k)
	case types.QueryPlans:
		return queryPlans(ctx, req, k)
	case types.QueryPlansForProvider:
		return queryPlansForProvider(ctx, req, k)

	case types.QueryNodesForPlan:
		return queryNodesForPlan(ctx, req, k)
	default:
		return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
	}
}
