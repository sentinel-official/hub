package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	if _, err := hubtypes.NodeAddressFromBech32(n.Address); err != nil {
		return err
	}
	if (n.Provider != "" && n.Price != nil) ||
		(n.Provider == "" && n.Price == nil) {
		return fmt.Errorf("invalid provider and price combination; expected one of them empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(n.Provider); err != nil {
		return err
	}
	if n.Price != nil && !n.Price.IsValid() {
		return fmt.Errorf("invalid price; expected non-nil and valid value")
	}
	if len(n.RemoteURL) == 0 || len(n.RemoteURL) > 64 {
		return fmt.Errorf("invalid remote url length; expected length is between 1 and 64")
	}
	if !n.Status.Equal(hubtypes.StatusActive) && !n.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("invalid status; exptected active or inactive")
	}
	if n.StatusAt.IsZero() {
		return fmt.Errorf("invalid status at; expected non-zero value")
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

type Nodes []Node
