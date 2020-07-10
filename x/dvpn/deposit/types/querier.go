package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QueryDeposit  = "deposit"
	QueryDeposits = "deposits"
)

type QueryDepositParams struct {
	Address sdk.AccAddress
}

func NewQueryDepositParams(address sdk.AccAddress) QueryDepositParams {
	return QueryDepositParams{
		Address: address,
	}
}
