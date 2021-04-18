package v0_5

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	legacyhubtypes "github.com/sentinel-official/hub/types/legacy/v0.5"
)

type (
	Plan struct {
		ID       uint64                `json:"id"`
		Provider hubtypes.ProvAddress  `json:"provider"`
		Price    sdk.Coins             `json:"price"`
		Validity time.Duration         `json:"validity"`
		Bytes    sdk.Int               `json:"bytes"`
		Status   legacyhubtypes.Status `json:"status"`
		StatusAt time.Time             `json:"status_at"`
	}

	Plans []Plan
)
