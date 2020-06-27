package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryPlan             = "query_plan"
	QueryPlans            = "query_plans"
	QueryPlansForProvider = "query_plans_for_provider"
	QueryNodesForPlan     = "query_nodes_for_plan"

	QuerySubscription            = "query_subscription"
	QuerySubscriptions           = "query_subscriptions"
	QuerySubscriptionsForAddress = "query_subscriptions_for_address"
	QuerySubscriptionsForPlan    = "query_subscriptions_for_plan"
	QuerySubscriptionsForNode    = "query_subscriptions_for_node"
	QueryMembersForSubscription  = "query_members_for_subscription"
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

type QueryNodesForPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQueryNodesForPlanParams(id uint64) QueryNodesForPlanParams {
	return QueryNodesForPlanParams{
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

type QuerySubscriptionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
}

func NewQuerySubscriptionsForAddressParams(address sdk.AccAddress) QuerySubscriptionsForAddressParams {
	return QuerySubscriptionsForAddressParams{
		Address: address,
	}
}

type QuerySubscriptionsForPlanParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySubscriptionsForPlanParams(id uint64) QuerySubscriptionsForPlanParams {
	return QuerySubscriptionsForPlanParams{
		ID: id,
	}
}

type QuerySubscriptionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQuerySubscriptionsForNodeParams(address hub.NodeAddress) QuerySubscriptionsForNodeParams {
	return QuerySubscriptionsForNodeParams{
		Address: address,
	}
}

type QueryMembersForSubscriptionParams struct {
	ID uint64 `json:"id"`
}

func NewQueryMembersForSubscriptionParams(id uint64) QueryMembersForSubscriptionParams {
	return QueryMembersForSubscriptionParams{
		ID: id,
	}
}
