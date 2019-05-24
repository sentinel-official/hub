package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QueryDepositsOfAddress = "depositsOfAddress"
	QueryAllDeposits       = "allDeposits"
)

type QueryDepositsOfAddressPrams struct {
	Address csdkTypes.AccAddress
}

func NewQueryDepositsOfAddressParams(address csdkTypes.AccAddress) QueryDepositsOfAddressPrams {
	return QueryDepositsOfAddressPrams{
		Address: address,
	}
}

func queryAllDeposits(ctx csdkTypes.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdkTypes.Error) {
	deposits := k.GetAllDeposits(ctx)

	res, resErr := cdc.MarshalJSON(deposits)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryDepositsOfAddress(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

	var params QueryDepositsOfAddressPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	deposits, found := k.GetDeposit(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(deposits)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}
	return res, nil
}
