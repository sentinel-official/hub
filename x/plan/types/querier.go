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

func NewQueryPlansForProviderRequest(addr hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) *QueryPlansForProviderRequest {
	return &QueryPlansForProviderRequest{
		Address:    addr.String(),
		Status:     status,
		Pagination: pagination,
	}
}
