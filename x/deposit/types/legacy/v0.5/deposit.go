package v0_5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Deposit struct {
		Address sdk.AccAddress `json:"address"`
		Coins   sdk.Coins      `json:"coins"`
	}

	Deposits []Deposit
)
