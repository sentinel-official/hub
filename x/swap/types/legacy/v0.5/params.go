package v0_5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Params struct {
		SwapEnabled bool           `json:"swap_enabled"`
		SwapDenom   string         `json:"swap_denom"`
		ApproveBy   sdk.AccAddress `json:"approve_by"`
	}
)
