package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryProvider  = "query_provider"
	QueryProviders = "query_providers"
)

type QueryProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryProviderParams(address hub.ProvAddress) QueryProviderParams {
	return QueryProviderParams{
		Address: address,
	}
}
