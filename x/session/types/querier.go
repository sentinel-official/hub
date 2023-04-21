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

func NewQuerySessionsForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QuerySessionsForAccountRequest {
	return &QuerySessionsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForNodeRequest(addr hubtypes.NodeAddress, pagination *query.PageRequest) *QuerySessionsForNodeRequest {
	return &QuerySessionsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySessionsForSubscriptionRequest(id uint64, pagination *query.PageRequest) *QuerySessionsForSubscriptionRequest {
	return &QuerySessionsForSubscriptionRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
