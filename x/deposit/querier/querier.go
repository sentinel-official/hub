package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

func NewQuerier(k keeper.Keeper, cdc *codec.Codec) csdk.Querier {
	return func(ctx csdk.Context, path []string, req abci.RequestQuery) (res []byte, err csdk.Error) {
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
