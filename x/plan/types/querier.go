package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

func NewQueryPlanRequest(id uint64) *QueryPlanRequest {
	return &QueryPlanRequest{
		Id: id,
	}
}

func NewQueryPlansRequest(status hub.Status, pagination *query.PageRequest) *QueryPlansRequest {
	return &QueryPlansRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryPlansForProviderRequest(address hub.ProvAddress, status hub.Status, pagination *query.PageRequest) *QueryPlansForProviderRequest {
	return &QueryPlansForProviderRequest{
		Address:    address.String(),
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForPlanRequest(id uint64, pagination *query.PageRequest) *QueryNodesForPlanRequest {
	return &QueryNodesForPlanRequest{
		Id:         id,
		Pagination: pagination,
	}
}
