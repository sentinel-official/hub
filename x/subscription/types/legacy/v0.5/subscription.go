package v05

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	legacyhubtypes "github.com/sentinel-official/hub/types/legacy/v0.5"
)

type (
	Subscription struct {
		ID       uint64                `json:"id"`
		Owner    sdk.AccAddress        `json:"owner"`
		Node     hubtypes.NodeAddress  `json:"node"`
		Price    sdk.Coin              `json:"price"`
		Deposit  sdk.Coin              `json:"deposit"`
		Plan     uint64                `json:"plan"`
		Expiry   time.Time             `json:"expiry"`
		Free     sdk.Int               `json:"free"`
		Status   legacyhubtypes.Status `json:"status"`
		StatusAt time.Time             `json:"status_at"`
	}

	Subscriptions []Subscription
)
