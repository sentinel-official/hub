package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

const (
	QuerySwap  = "Swap"
	QuerySwaps = "Swaps"
)

func NewQuerySwapRequest(txHash []byte) QuerySwapRequest {
	return QuerySwapRequest{
		TxHash: txHash,
	}
}

type QuerySwapsParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQuerySwapsRequest(pagination *query.PageRequest) QuerySwapsRequest {
	return QuerySwapsRequest{
		Pagination: pagination,
	}
}
