package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"

	csdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QueryDepositOfAddress = "depositOfAddress"
	QueryAllDeposits      = "allDeposits"
)

type QueryDepositOfAddressPrams struct {
	Address csdk.AccAddress
}

func NewQueryDepositOfAddressParams(address csdk.AccAddress) QueryDepositOfAddressPrams {
	return QueryDepositOfAddressPrams{
		Address: address,
	}
}

// nolint:dupl
func queryDepositOfAddress(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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

func queryAllDeposits(ctx csdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdk.Error) {
	deposits := k.GetAllDeposits(ctx)

	res, resErr := cdc.MarshalJSON(deposits)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
