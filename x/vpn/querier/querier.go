package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func NewQuerier(k keeper.Keeper, cdc *codec.Codec) csdkTypes.Querier {
	return func(ctx csdkTypes.Context, path []string, req abciTypes.RequestQuery) (res []byte, err csdkTypes.Error) {
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
		case QuerySession:
			return querySession(ctx, cdc, req, k)
		case QuerySessionsOfSubscription:
			return querySessionsOfSubscription(ctx, cdc, req, k)
		case QueryAllSessions:
			return queryAllSessions(ctx, cdc, k)
		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}
