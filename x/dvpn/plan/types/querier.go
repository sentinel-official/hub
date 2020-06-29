package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan             = "query_plan"
	QueryPlans            = "query_plans"
	QueryPlansForProvider = "query_plans_for_provider"
	QueryPlansForNode     = "query_plans_for_node"
	QueryNodesForPlan     = "query_nodes_for_plan"
)

type QueryPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQueryPlanParams(id uint64) QueryPlanParams {
	return QueryPlanParams{
		ID: id,
	}
}

type QueryPlansForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryPlansForProviderParams(address hub.ProvAddress) QueryPlansForProviderParams {
	return QueryPlansForProviderParams{
		Address: address,
	}
}

type QueryPlansForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQueryPlansForNodeParams(address hub.NodeAddress) QueryPlansForNodeParams {
	return QueryPlansForNodeParams{
		Address: address,
	}
}

type QueryNodesForPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQueryNodesForPlanParams(id uint64) QueryNodesForPlanParams {
	return QueryNodesForPlanParams{
		ID: id,
	}
}
