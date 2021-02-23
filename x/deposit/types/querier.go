package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

const (
	QueryDeposit  = "deposit"
	QueryDeposits = "deposits"
)

func NewQueryDepositRequest(address string) QueryDepositRequest {
	return QueryDepositRequest{
		Address: address,
	}
}

func NewQueryDepositsRequest(pagination *query.PageRequest) QueryDepositsRequest {
	return QueryDepositsRequest{
		Pagination: pagination,
	}
}
