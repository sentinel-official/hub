package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryPlan:
		return queryPlan(ctx, req, k)
	case types.QueryPlans:
		return queryPlans(ctx, req, k)
	case types.QueryPlansOfProvider:
		return queryPlansOfProvider(ctx, req, k)
	case types.QueryNodesOfPlan:
		return queryNodesOfPlan(ctx, req, k)
	case types.QuerySubscription:
		return querySubscription(ctx, req, k)
	case types.QuerySubscriptions:
		return querySubscriptions(ctx, req, k)
	case types.QuerySubscriptionsOfAddress:
		return querySubscriptionsOfAddress(ctx, req, k)
	case types.QuerySubscriptionsOfPlan:
		return querySubscriptionsOfPlan(ctx, req, k)
	case types.QuerySubscriptionsOfNode:
		return querySubscriptionsOfNode(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
