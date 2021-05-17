package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

func NewQuerySwapRequest(txHash EthereumHash) *QuerySwapRequest {
	return &QuerySwapRequest{
		TxHash: txHash.Bytes(),
	}
}

func NewQuerySwapsRequest(pagination *query.PageRequest) *QuerySwapsRequest {
	return &QuerySwapsRequest{
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
