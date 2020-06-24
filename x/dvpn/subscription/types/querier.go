package types

import (
	"fmt"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryKeyPlan            = "query_plan"
	QueryKeyPlans           = "query_plans"
	QueryKeyPlansOfProvider = "query_plans_of_provider"
	QueryKeyNodesOfPlan     = "query_nodes_of_plan"
)

var (
	QueryPlanPath            = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyPlan)
	QueryPlansPath           = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyPlans)
	QueryPlansOfProviderPath = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyPlansOfProvider)
	QueryNodesOfPlanPath     = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyNodesOfPlan)
)

type QueryPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQueryPlanParams(id uint64) QueryPlanParams {
	return QueryPlanParams{
		ID: id,
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
	ID uint64 `json:"id"`
}

func NewQueryNodesOfPlanParams(id uint64) QueryNodesOfPlanParams {
	return QueryNodesOfPlanParams{
		ID: id,
	}
}
