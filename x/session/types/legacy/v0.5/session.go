package v0_5

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	legacyhubtypes "github.com/sentinel-official/hub/types/legacy/v0.5"
)

type (
	Session struct {
		ID           uint64                   `json:"id"`
		Subscription uint64                   `json:"subscription"`
		Node         hubtypes.NodeAddress     `json:"node"`
		Address      sdk.AccAddress           `json:"address"`
		Duration     time.Duration            `json:"duration"`
		Bandwidth    legacyhubtypes.Bandwidth `json:"bandwidth"`
		Status       legacyhubtypes.Status    `json:"status"`
		StatusAt     time.Time                `json:"status_at"`
	}

	Sessions []Session
)
