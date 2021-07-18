package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Session) GetAddress() sdk.AccAddress {
	if m.Address == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Session) GetNode() hubtypes.NodeAddress {
	if m.Node == "" {
		return nil
	}

	address, err := hubtypes.NodeAddressFromBech32(m.Node)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Session) Validate() error {
	if m.Id == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.Subscription == 0 {
		return fmt.Errorf("subscription cannot be zero")
	}
	if m.Node == "" {
		return fmt.Errorf("node cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Node); err != nil {
		return errors.Wrapf(err, "invalid node %s", m.Node)
	}
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Duration < 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	if m.Bandwidth.IsAnyNegative() {
		return fmt.Errorf("bandwidth cannot be negative")
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
	Sessions []Session
)
