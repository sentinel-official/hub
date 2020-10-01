package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
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

	case types.QueryQuota:
		return queryQuota(ctx, req, k)
	case types.QueryQuotas:
		return queryQuotas(ctx, req, k)
	default:
		return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
	}
}
