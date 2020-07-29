package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func QueryDeposit(ctx context.CLIContext, address sdk.AccAddress) (*types.Deposit, error) {
	params := types.NewQueryDepositParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryDeposit)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no deposit found")
	}

	var deposit types.Deposit
	if err = ctx.Codec.UnmarshalJSON(res, &deposit); err != nil {
		return nil, err
	}

	return &deposit, nil
}

func QueryDeposits(ctx context.CLIContext) (types.Deposits, error) {
	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryDeposits)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no deposits found")
	}

	var deposits types.Deposits
	if err = ctx.Codec.UnmarshalJSON(res, &deposits); err != nil {
		return nil, err
	}

	return deposits, nil
}
