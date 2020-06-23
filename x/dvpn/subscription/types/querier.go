package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan            = "query_plan"
	QueryPlans           = "query_plans"
	QueryPlansOfProvider = "query_plans_of_provider"
	QueryNodesOfPlan     = "query_nodes_of_plan"
)

type QueryPlanParams struct {
	Address hub.ProvAddress `json:"address"`
	ID      uint64          `json:"id"`
}

func NewQueryPlanParams(address hub.ProvAddress, id uint64) QueryPlanParams {
	return QueryPlanParams{
		Address: address,
		ID:      id,
	}
}

type QueryPlansOfProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryPlansOfProviderParams(address hub.ProvAddress) QueryPlansOfProviderParams {
	return QueryPlansOfProviderParams{
		Address: address,
	}
}

type QueryNodesOfPlanParams struct {
	Address hub.ProvAddress `json:"address"`
	ID      uint64          `json:"id"`
}

func NewQueryNodesOfPlanParams(address hub.ProvAddress, id uint64) QueryNodesOfPlanParams {
	return QueryNodesOfPlanParams{
		Address: address,
		ID:      id,
	}
}
