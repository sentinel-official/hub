package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (s *Subscription) GetNode() hubtypes.NodeAddress {
	if s.Node == "" {
		return nil
	}

	address, err := hubtypes.NodeAddressFromBech32(s.Node)
	if err != nil {
		panic(err)
	}

	return address
}

func (s *Subscription) String() string {
	if s.Plan == 0 {
		return fmt.Sprintf(strings.TrimSpace(`
Id:        %d
Owner:     %s
Node:      %s
Price:     %s
Deposit:   %s
Free:      %s
Status:    %s
Status at: %s
`), s.Id, s.Owner, s.Node, s.Price, s.Deposit, s.Free, s.Status, s.StatusAt)
	}

	return fmt.Sprintf(strings.TrimSpace(`
Id:        %d
Owner:     %s
Plan:      %d
Expiry:    %s
Free:      %s
Status:    %s
Status at: %s
`), s.Id, s.Owner, s.Plan, s.Expiry, s.Free, s.Status, s.StatusAt)
}

func (s *Subscription) Amount(consumed sdk.Int) sdk.Coin {
	var (
		amount sdk.Int
		x      = hubtypes.Gigabyte.Quo(s.Price.Amount)
	)

	if x.IsPositive() {
		amount = hubtypes.NewBandwidth(consumed, sdk.ZeroInt()).
			CeilTo(x).
			Sum().Quo(x)
	} else {
		y := sdk.NewDecFromInt(s.Price.Amount).
			QuoInt(hubtypes.Gigabyte).
			Ceil().TruncateInt()
		amount = consumed.Mul(y)
	}

	return sdk.NewCoin(s.Price.Denom, amount)
}

func (s *Subscription) Validate() error {
	if s.Id == 0 {
		return fmt.Errorf("id should not be zero")
	}
	if _, err := sdk.AccAddressFromBech32(s.Owner); err != nil {
		return fmt.Errorf("owner should not nil or empty")
	}

	if s.Plan == 0 {
		if _, err := hubtypes.NodeAddressFromBech32(s.Node); err != nil {
			return fmt.Errorf("node should not be nil or empty")
		}
		if !s.Price.IsValid() {
			return fmt.Errorf("price should be valid")
		}
		if !s.Deposit.IsValid() {
			return fmt.Errorf("deposit should be valid")
		}
	} else {
		if s.Expiry.IsZero() {
			return fmt.Errorf("expiry should not be zero")
		}
	}

	if s.Free.IsNegative() {
		return fmt.Errorf("free should not be negative")
	}
	if !s.Status.IsValid() {
		return fmt.Errorf("status should be valid")
	}
	if s.StatusAt.IsZero() {
		return fmt.Errorf("status_at should not be zero")
	}

	return nil
}

type Subscriptions []Subscription
