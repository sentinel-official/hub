package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode             = "Node"
	QueryNodes            = "Nodes"
	QueryNodesForProvider = "NodesForProvider"
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
	Skip   int        `json:"skip"`
	Limit  int        `json:"limit"`
}

func NewQueryNodesParams(status hub.Status, skip, limit int) QueryNodesParams {
	return QueryNodesParams{
		Status: status,
		Skip:   skip,
		Limit:  limit,
	}
}

type QueryNodesForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
	Status  hub.Status      `json:"status"`
	Skip    int             `json:"skip"`
	Limit   int             `json:"limit"`
}

func NewQueryNodesForProviderParams(address hub.ProvAddress, status hub.Status, skip, limit int) QueryNodesForProviderParams {
	return QueryNodesForProviderParams{
		Address: address,
		Status:  status,
		Skip:    skip,
		Limit:   limit,
	}
}
