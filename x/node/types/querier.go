package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQueryNodeRequest(addr hubtypes.NodeAddress) *QueryNodeRequest {
	return &QueryNodeRequest{
		Address: addr.String(),
	}
}

func NewQueryNodesRequest(status hubtypes.Status, pagination *query.PageRequest) *QueryNodesRequest {
	return &QueryNodesRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForPlanRequest(id uint64, status hubtypes.Status, pagination *query.PageRequest) *QueryNodesForPlanRequest {
	return &QueryNodesForPlanRequest{
		Id:         id,
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForProviderRequest(addr hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) *QueryNodesForProviderRequest {
	return &QueryNodesForProviderRequest{
		Address:    addr.String(),
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
