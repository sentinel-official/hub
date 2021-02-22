package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

func (s Session) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Id:           %d
Subscription: %d
Node:         %s
Address:      %s
Duration:     %s
Bandwidth:    %s
Status:       %s
Status at:    %s
`), s.Id, s.Subscription, s.Node, s.Address, s.Duration, s.Bandwidth, s.Status, s.StatusAt)
}

func (s Session) Validate() error {
	if s.Id == 0 {
		return fmt.Errorf("id should not be zero")
	}
	if s.Subscription == 0 {
		return fmt.Errorf("subscription should not be zero")
	}
	if _, err := hub.NodeAddressFromBech32(s.Node); err != nil {
		return fmt.Errorf("node should not be nil or empty")
	}
	if _, err := sdk.AccAddressFromBech32(s.Address); err != nil {
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
