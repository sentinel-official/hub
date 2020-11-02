package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QueryDeposit  = "deposit"
	QueryDeposits = "deposits"
)

// QueryDepositParams is the request parameters for querying a deposit.
type QueryDepositParams struct {
	Address sdk.AccAddress `json:"address"`
}

func NewQueryDepositParams(address sdk.AccAddress) QueryDepositParams {
	return QueryDepositParams{
		Address: address,
	}
}

// QueryDepositsParams is the request parameters for querying the deposits.
type QueryDepositsParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQueryDepositsParams(skip, limit int) QueryDepositsParams {
	return QueryDepositsParams{
		Skip:  skip,
		Limit: limit,
	}
}
