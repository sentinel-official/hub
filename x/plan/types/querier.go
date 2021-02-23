package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan             = "Plan"
	QueryPlans            = "Plans"
	QueryPlansForProvider = "PlansForProvider"

	QueryNodesForPlan = "NodesForPlan"
)

func NewQueryPlanRequest(id uint64) QueryPlanRequest {
	return QueryPlanRequest{
		Id: id,
	}
}

func NewQueryPlansRequest(status hub.Status, pagination *query.PageRequest) QueryPlansRequest {
	return QueryPlansRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryPlansForProviderRequest(address string, status hub.Status, pagination *query.PageRequest) QueryPlansForProviderRequest {
	return QueryPlansForProviderRequest{
		Address:    address,
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForPlanRequest(id uint64, pagination *query.PageRequest) QueryNodesForPlanRequest {
	return QueryNodesForPlanRequest{
		Id:         id,
		Pagination: pagination,
	}
}
