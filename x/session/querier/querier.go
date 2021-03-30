package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	switch path[0] {
	case types.QuerySession:
		return querySession(ctx, req, k)
	case types.QuerySessions:
		return querySessions(ctx, req, k)
	case types.QuerySessionsForSubscription:
		return querySessionsForSubscription(ctx, req, k)
	case types.QuerySessionsForNode:
		return querySessionsForNode(ctx, req, k)
	case types.QuerySessionsForAddress:
		return querySessionsForAddress(ctx, req, k)

	case types.QueryActiveSession:
		return queryActiveSession(ctx, req, k)
	default:
		return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
	}
}
