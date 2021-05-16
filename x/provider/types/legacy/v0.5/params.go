package v05

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Params struct {
		Deposit sdk.Coin `json:"deposit"`
	}
)
