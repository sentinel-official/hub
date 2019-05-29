package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func NewQuerier(k keeper.Keeper, cdc *codec.Codec) csdkTypes.Querier {
	return func(ctx csdkTypes.Context, path []string, req abciTypes.RequestQuery) (res []byte, err csdkTypes.Error) {
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
