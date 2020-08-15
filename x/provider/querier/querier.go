package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryProvider:
		return queryProvider(ctx, req, k)
	case types.QueryProviders:
		return queryProviders(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
