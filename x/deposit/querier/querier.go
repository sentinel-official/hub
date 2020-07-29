package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	switch path[0] {
	case types.QueryDeposit:
		return queryDeposit(ctx, req, k)
	case types.QueryDeposits:
		return queryDeposits(ctx, req, k)
	default:
		return nil, types.ErrorUnknownQueryType(path[0])
	}
}
