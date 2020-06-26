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
	return fmt.Sprintf(strings.TrimSpace(`
ID: %d
Address: %s
Plan: %d
Duration: %s
Expires at: %s
Node: %s
Price: %s
Deposit: %s
Bandwidth: %s
Status: %s
Status at: %s
`), s.ID, s.Address, s.Plan, s.Duration, s.ExpiresAt, s.Node, s.Price, s.Deposit, s.Bandwidth, s.Status, s.StatusAt)
}

type Subscriptions []Subscription
