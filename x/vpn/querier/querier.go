package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

// nolint:gocyclo
func NewQuerier(k keeper.Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryNode:
			return queryNode(ctx, cdc, req, k)
		case QueryNodesOfAddress:
			return queryNodesOfAddress(ctx, cdc, req, k)
		case QueryAllNodes:
			return queryAllNodes(ctx, cdc, k)
		case QuerySubscription:
			return querySubscription(ctx, cdc, req, k)
		case QuerySubscriptionsOfNode:
			return querySubscriptionsOfNode(ctx, cdc, req, k)
		case QuerySubscriptionsOfAddress:
			return querySubscriptionsOfAddress(ctx, cdc, req, k)
		case QueryAllSubscriptions:
			return queryAllSubscriptions(ctx, cdc, k)
		case QuerySessionsCountOfSubscription:
			return querySessionsCountOfSubscription(ctx, cdc, req, k)
		case QuerySession:
			return querySession(ctx, cdc, req, k)
		case QuerySessionOfSubscription:
			return querySessionOfSubscription(ctx, cdc, req, k)
		case QuerySessionsOfSubscription:
			return querySessionsOfSubscription(ctx, cdc, req, k)
		case QueryAllSessions:
			return queryAllSessions(ctx, cdc, k)
		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}
