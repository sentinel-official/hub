package v0_5

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Params struct {
		Deposit          sdk.Coin      `json:"deposit"`
		InactiveDuration time.Duration `json:"inactive_duration"`
	}
)
