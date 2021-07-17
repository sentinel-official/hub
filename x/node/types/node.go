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

	address, err := hubtypes.NodeAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Node) GetProvider() hubtypes.ProvAddress {
	if m.Provider == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(m.Provider)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Node) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Provider == "" && m.Price == nil {
		return fmt.Errorf("both provider and price cannot be empty")
	}
	if m.Provider != "" && m.Price != nil {
		return fmt.Errorf("either provider or price must be empty")
	}
	if m.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
			return errors.Wrapf(err, "invalid provider %s", m.Provider)
		}
	}
	if m.Price != nil {
		if m.Price.Len() == 0 {
			return fmt.Errorf("price cannot be empty")
		}
		if !m.Price.IsValid() {
			return fmt.Errorf("price must be valid")
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

	if !m.Status.Equal(hubtypes.StatusActive) && !m.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("status must be either active or inactive")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

func (m *Node) PriceForDenom(s string) (sdk.Coin, bool) {
	for _, coin := range m.Price {
		if coin.Denom == s {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (m *Node) BytesForCoin(coin sdk.Coin) (sdk.Int, error) {
	price, found := m.PriceForDenom(coin.Denom)
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
