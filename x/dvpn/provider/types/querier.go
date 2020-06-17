package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryProvider  = "query_provider"
	QueryProviders = "query_providers"
)

type QueryProviderParams struct {
	ID hub.ProviderID `json:"id"`
}

func NewQueryProviderParams(id hub.ProviderID) QueryProviderParams {
	return QueryProviderParams{
		ID: id,
	}
}
