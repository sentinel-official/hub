package v0_5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Quota struct {
		Address   sdk.AccAddress `json:"address"`
		Consumed  sdk.Int        `json:"consumed"`
		Allocated sdk.Int        `json:"allocated"`
	}

	Quotas []Quota
)
