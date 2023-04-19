package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQuerySubscriptionRequest(id uint64) *QuerySubscriptionRequest {
	return &QuerySubscriptionRequest{
		Id: id,
	}
}

func NewQuerySubscriptionsRequest(pagination *query.PageRequest) *QuerySubscriptionsRequest {
	return &QuerySubscriptionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QuerySubscriptionsForAccountRequest {
	return &QuerySubscriptionsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForNodeRequest(addr hubtypes.NodeAddress, pagination *query.PageRequest) *QuerySubscriptionsForNodeRequest {
	return &QuerySubscriptionsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForPlanRequest(id uint64, pagination *query.PageRequest) *QuerySubscriptionsForPlanRequest {
	return &QuerySubscriptionsForPlanRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryQuotaRequest(id uint64, addr sdk.AccAddress) *QueryQuotaRequest {
	return &QueryQuotaRequest{
		Id:      id,
		Address: addr.String(),
	}
}

func NewQueryQuotasRequest(id uint64, pagination *query.PageRequest) *QueryQuotasRequest {
	return &QueryQuotasRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
