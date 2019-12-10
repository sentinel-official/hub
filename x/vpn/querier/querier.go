package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryNode:
			return queryNode(ctx, req, k)
		case types.QueryNodesOfAddress:
			return queryNodesOfAddress(ctx, req, k)
		case types.QueryAllNodes:
			return queryAllNodes(ctx, k)
		case types.QueryFreeNodesOfClient:
			return queryFreeNodesOfClient(ctx, req, k)
		case types.QueryFreeClientsOfNode:
			return queryFreeClientsOfNode(ctx, req, k)
		case types.QueryNodesOfResolver:
			return queryNodesOfResolver(ctx, req, k)
		case types.QueryResolversOfNode:
			return queryResolversOfNode(ctx, req, k)
		case types.QuerySubscription:
			return querySubscription(ctx, req, k)
		case types.QuerySubscriptionsOfNode:
			return querySubscriptionsOfNode(ctx, req, k)
		case types.QuerySubscriptionsOfAddress:
			return querySubscriptionsOfAddress(ctx, req, k)
		case types.QueryAllSubscriptions:
			return queryAllSubscriptions(ctx, k)
		case types.QuerySessionsCountOfSubscription:
			return querySessionsCountOfSubscription(ctx, req, k)
		case types.QuerySession:
			return querySession(ctx, req, k)
		case types.QuerySessionOfSubscription:
			return querySessionOfSubscription(ctx, req, k)
		case types.QuerySessionsOfSubscription:
			return querySessionsOfSubscription(ctx, req, k)
		case types.QueryAllSessions:
			return queryAllSessions(ctx, k)
		case types.QueryResolvers:
			return queryResolvers(ctx, req, k)

		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}
