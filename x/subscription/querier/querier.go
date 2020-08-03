package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QuerySubscription:
		return querySubscription(ctx, req, k)
	case types.QuerySubscriptions:
		return querySubscriptions(ctx, req, k)
	case types.QuerySubscriptionsForAddress:
		return querySubscriptionsForAddress(ctx, req, k)
	case types.QuerySubscriptionsForPlan:
		return querySubscriptionsForPlan(ctx, req, k)
	case types.QuerySubscriptionsForNode:
		return querySubscriptionsForNode(ctx, req, k)
	case types.QueryQuotaForSubscription:
		return queryQuotaForSubscription(ctx, req, k)
	case types.QueryQuotasForSubscription:
		return queryQuotasForSubscription(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
