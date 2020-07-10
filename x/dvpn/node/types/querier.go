package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode             = "query_node"
	QueryNodes            = "query_nodes"
	QueryNodesForProvider = "query_nodes_for_provider"
)

type QueryNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQueryNodeParams(address hub.NodeAddress) QueryNodeParams {
	return QueryNodeParams{
		Address: address,
	}
}

type QueryNodesForProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryNodesForProviderParams(address hub.ProvAddress) QueryNodesForProviderParams {
	return QueryNodesForProviderParams{
		Address: address,
	}
}
