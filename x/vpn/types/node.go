package types

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Node struct {
	ID      hub.NodeID     `json:"id"`
	Owner   sdk.AccAddress `json:"owner"`
	Deposit sdk.Coin       `json:"deposit"`

	Type          string        `json:"type"`
	Version       string        `json:"version"`
	Moniker       string        `json:"moniker"`
	PricesPerGB   sdk.Coins     `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth `json:"internet_speed"`
	Encryption    string        `json:"encryption"`

	Status           string `json:"status"`
	StatusModifiedAt int64  `json:"status_modified_at"`
}

func (n Node) String() string {
	return fmt.Sprintf(`Node
  ID:                  %s
  Owner Address:       %s
  Deposit:             %s
  Type:                %s
  Version:             %s
  Moniker:             %s
  Price Per GB:        %s
  Internet Speed:      %s
  Encryption:          %s
  Status:              %s
  Status Modified At:  %d`, n.ID, n.Owner, n.Deposit, n.Type, n.Version,
		n.Moniker, n.PricesPerGB, n.InternetSpeed, n.Encryption,
		n.Status, n.StatusModifiedAt)
}

func (n Node) UpdateInfo(_node Node) Node {
	if _node.Type != "" {
		n.Type = _node.Type
	}
	if _node.Version != "" {
		n.Version = _node.Version
	}
	if _node.Moniker != "" {
		n.Moniker = _node.Moniker
	}
	if _node.PricesPerGB != nil &&
		_node.PricesPerGB.Len() > 0 && _node.PricesPerGB.IsValid() {
		n.PricesPerGB = _node.PricesPerGB
	}
	if !_node.InternetSpeed.AnyNil() && _node.InternetSpeed.AllPositive() {
		n.InternetSpeed = _node.InternetSpeed
	}
	if _node.Encryption != "" {
		n.Encryption = _node.Encryption
	}

	return n
}

func (n Node) FindPricePerGB(denom string) (coin sdk.Coin) {
	index := sort.Search(n.PricesPerGB.Len(), func(i int) bool {
		return n.PricesPerGB[i].Denom >= denom
	})

	if index == n.PricesPerGB.Len() ||
		(index < n.PricesPerGB.Len() && n.PricesPerGB[index].Denom != denom) {
		return coin
	}

	return n.PricesPerGB[index]
}

func (n Node) DepositToBandwidth(deposit sdk.Coin) (bandwidth hub.Bandwidth, err sdk.Error) {
	pricePerGB := n.FindPricePerGB(deposit.Denom)
	if pricePerGB.Denom == "" || pricePerGB.Amount.IsNegative() {
		return bandwidth, ErrorInvalidDeposit()
	}

	x := hub.MB100.Quo(sdk.NewInt(2))
	if !pricePerGB.Amount.IsZero() {
		x = deposit.Amount.Mul(hub.MB500).Quo(pricePerGB.Amount)
	}

	return hub.NewBandwidth(x, x), nil
}

func (n Node) IsValid() error {
	if n.Owner == nil || n.Owner.Empty() {
		return fmt.Errorf("invalid owner")
	}
	if n.Deposit.Denom == "" {
		return fmt.Errorf("invalid deposit")
	}
	if n.Type == "" || len(n.Type) < 4 || len(n.Type) > 16 {
		return fmt.Errorf("invalid type")
	}

	if n.Version == "" || len(n.Version) < 4 || len(n.Version) > 16 {
		return fmt.Errorf("invalid version")
	}
	if n.Moniker == "" || len(n.Moniker) < 4 || len(n.Moniker) > 32 {
		return fmt.Errorf("invalid moniker")
	}
	if n.PricesPerGB == nil || !n.PricesPerGB.IsValid() {
		return fmt.Errorf("invalid price per gb")
	}
	if n.InternetSpeed.AnyNil() || !n.InternetSpeed.AllPositive() {
		return fmt.Errorf("invalid internet speed")
	}

	if n.Encryption == "" || len(n.Encryption) < 4 || len(n.Encryption) > 16 {
		return fmt.Errorf("invalid encryption")
	}

	if n.Status != StatusRegistered &&
		n.Status != StatusDeRegistered {
		return fmt.Errorf("invalid status")
	}

	return nil
}
