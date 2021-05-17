package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQueryProviderRequest(address hubtypes.ProvAddress) *QueryProviderRequest {
	return &QueryProviderRequest{
		Address: address.String(),
	}
}

func NewQueryProvidersRequest(pagination *query.PageRequest) *QueryProvidersRequest {
	return &QueryProvidersRequest{
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
