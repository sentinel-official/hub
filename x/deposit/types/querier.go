package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func NewQueryDepositRequest(addr sdk.AccAddress) *QueryDepositRequest {
	return &QueryDepositRequest{
		Address: addr.String(),
	}
}

func NewQueryDepositsRequest(pagination *query.PageRequest) *QueryDepositsRequest {
	return &QueryDepositsRequest{
		Pagination: pagination,
	}
}
