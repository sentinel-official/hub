package types

import (
	"fmt"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryKeyNode            = "query_node"
	QueryKeyNodes           = "query_nodes"
	QueryKeyNodesOfProvider = "query_nodes_of_provider"
)

var (
	QueryNodePath            = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyNode)
	QueryNodesPath           = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyNodes)
	QueryNodesOfProviderPath = fmt.Sprintf("custom/%s/%s/%s", StoreKey, QuerierRoute, QueryKeyNodesOfProvider)
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
