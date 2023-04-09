package types

import (
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
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
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.GigabytePrices != nil {
		if m.GigabytePrices.Len() == 0 {
			return fmt.Errorf("gigabyte_prices cannot be empty")
		}
		if !m.GigabytePrices.IsValid() {
			return fmt.Errorf("gigabyte_prices must be valid")
		}
	}
	if m.HourlyPrices != nil {
		if m.HourlyPrices.Len() == 0 {
			return fmt.Errorf("hourly_prices cannot be empty")
		}
		if !m.HourlyPrices.IsValid() {
			return fmt.Errorf("hourly_prices must be valid")
		}
	}
	if m.RemoteURL == "" {
		return fmt.Errorf("remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return errors.Wrapf(err, "invalid remote_url %s", m.RemoteURL)
	}
	if remoteURL.Scheme != "https" {
		return fmt.Errorf("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return fmt.Errorf("remote_url port cannot be empty")
	}

	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt < 0 {
		return fmt.Errorf("status_at cannot be negative")
	}
	if m.StatusAt == 0 {
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

func (m *Node) Bytes(coin sdk.Coin) (sdk.Int, error) {
	price, found := m.GigabytePrice(coin.Denom)
	if !found {
		return sdk.ZeroInt(), fmt.Errorf("price for denom %s does not exist", coin.Denom)
	}

	x := hubtypes.Gigabyte.Quo(price.Amount)
	if x.IsPositive() {
		return coin.Amount.Mul(x), nil
	}

	y := sdk.NewDecFromInt(price.Amount).
		QuoInt(hubtypes.Gigabyte).
		Ceil().TruncateInt()

	return coin.Amount.Quo(y), nil
}

type (
	Nodes []Node
)
