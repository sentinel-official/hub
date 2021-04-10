package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

func NewQueryProviderRequest(address hub.ProvAddress) *QueryProviderRequest {
	return &QueryProviderRequest{
		Address: address.String(),
	}
}

func NewQueryProvidersRequest(pagination *query.PageRequest) *QueryProvidersRequest {
	return &QueryProvidersRequest{
		Pagination: pagination,
	}
}
