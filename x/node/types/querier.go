package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

const (
	QueryNode             = "Node"
	QueryNodes            = "Nodes"
	QueryNodesForProvider = "NodesForProvider"
)

func NewQueryNodeRequest(address string) QueryNodeRequest {
	return QueryNodeRequest{
		Address: address,
	}
}

func NewQueryNodesRequest(status hub.Status, pagination *query.PageRequest) QueryNodesRequest {
	return QueryNodesRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForProviderRequest(address string, status hub.Status, pagination *query.PageRequest) QueryNodesForProviderRequest {
	return QueryNodesForProviderRequest{
		Address:    address,
		Status:     status,
		Pagination: pagination,
	}
}
