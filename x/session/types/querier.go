package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQuerySessionRequest(id uint64) *QuerySessionRequest {
	return &QuerySessionRequest{
		Id: id,
	}
}

func NewQuerySessionsRequest(pagination *query.PageRequest) *QuerySessionsRequest {
	return &QuerySessionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySessionsForAddressRequest(address sdk.AccAddress, status hubtypes.Status, pagination *query.PageRequest) *QuerySessionsForAddressRequest {
	return &QuerySessionsForAddressRequest{
		Address:    address.String(),
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
