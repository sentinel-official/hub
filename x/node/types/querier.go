package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hub "github.com/sentinel-official/hub/types"
)

func NewQueryNodeRequest(address hub.NodeAddress) *QueryNodeRequest {
	return &QueryNodeRequest{
		Address: address.String(),
	}
}

func NewQueryNodesRequest(status hub.Status, pagination *query.PageRequest) *QueryNodesRequest {
	return &QueryNodesRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForProviderRequest(address hub.ProvAddress, status hub.Status, pagination *query.PageRequest) *QueryNodesForProviderRequest {
	return &QueryNodesForProviderRequest{
		Address:    address.String(),
		Status:     status,
		Pagination: pagination,
	}
}
