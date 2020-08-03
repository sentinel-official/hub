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

	Plan      uint64    `json:"plan,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	Node    hub.NodeAddress `json:"node,omitempty"`
	Price   sdk.Coin        `json:"price,omitempty"`
	Deposit sdk.Coin        `json:"deposit,omitempty"`

	Unallocated hub.Bandwidth `json:"unallocated"`
	Status      hub.Status    `json:"status"`
	StatusAt    time.Time     `json:"status_at"`
}

func (s Subscription) String() string {
	if s.Plan == 0 {
		return fmt.Sprintf(strings.TrimSpace(`
ID:          %d
Address:     %s
Node:        %s
Price:       %s
Deposit:     %s
Unallocated: %s
Status:      %s
Status at:   %s
`), s.ID, s.Address, s.Node, s.Price, s.Deposit, s.Unallocated, s.Status, s.StatusAt)
	}

	return fmt.Sprintf(strings.TrimSpace(`
ID:          %d
Address:     %s
Plan:        %d
Expires at:  %s
Unallocated: %s
Status:      %s
Status at:   %s
`), s.ID, s.Address, s.Plan, s.ExpiresAt, s.Unallocated, s.Status, s.StatusAt)
}

func (s Subscription) Amount(consumed hub.Bandwidth) sdk.Coin {
	amount := consumed.
		CeilTo(hub.Gigabyte.Quo(s.Price.Amount)).
		Sum().
		Mul(s.Price.Amount).
		Quo(hub.Gigabyte)

	coin := sdk.NewCoin(s.Price.Denom, amount)
	if s.Deposit.IsLT(coin) {
		return s.Deposit
	}

	return coin
}

type Subscriptions []Subscription

type Quota struct {
	Address   sdk.AccAddress `json:"address"`
	Consumed  hub.Bandwidth  `json:"consumed"`
	Allocated hub.Bandwidth  `json:"allocated"`
}

func (q Quota) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address:   %s
Consumed:  %s
Allocated: %s
`), q.Address, q.Consumed, q.Allocated)
}

type Quotas []Quota
