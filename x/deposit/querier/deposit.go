package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func queryDeposit(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryDepositParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	deposit, found := k.GetDeposit(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(deposit)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryDeposits(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryDepositsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	deposits := k.GetDeposits(ctx)

	start, end := client.Paginate(len(deposits), params.Page, params.Limit, len(deposits))
	if start < 0 || end < 0 {
		deposits = types.Deposits{}
	} else {
		deposits = deposits[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(deposits)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
