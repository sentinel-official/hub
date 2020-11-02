package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySubscription            = "subscription"
	QuerySubscriptions           = "subscriptions"
	QuerySubscriptionsForAddress = "subscriptions_for_address"
	QuerySubscriptionsForPlan    = "subscriptions_for_plan"
	QuerySubscriptionsForNode    = "subscriptions_for_node"

	QueryQuota  = "quota"
	QueryQuotas = "quotas"
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
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQuerySubscriptionsParams(skip, limit int) QuerySubscriptionsParams {
	return QuerySubscriptionsParams{
		Skip:  skip,
		Limit: limit,
	}
}

type QuerySubscriptionsForAddressParams struct {
	Address sdk.AccAddress `json:"address"`
	Skip    int            `json:"skip"`
	Limit   int            `json:"limit"`
}

func NewQuerySubscriptionsForAddressParams(address sdk.AccAddress, skip, limit int) QuerySubscriptionsForAddressParams {
	return QuerySubscriptionsForAddressParams{
		Address: address,
		Skip:    skip,
		Limit:   limit,
	}
}

type QuerySubscriptionsForPlanParams struct {
	ID    uint64 `json:"id"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

func NewQuerySubscriptionsForPlanParams(id uint64, skip, limit int) QuerySubscriptionsForPlanParams {
	return QuerySubscriptionsForPlanParams{
		ID:    id,
		Skip:  skip,
		Limit: limit,
	}
}

type QuerySubscriptionsForNodeParams struct {
	Address hub.NodeAddress `json:"address"`
	Skip    int             `json:"skip"`
	Limit   int             `json:"limit"`
}

func NewQuerySubscriptionsForNodeParams(address hub.NodeAddress, skip, limit int) QuerySubscriptionsForNodeParams {
	return QuerySubscriptionsForNodeParams{
		Address: address,
		Skip:    skip,
		Limit:   limit,
	}
}

type QueryQuotaParams struct {
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewQueryQuotaParams(id uint64, address sdk.AccAddress) QueryQuotaParams {
	return QueryQuotaParams{
		ID:      id,
		Address: address,
	}
}

type QueryQuotasParams struct {
	ID    uint64 `json:"id"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

func NewQueryQuotasParams(id uint64, skip, limit int) QueryQuotasParams {
	return QueryQuotasParams{
		ID:    id,
		Skip:  skip,
		Limit: limit,
	}
}
