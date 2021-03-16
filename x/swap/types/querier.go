package types

const (
	QuerySwap  = "Swap"
	QuerySwaps = "Swaps"
)

type QuerySwapParams struct {
	TxHash EthereumHash `json:"tx_hash"`
}

func NewQuerySwapParams(txHash EthereumHash) QuerySwapParams {
	return QuerySwapParams{
		TxHash: txHash,
	}
}

type QuerySwapsParams struct {
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

func NewQuerySwapsParams(skip, limit int) QuerySwapsParams {
	return QuerySwapsParams{
		Skip:  skip,
		Limit: limit,
	}
}
