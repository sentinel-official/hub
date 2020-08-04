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
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryPlansParams(page, limit int) QueryPlansParams {
	return QueryPlansParams{
		Page:  page,
		Limit: limit,
	}
}

type QueryPlansForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
}

func NewQueryPlansForProviderParams(address hub.ProvAddress, page, limit int) QueryPlansForProviderParams {
	return QueryPlansForProviderParams{
		Address: address,
		Page:    page,
		Limit:   limit,
	}
}

type QueryNodesForPlanParams struct {
	ID    uint64 `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func NewQueryNodesForPlanParams(id uint64, page, limit int) QueryNodesForPlanParams {
	return QueryNodesForPlanParams{
		ID:    id,
		Page:  page,
		Limit: limit,
	}
}
