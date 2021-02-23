package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

const (
	QueryProvider  = "Provider"
	QueryProviders = "Providers"
)

func NewQueryProviderRequest(address string) QueryProviderRequest {
	return QueryProviderRequest{
		Address: address,
	}
}

func NewQueryProvidersRequest(pagination *query.PageRequest) QueryProvidersRequest {
	return QueryProvidersRequest{
		Pagination: pagination,
	}
}
