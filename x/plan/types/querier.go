package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQueryPlanRequest(id uint64) *QueryPlanRequest {
	return &QueryPlanRequest{
		Id: id,
	}
}

func NewQueryPlansRequest(status hubtypes.Status, pagination *query.PageRequest) *QueryPlansRequest {
	return &QueryPlansRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryPlansForProviderRequest(address hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) *QueryPlansForProviderRequest {
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
