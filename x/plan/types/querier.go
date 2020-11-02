package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan             = "plan"
	QueryPlans            = "plans"
	QueryPlansForProvider = "plans_for_provider"

	QueryNodesForPlan = "nodes_for_plan"
)

type QueryPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQueryPlanParams(id uint64) QueryPlanParams {
	return QueryPlanParams{
		ID: id,
	}
}

type QueryPlansParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQueryPlansParams(skip, limit int) QueryPlansParams {
	return QueryPlansParams{
		Skip:  skip,
		Limit: limit,
	}
}

type QueryPlansForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Skip    int             `json:"skip"`
	Limit   int             `json:"limit"`
}

func NewQueryPlansForProviderParams(address hub.ProvAddress, skip, limit int) QueryPlansForProviderParams {
	return QueryPlansForProviderParams{
		Address: address,
		Skip:    skip,
		Limit:   limit,
	}
}

type QueryNodesForPlanParams struct {
	ID    uint64 `json:"id"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

func NewQueryNodesForPlanParams(id uint64, skip, limit int) QueryNodesForPlanParams {
	return QueryNodesForPlanParams{
		ID:    id,
		Skip:  skip,
		Limit: limit,
	}
}
