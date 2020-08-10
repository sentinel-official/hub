package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryProvider  = "provider"
	QueryProviders = "providers"
)

type QueryProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryProviderParams(address hub.ProvAddress) QueryProviderParams {
	return QueryProviderParams{
		Address: address,
	}
}

type QueryProvidersParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryProvidersParams(page, limit int) QueryProvidersParams {
	return QueryProvidersParams{
		Page:  page,
		Limit: limit,
	}
}
