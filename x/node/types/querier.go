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

func NewQueryPayoutRequest(id uint64) *QueryPayoutRequest {
	return &QueryPayoutRequest{
		Id: id,
	}
}

func NewQueryPayoutsRequest(pagination *query.PageRequest) *QueryPayoutsRequest {
	return &QueryPayoutsRequest{
		Pagination: pagination,
	}
}

func NewQueryPayoutsForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QueryPayoutsForNodeRequest {
	return &QueryPayoutsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryPayoutsForNodeRequest(addr hubtypes.NodeAddress, pagination *query.PageRequest) *QueryPayoutsForNodeRequest {
	return &QueryPayoutsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
