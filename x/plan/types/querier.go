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
	Status hub.Status `json:"status"`
	Skip   int        `json:"skip"`
	Limit  int        `json:"limit"`
}

func NewQueryPlansParams(status hub.Status, skip, limit int) QueryPlansParams {
	return QueryPlansParams{
		Status: status,
		Skip:   skip,
		Limit:  limit,
	}
}

type QueryPlansForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Status  hub.Status      `json:"status"`
	Skip    int             `json:"skip"`
	Limit   int             `json:"limit"`
}

func NewQueryPlansForProviderParams(address hub.ProvAddress, status hub.Status, skip, limit int) QueryPlansForProviderParams {
	return QueryPlansForProviderParams{
		Address: address,
		Status:  status,
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
