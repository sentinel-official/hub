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

	addr, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Session) GetNodeAddress() hubtypes.NodeAddress {
	if m.NodeAddress == "" {
		return nil
	}

	addr, err := hubtypes.NodeAddressFromBech32(m.NodeAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Session) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.SubscriptionID == 0 {
		return fmt.Errorf("subscription_id cannot be zero")
	}
	if m.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return errors.Wrapf(err, "invalid node_address %s", m.NodeAddress)
	}
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Bandwidth.IsAnyNil() {
		return fmt.Errorf("bandwidth cannot be empty")
	}
	if m.Bandwidth.IsAnyNegative() {
		return fmt.Errorf("bandwidth cannot be negative")
	}
	if m.Duration < 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	if m.InactiveAt.IsZero() {
		return fmt.Errorf("inactive_at cannot be zero")
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactivePending) {
		return fmt.Errorf("status must be oneof [active, inactive_pending]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

type (
	Sessions []Session
)
