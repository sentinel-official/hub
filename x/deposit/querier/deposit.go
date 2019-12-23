package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func queryDepositOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryDepositOfAddressPrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}
	
	deposit, found := k.GetDeposit(ctx, params.Address)
	if !found {
		return nil, nil
	}
	
	res, err := types.ModuleCdc.MarshalJSON(deposit)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	
	return res, nil
}

func queryAllDeposits(ctx sdk.Context, k keeper.Keeper) ([]byte, sdk.Error) {
	deposits := k.GetAllDeposits(ctx)
	
	res, err := types.ModuleCdc.MarshalJSON(deposits)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	
	return res, nil
}
