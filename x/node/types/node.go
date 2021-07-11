package types

import (
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (n *Node) GetAddress() hubtypes.NodeAddress {
	if n.Address == "" {
		return nil
	}

	address, err := hubtypes.NodeAddressFromBech32(n.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (n *Node) GetProvider() hubtypes.ProvAddress {
	if n.Provider == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(n.Provider)
	if err != nil {
		panic(err)
	}

	return address
}

func (n *Node) Validate() error {
	if n.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(n.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", n.Address)
	}
	if n.Provider == "" && n.Price == nil {
		return fmt.Errorf("both provider and price cannot be empty")
	}
	if n.Provider != "" && n.Price != nil {
		return fmt.Errorf("either provider or price must be empty")
	}
	if n.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(n.Provider); err != nil {
			return errors.Wrapf(err, "invalid provider %s", n.Provider)
		}
	}
	if n.Price != nil {
		if n.Price.Len() == 0 {
			return fmt.Errorf("price cannot be empty")
		}
		if !n.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
	}
	if n.RemoteURL == "" {
		return fmt.Errorf("remote_url cannot be empty")
	}
	if len(n.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d", 64)
	}

	remoteURL, err := url.ParseRequestURI(n.RemoteURL)
	if err != nil {
		return errors.Wrapf(err, "invalid remote_url %s", n.RemoteURL)
	}
	if remoteURL.Scheme != "https" {
		return fmt.Errorf("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return fmt.Errorf("remote_url port cannot be empty")
	}

	if !n.Status.Equal(hubtypes.StatusActive) && !n.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("status must be either active or inactive")
	}
	if n.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

func (n *Node) PriceForDenom(s string) (sdk.Coin, bool) {
	for _, coin := range n.Price {
		if coin.Denom == s {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (n *Node) BytesForCoin(coin sdk.Coin) (sdk.Int, error) {
	price, found := n.PriceForDenom(coin.Denom)
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
