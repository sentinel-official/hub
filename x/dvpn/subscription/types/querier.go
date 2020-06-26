package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan            = "query_plan"
	QueryPlans           = "query_plans"
	QueryPlansOfProvider = "query_plans_of_provider"
	QueryNodesOfPlan     = "query_nodes_of_plan"

	QuerySubscription           = "query_subscription"
	QuerySubscriptions          = "query_subscriptions"
	QuerySubscriptionsOfAddress = "query_subscriptions_of_address"
	QuerySubscriptionsOfPlan    = "query_subscriptions_of_plan"
	QuerySubscriptionsOfNode    = "query_subscriptions_of_node"
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

type QuerySubscriptionParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySubscriptionParams(id uint64) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

type QuerySubscriptionsOfAddressParams struct {
	Address sdk.AccAddress `json:"address"`
}

func NewQuerySubscriptionsOfAddressParams(address sdk.AccAddress) QuerySubscriptionsOfAddressParams {
	return QuerySubscriptionsOfAddressParams{
		Address: address,
	}
}

type QuerySubscriptionsOfPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySubscriptionsOfPlanParams(id uint64) QuerySubscriptionsOfPlanParams {
	return QuerySubscriptionsOfPlanParams{
		ID: id,
	}
}

type QuerySubscriptionsOfNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQuerySubscriptionsOfNodeParams(address hub.NodeAddress) QuerySubscriptionsOfNodeParams {
	return QuerySubscriptionsOfNodeParams{
		Address: address,
	}
}
