package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/querier"
)

func QueryDepositsOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]byte, error) {
	params := querier.NewQueryDepositsOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", deposit.QuerierRoute, deposit.QueryDepositsOfAddress), paramBytes)
}

func QueryAllDeposits(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", deposit.QuerierRoute, deposit.QueryAllDeposits), nil)
}
