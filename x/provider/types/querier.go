package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQueryProviderRequest(addr hubtypes.ProvAddress) *QueryProviderRequest {
	return &QueryProviderRequest{
		Address: addr.String(),
	}
}

func NewQueryProvidersRequest(status hubtypes.Status, pagination *query.PageRequest) *QueryProvidersRequest {
	return &QueryProvidersRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
