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

func NewQuerySubscriptionsForAddressRequest(address sdk.AccAddress, status hubtypes.Status, pagination *query.PageRequest) *QuerySubscriptionsForAddressRequest {
	return &QuerySubscriptionsForAddressRequest{
		Address:    address.String(),
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryQuotaRequest(id uint64, address sdk.AccAddress) *QueryQuotaRequest {
	return &QueryQuotaRequest{
		Id:      id,
		Address: address.String(),
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
