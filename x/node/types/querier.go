package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func NewQueryLeaseRequest(id uint64) *QueryLeaseRequest {
	return &QueryLeaseRequest{
		Id: id,
	}
}

func NewQueryLeasesRequest(pagination *query.PageRequest) *QueryLeasesRequest {
	return &QueryLeasesRequest{
		Pagination: pagination,
	}
}

func NewQueryLeasesForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QueryLeasesForNodeRequest {
	return &QueryLeasesForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryLeasesForNodeRequest(addr hubtypes.NodeAddress, pagination *query.PageRequest) *QueryLeasesForNodeRequest {
	return &QueryLeasesForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
