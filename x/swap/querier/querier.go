package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QuerySwap:
			return querySwap(ctx, req, k)
		case types.QuerySwaps:
			return querySwaps(ctx, req, k)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
		}
	}
}
