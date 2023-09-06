package types

import (
	"fmt"
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
)

func (m *Node) GetAddress() hubtypes.NodeAddress {
	if m.Address == "" {
		return nil
	}

	addr, err := hubtypes.NodeAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Node) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.GigabytePrices == nil {
		return fmt.Errorf("gigabyte_prices cannot be nil")
	}
	if m.GigabytePrices.Len() == 0 {
		return fmt.Errorf("gigabyte_prices cannot be empty")
	}
	if m.GigabytePrices.IsAnyNil() {
		return fmt.Errorf("gigabyte_prices cannot contain nil")
	}
	if !m.GigabytePrices.IsValid() {
		return fmt.Errorf("gigabyte_prices must be valid")
	}
	if m.HourlyPrices == nil {
		return fmt.Errorf("hourly_prices cannot be nil")
	}
	if m.HourlyPrices.Len() == 0 {
		return fmt.Errorf("hourly_prices cannot be empty")
	}
	if m.HourlyPrices.IsAnyNil() {
		return fmt.Errorf("hourly_prices cannot contain nil")
	}
	if !m.HourlyPrices.IsValid() {
		return fmt.Errorf("hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return fmt.Errorf("remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d chars", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid remote_url %s", m.RemoteURL)
	}
	if remoteURL.Scheme != "https" {
		return fmt.Errorf("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return fmt.Errorf("remote_url port cannot be empty")
	}

	if m.InactiveAt.IsZero() {
		if !m.Status.Equal(hubtypes.StatusInactive) {
			return fmt.Errorf("invalid inactive_at %s; expected positive", m.InactiveAt)
		}
	}
	if !m.InactiveAt.IsZero() {
		if !m.Status.Equal(hubtypes.StatusActive) {
			return fmt.Errorf("invalid inactive_at %s; expected zero", m.InactiveAt)
		}
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

func (m *Node) GigabytePrice(denom string) (sdk.Coin, bool) {
	for _, v := range m.GigabytePrices {
		if v.Denom == denom {
			return v, true
		}
	}

	return sdk.Coin{}, false
}

func (m *Node) HourlyPrice(denom string) (sdk.Coin, bool) {
	for _, v := range m.HourlyPrices {
		if v.Denom == denom {
			return v, true
		}
	}

	return sdk.Coin{}, false
}

type (
	Nodes []Node
)
