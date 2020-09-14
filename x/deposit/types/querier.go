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
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryDepositsParams(page, limit int) QueryDepositsParams {
	return QueryDepositsParams{
		Page:  page,
		Limit: limit,
	}
}
