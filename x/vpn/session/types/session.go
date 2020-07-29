package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Session struct {
	ID           uint64          `json:"id"`
	Subscription uint64          `json:"subscription"`
	Node         hub.NodeAddress `json:"node"`
	Address      sdk.AccAddress  `json:"address"`
	Duration     time.Duration   `json:"duration"`
	Bandwidth    hub.Bandwidth   `json:"bandwidth"`
	Status       hub.Status      `json:"status"`
	StatusAt     time.Time       `json:"status_at"`
}

func (s Session) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
ID:           %d
Subscription: %d
Node:         %s
Address:      %s
Duration:     %s
Bandwidth:    %s
Status:       %s
Status at:    %s
`), s.ID, s.Subscription, s.Node, s.Address, s.Duration, s.Bandwidth, s.Status, s.StatusAt)
}

type Sessions []Session
