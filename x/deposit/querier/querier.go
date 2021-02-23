package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func Querier(ctx sdk.Context, path []string, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	switch path[0] {
	case types.QueryDeposit:
		return queryDeposit(ctx, req, k)
	case types.QueryDeposits:
		return queryDeposits(ctx, req, k)
	default:
		return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
	}
}
