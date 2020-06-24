package types

import (
	"fmt"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryKeyProvider  = "query_provider"
	QueryKeyProviders = "query_providers"
)

var (
	QueryProviderPath  = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyProvider)
	QueryProvidersPath = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyProviders)
)

type QueryProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryProviderParams(address hub.ProvAddress) QueryProviderParams {
	return QueryProviderParams{
		Address: address,
	}
}
