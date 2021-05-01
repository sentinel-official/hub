package v0_5

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	legacyhubtypes "github.com/sentinel-official/hub/types/legacy/v0.5"
)

type (
	Node struct {
		Address   hubtypes.NodeAddress  `json:"address"`
		Provider  hubtypes.ProvAddress  `json:"provider"`
		Price     sdk.Coins             `json:"price"`
		RemoteURL string                `json:"remote_url"`
		Status    legacyhubtypes.Status `json:"status"`
		StatusAt  time.Time             `json:"status_at"`
	}

	Nodes []Node
)
