package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QueryDeposit  = "deposit"
	QueryDeposits = "deposits"
)

type QueryDepositParams struct {
	Address sdk.AccAddress `json:"address"`
}

func NewQueryDepositParams(address sdk.AccAddress) QueryDepositParams {
	return QueryDepositParams{
		Address: address,
	}
}

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
