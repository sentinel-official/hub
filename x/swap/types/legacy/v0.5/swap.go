package v0_5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Swap struct {
		TxHash   EthereumHash   `json:"tx_hash"`
		Receiver sdk.AccAddress `json:"receiver"`
		Amount   sdk.Coin       `json:"amount"`
	}

	Swaps []Swap
)
