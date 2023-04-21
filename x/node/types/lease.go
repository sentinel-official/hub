package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (l *Lease) GetNodeAddress() hubtypes.NodeAddress {
	if l.NodeAddress == "" {
		return nil
	}

	addr, err := hubtypes.NodeAddressFromBech32(l.NodeAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (l *Lease) GetAccountAddress() sdk.AccAddress {
	if l.AccountAddress == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(l.AccountAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (l *Lease) Validate() error {
	if l.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if l.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(l.NodeAddress); err != nil {
		return errors.Wrapf(err, "invalid node_address %s", l.NodeAddress)
	}
	if l.AccountAddress == "" {
		return fmt.Errorf("account_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(l.AccountAddress); err != nil {
		return errors.Wrapf(err, "invalid account_address %s", l.AccountAddress)
	}
	if l.Hours < 0 {
		return fmt.Errorf("hours cannot be negative")
	}
	if l.Hours == 0 {
		return fmt.Errorf("hours cannot be zero")
	}
	if l.Price.IsNegative() {
		return fmt.Errorf("price cannot be negative")
	}
	if l.Price.IsZero() {
		return fmt.Errorf("price cannot be zero")
	}
	if !l.Price.IsValid() {
		return fmt.Errorf("price must be valid")
	}

	return nil
}

type (
	Leases []Lease
)
