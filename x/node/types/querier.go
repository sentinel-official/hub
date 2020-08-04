package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode             = "node"
	QueryNodes            = "nodes"
	QueryNodesForProvider = "nodes_for_provider"
)

type QueryNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQueryNodeParams(address hub.NodeAddress) QueryNodeParams {
	return QueryNodeParams{
		Address: address,
	}
}

type QueryNodesParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewQueryNodesParams(page, limit int) QueryNodesParams {
	return QueryNodesParams{
		Page:  page,
		Limit: limit,
	}
}

type QueryNodesForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
}

func NewQueryNodesForProviderParams(address hub.ProvAddress, page, limit int) QueryNodesForProviderParams {
	return QueryNodesForProviderParams{
		Address: address,
		Page:    page,
		Limit:   limit,
	}
}
