package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Subscription struct {
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`

	Plan      uint64        `json:"plan,omitempty"`
	Duration  time.Duration `json:"duration,omitempty"` // Remaining duration
	ExpiresAt time.Time     `json:"expires_at,omitempty"`

	Node    hub.NodeAddress `json:"node,omitempty"`
	Price   sdk.Coin        `json:"price,omitempty"`
	Deposit sdk.Coin        `json:"deposit,omitempty"`

	Bandwidth hub.Bandwidth `json:"bandwidth"` // Remaining bandwidth
	Status    hub.Status    `json:"status"`
	StatusAt  time.Time     `json:"status_at"`
}

func (s Subscription) String() string {
	if s.Plan > 0 {
		return fmt.Sprintf(strings.TrimSpace(`
ID:         %d
Address:    %s
Plan:       %d
Bandwidth:  %s
Duration:   %s
Expires at: %s
Status:     %s
Status at:  %s
`), s.ID, s.Address, s.Plan, s.Bandwidth, s.Duration, s.ExpiresAt, s.Status, s.StatusAt)
	}

	return fmt.Sprintf(strings.TrimSpace(`
ID:        %d
Address:   %s
Node:      %s
Price:     %s
Deposit:   %s
Bandwidth: %s
Status:    %s
Status at: %s
`), s.ID, s.Address, s.Node, s.Price, s.Deposit, s.Bandwidth, s.Status, s.StatusAt)
}

type Subscriptions []Subscription
