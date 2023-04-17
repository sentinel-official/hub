package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return nil
}

type (
	Leases []Lease
)
