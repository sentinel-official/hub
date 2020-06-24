package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryKeyPlan:
		return queryPlan(ctx, req, k)
	case types.QueryKeyPlans:
		return queryPlans(ctx, req, k)
	case types.QueryKeyPlansOfProvider:
		return queryPlansOfProvider(ctx, req, k)
	case types.QueryKeyNodesOfPlan:
		return queryNodesOfPlan(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
