package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

const (
	QueryDeposit  = "deposit"
	QueryDeposits = "deposits"
)

func NewQueryDepositRequest(address sdk.AccAddress) *QueryDepositRequest {
	return &QueryDepositRequest{
		Address: address.String(),
	}
}

func NewQueryDepositsRequest(pagination *query.PageRequest) *QueryDepositsRequest {
	return &QueryDepositsRequest{
		Pagination: pagination,
	}
}
