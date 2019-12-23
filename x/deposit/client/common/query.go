package common

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/sentinel-official/hub/x/deposit/types"
)

func QueryDepositOfAddress(ctx context.CLIContext, s string) (*types.Deposit, error) {
	address, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryDepositOfAddressParams(address)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositOfAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no deposit found")
	}
	
	var d types.Deposit
	if err = ctx.Codec.UnmarshalJSON(res, &d); err != nil {
		return nil, err
	}
	
	return &d, nil
}

func QueryAllDeposits(ctx context.CLIContext) ([]types.Deposit, error) {
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllDeposits)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no deposits found")
	}
	
	var d []types.Deposit
	if err = ctx.Codec.UnmarshalJSON(res, &d); err != nil {
		return nil, err
	}
	
	return d, nil
}
