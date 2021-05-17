package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewQueryNodeRequest(address hubtypes.NodeAddress) *QueryNodeRequest {
	return &QueryNodeRequest{
		Address: address.String(),
	}
}

func NewQueryNodesRequest(status hubtypes.Status, pagination *query.PageRequest) *QueryNodesRequest {
	return &QueryNodesRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryNodesForProviderRequest(address hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) *QueryNodesForProviderRequest {
	return &QueryNodesForProviderRequest{
		Address:    address.String(),
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
