package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySubscription            = "query_subscription"
	QuerySubscriptions           = "query_subscriptions"
	QuerySubscriptionsForAddress = "query_subscriptions_for_address"
	QuerySubscriptionsForPlan    = "query_subscriptions_for_plan"
	QuerySubscriptionsForNode    = "query_subscriptions_for_node"

	QueryQuotaForSubscription  = "query_quota_for_subscription"
	QueryQuotasForSubscription = "query_quotas_for_subscription"
)

type QuerySubscriptionParams struct {
	ID uint64 `json:"id"`
}

func NewQuerySubscriptionParams(id uint64) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

type QuerySubscriptionsParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQuerySubscriptionsParams(page, limit int) QuerySubscriptionsParams {
	return QuerySubscriptionsParams{
		Page:  page,
		Limit: limit,
	}
}

type QuerySubscriptionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
	Page    int            `json:"page"`
	Limit   int            `json:"limit"`
}

func NewQuerySubscriptionsForAddressParams(address sdk.AccAddress, page, limit int) QuerySubscriptionsForAddressParams {
	return QuerySubscriptionsForAddressParams{
		Address: address,
		Page:    page,
		Limit:   limit,
	}
}

type QuerySubscriptionsForPlanParams struct {
	ID    uint64 `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func NewQuerySubscriptionsForPlanParams(id uint64, page, limit int) QuerySubscriptionsForPlanParams {
	return QuerySubscriptionsForPlanParams{
		ID:    id,
		Page:  page,
		Limit: limit,
	}
}

type QuerySubscriptionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
}

func NewQuerySubscriptionsForNodeParams(address hub.NodeAddress, page, limit int) QuerySubscriptionsForNodeParams {
	return QuerySubscriptionsForNodeParams{
		Address: address,
		Page:    page,
		Limit:   limit,
	}
}

type QueryQuotaForSubscriptionParams struct {
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewQueryQuotaForSubscriptionParams(id uint64, address sdk.AccAddress) QueryQuotaForSubscriptionParams {
	return QueryQuotaForSubscriptionParams{
		ID:      id,
		Address: address,
	}
}

type QueryQuotasForSubscriptionParams struct {
	ID    uint64 `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func NewQueryQuotasForSubscriptionParams(id uint64, page, limit int) QueryQuotasForSubscriptionParams {
	return QueryQuotasForSubscriptionParams{
		ID:    id,
		Page:  page,
		Limit: limit,
	}
}
