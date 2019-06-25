package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/sentinel-hub/x/deposit/keeper"
	"github.com/sentinel-official/sentinel-hub/x/deposit/types"
)

const (
	QueryDepositOfAddress = "depositOfAddress"
	QueryAllDeposits      = "allDeposits"
)

type QueryDepositOfAddressPrams struct {
	Address sdk.AccAddress
}

func NewQueryDepositOfAddressParams(address sdk.AccAddress) QueryDepositOfAddressPrams {
	return QueryDepositOfAddressPrams{
		Address: address,
	}
}

// nolint:dupl
func queryDepositOfAddress(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

	var params QueryDepositOfAddressPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	deposit, found := k.GetDeposit(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(deposit)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllDeposits(ctx sdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, sdk.Error) {
	deposits := k.GetAllDeposits(ctx)

	res, resErr := cdc.MarshalJSON(deposits)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
