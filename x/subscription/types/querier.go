package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QuerySubscription            = "Subscription"
	QuerySubscriptions           = "Subscriptions"
	QuerySubscriptionsForNode    = "SubscriptionsForNode"
	QuerySubscriptionsForPlan    = "SubscriptionsForPlan"
	QuerySubscriptionsForAddress = "SubscriptionsForAddress"

	QueryQuota  = "Quota"
	QueryQuotas = "Quotas"
)

func NewQuerySubscriptionRequest(id uint64) QuerySubscriptionRequest {
	return QuerySubscriptionRequest{
		Id: id,
	}
}

func NewQuerySubscriptionsRequest(pagination *query.PageRequest) QuerySubscriptionsRequest {
	return QuerySubscriptionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForNodeRequest(address string, pagination *query.PageRequest) QuerySubscriptionsForNodeRequest {
	return QuerySubscriptionsForNodeRequest{
		Address:    address,
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForPlanRequest(id uint64, pagination *query.PageRequest) QuerySubscriptionsForPlanRequest {
	return QuerySubscriptionsForPlanRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForAddressRequest(address string, status hub.Status, pagination *query.PageRequest) QuerySubscriptionsForAddressRequest {
	return QuerySubscriptionsForAddressRequest{
		Address:    address,
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryQuotaRequest(id uint64, address string) QueryQuotaRequest {
	return QueryQuotaRequest{
		Id:      id,
		Address: address,
	}
}

func NewQueryQuotasRequest(id uint64, pagination *query.PageRequest) QueryQuotasRequest {
	return QueryQuotasRequest{
		Id:         id,
		Pagination: pagination,
	}
}
