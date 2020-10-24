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
	Status hub.Status `json:"status"`
	Page   int        `json:"page"`
	Limit  int        `json:"limit"`
}

func NewQueryNodesParams(status hub.Status, page, limit int) QueryNodesParams {
	return QueryNodesParams{
		Status: status,
		Page:   page,
		Limit:  limit,
	}
}

type QueryNodesForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Status  hub.Status      `json:"status"`
	Page    int             `json:"page"`
	Limit   int             `json:"limit"`
}

func NewQueryNodesForProviderParams(address hub.ProvAddress, status hub.Status, page, limit int) QueryNodesForProviderParams {
	return QueryNodesForProviderParams{
		Address: address,
		Status:  status,
		Page:    page,
		Limit:   limit,
	}
}
