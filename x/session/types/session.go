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

func (s Session) Validate() error {
	if s.ID == 0 {
		return fmt.Errorf("id should not be zero")
	}
	if s.Subscription == 0 {
		return fmt.Errorf("subscription should not be zero")
	}
	if s.Node == nil || s.Node.Empty() {
		return fmt.Errorf("node should not be nil or empty")
	}
	if s.Address == nil || s.Address.Empty() {
		return fmt.Errorf("address should not be nil or empty")
	}
	if s.Duration <= 0 {
		return fmt.Errorf("duration should be positive")
	}
	if s.Bandwidth.IsValid() {
		return fmt.Errorf("bandwidth should be valid")
	}
	if !s.Status.Equal(hub.StatusActive) && !s.Status.Equal(hub.StatusInactive) {
		return fmt.Errorf("status should be either active or inactive")
	}
	if s.StatusAt.IsZero() {
		return fmt.Errorf("status_at should not be zero")
	}

	return nil
}

type Sessions []Session
