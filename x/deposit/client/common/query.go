// nolint:dupl
package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
)

func QueryDepositOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, _address string) (*deposit.Deposit, error) {
	address, err := csdk.AccAddressFromBech32(_address)
	if err != nil {
		return nil, err
	}

	params := deposit.NewQueryDepositOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", deposit.QuerierRoute, deposit.QueryDepositOfAddress), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no deposit found")
	}

	var _deposit deposit.Deposit
	if err = cdc.UnmarshalJSON(res, &_deposit); err != nil {
		return nil, err
	}

	return &_deposit, nil
}

func QueryAllDeposits(cliCtx context.CLIContext, cdc *codec.Codec) ([]deposit.Deposit, error) {
	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", deposit.QuerierRoute, deposit.QueryAllDeposits), nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no deposits found")
	}

	var deposits []deposit.Deposit
	if err = cdc.UnmarshalJSON(res, &deposits); err != nil {
		return nil, err
	}

	return deposits, nil
}
