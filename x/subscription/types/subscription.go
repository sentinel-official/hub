package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Subscription) GetNode() hubtypes.NodeAddress {
	if m.Node == "" {
		return nil
	}

	address, err := hubtypes.NodeAddressFromBech32(m.Node)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Subscription) GetOwner() sdk.AccAddress {
	if m.Owner == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Subscription) Amount(consumed sdk.Int) sdk.Coin {
	var (
		amount sdk.Int
		x      = hubtypes.Gigabyte.Quo(m.Price.Amount)
	)

	if x.IsPositive() {
		amount = hubtypes.NewBandwidth(consumed, sdk.ZeroInt()).
			CeilTo(x).
			Sum().Quo(x)
	} else {
		y := sdk.NewDecFromInt(m.Price.Amount).
			QuoInt(hubtypes.Gigabyte).
			Ceil().TruncateInt()
		amount = consumed.Mul(y)
	}

	return sdk.NewCoin(m.Price.Denom, amount)
}

func (m *Subscription) Validate() error {
	if m.Id == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.Owner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return errors.Wrapf(err, "invalid owner %s", m.Owner)
	}
	if m.Node == "" && m.Plan == 0 {
		return fmt.Errorf("both node and plan cannot be empty")
	}
	if m.Node != "" && m.Plan != 0 {
		return fmt.Errorf("either node or plan must be empty")
	}
	if m.Node != "" {
		if _, err := hubtypes.NodeAddressFromBech32(m.Node); err != nil {
			return errors.Wrapf(err, "invalid node %s", m.Node)
		}
		if m.Price.IsZero() {
			return fmt.Errorf("price cannot be zero")
		}
		if !m.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
		if m.Deposit.IsZero() {
			return fmt.Errorf("deposit cannot be zero")
		}
		if !m.Deposit.IsValid() {
			return fmt.Errorf("deposit must be valid")
		}
	}
	if m.Plan != 0 {
		if m.Denom != "" {
			if err := sdk.ValidateDenom(m.Denom); err != nil {
				return errors.Wrapf(err, "invalid denom %s", m.Denom)
			}
		}
		if m.Expiry.IsZero() {
			return fmt.Errorf("expiry cannot be zero")
		}
	}
	if m.Free.IsNegative() {
		return fmt.Errorf("free cannot not be negative")
	}
	if !m.Status.IsValid() {
		return fmt.Errorf("status must be valid")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

type (
	Subscriptions []Subscription
)
