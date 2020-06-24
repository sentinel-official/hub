package types

import (
	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode            = "query_node"
	QueryNodes           = "query_nodes"
	QueryNodesOfProvider = "query_nodes_of_provider"
)

type QueryNodeParams struct {
	Address hub.NodeAddress `json:"address"`
}

func NewQueryNodeParams(address hub.NodeAddress) QueryNodeParams {
	return QueryNodeParams{
		Address: address,
	}
}

type QueryNodesOfProviderParams struct {
	Address hub.ProvAddress `json:"address"`
}

func NewQueryNodesOfProviderParams(address hub.ProvAddress) QueryNodesOfProviderParams {
	return QueryNodesOfProviderParams{
		Address: address,
	}
}
