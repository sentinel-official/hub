package types

import (
	"fmt"
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Node struct {
	ID          uint64               `json:"id"`
	Owner       csdkTypes.AccAddress `json:"owner"`
	OwnerPubKey crypto.PubKey        `json:"owner_pub_key"`
	Deposit     csdkTypes.Coin       `json:"deposit"`

	Type          string             `json:"type"`
	Version       string             `json:"version"`
	Moniker       string             `json:"moniker"`
	PricesPerGB   csdkTypes.Coins    `json:"prices_per_gb"`
	InternetSpeed sdkTypes.Bandwidth `json:"internet_speed"`
	Encryption    string             `json:"encryption"`

	SubscriptionsCount uint64 `json:"subscriptions_count"`
	Status             string `json:"status"`
	StatusModifiedAt   int64  `json:"status_modified_at"`
}

func (n Node) String() string {
	ownerPubKey, err := csdkTypes.Bech32ifyAccPub(n.OwnerPubKey)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(`Node
  ID:                  %d
  Owner Address:       %s
  Owner Public Key:    %s
  Deposit:             %s
  Type:                %s
  Version:             %s
  Moniker:             %s
  Price Per GB:        %s
  Internet speed:      %s
  Encryption:          %s
  Subscriptions Count: %d
  Status:              %s
  Status Modified At:  %d`, n.ID, n.Owner, ownerPubKey, n.Deposit, n.Type, n.Version,
		n.Moniker, n.PricesPerGB, n.InternetSpeed, n.Encryption,
		n.SubscriptionsCount, n.Status, n.StatusModifiedAt)
}

func (n Node) UpdateInfo(_node Node) Node {
	if len(_node.Type) != 0 {
		n.Type = _node.Type
	}
	if len(_node.Version) != 0 {
		n.Version = _node.Version
	}
	if len(_node.Moniker) != 0 {
		n.Moniker = _node.Moniker
	}
	if _node.PricesPerGB != nil &&
		_node.PricesPerGB.Len() > 0 && _node.PricesPerGB.IsValid() {

		n.PricesPerGB = _node.PricesPerGB
	}
	if !_node.InternetSpeed.AnyNil() && _node.InternetSpeed.AllPositive() {
		n.InternetSpeed = _node.InternetSpeed
	}
	if len(_node.Encryption) != 0 {
		n.Encryption = _node.Encryption
	}

	return n
}

func (n Node) FindPricePerGB(denom string) (coin csdkTypes.Coin) {
	index := sort.Search(n.PricesPerGB.Len(), func(i int) bool {
		return n.PricesPerGB[i].Denom >= denom
	})

	if index == n.PricesPerGB.Len() ||
		(index < n.PricesPerGB.Len() && n.PricesPerGB[index].Denom != denom) {
		return coin
	}

	return n.PricesPerGB[index]
}

func (n Node) DepositToBandwidth(deposit csdkTypes.Coin) (bandwidth sdkTypes.Bandwidth, err csdkTypes.Error) {
	pricePerGB := n.FindPricePerGB(deposit.Denom)
	if len(pricePerGB.Denom) == 0 || pricePerGB.Amount.IsZero() {
		return bandwidth, ErrorInvalidDeposit()
	}

	gb := deposit.Amount.Mul(sdkTypes.GB).Quo(pricePerGB.Amount)
	return sdkTypes.NewBandwidth(gb, gb), nil
}

// nolint: gocyclo
func (n Node) IsValid() error {
	if n.Owner == nil || n.Owner.Empty() {
		return fmt.Errorf("invalid owner")
	}
	if n.OwnerPubKey == nil {
		return fmt.Errorf("invalid owner public key")
	}
	if len(n.Deposit.Denom) == 0 {
		return fmt.Errorf("invalid deposit")
	}
	if len(n.Type) == 0 {
		return fmt.Errorf("invalid type")
	}
	if len(n.Version) == 0 {
		return fmt.Errorf("invalid version")
	}
	if len(n.Moniker) > 128 {
		return fmt.Errorf("invalid moniker")
	}
	if n.PricesPerGB == nil || !n.PricesPerGB.IsValid() {
		return fmt.Errorf("invalid price per gb")
	}
	if n.InternetSpeed.AnyNil() || !n.InternetSpeed.AllPositive() {
		return fmt.Errorf("invalid internet speed")
	}
	if len(n.Encryption) == 0 {
		return fmt.Errorf("invalid encryption")
	}
	if n.Status != StatusRegistered && n.Status != StatusActive &&
		n.Status != StatusInactive && n.Status != StatusDeRegistered {
		return fmt.Errorf("invalid status")
	}

	return nil
}
