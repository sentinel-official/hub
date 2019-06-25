package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func NewQuerier(k keeper.Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryDepositOfAddress:
			return queryDepositOfAddress(ctx, cdc, req, k)
		case QueryAllDeposits:
			return queryAllDeposits(ctx, cdc, k)
		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}
