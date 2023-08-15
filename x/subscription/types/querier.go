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

func NewQueryAllocationRequest(id uint64, addr sdk.AccAddress) *QueryAllocationRequest {
	return &QueryAllocationRequest{
		Id:      id,
		Address: addr.String(),
	}
}

func NewQueryAllocationsRequest(id uint64, pagination *query.PageRequest) *QueryAllocationsRequest {
	return &QueryAllocationsRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryPayoutRequest(id uint64) *QueryPayoutRequest {
	return &QueryPayoutRequest{
		Id: id,
	}
}

func NewQueryPayoutsRequest(pagination *query.PageRequest) *QueryPayoutsRequest {
	return &QueryPayoutsRequest{
		Pagination: pagination,
	}
}

func NewQueryPayoutsForAccountRequest(addr sdk.AccAddress, pagination *query.PageRequest) *QueryPayoutsForAccountRequest {
	return &QueryPayoutsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryPayoutsForNodeRequest(addr hubtypes.NodeAddress, pagination *query.PageRequest) *QueryPayoutsForNodeRequest {
	return &QueryPayoutsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
