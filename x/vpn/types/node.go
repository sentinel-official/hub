package types

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Node struct {
	ID          sdkTypes.ID          `json:"id"`
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
	if _node.PricesPerGB != nil && _node.PricesPerGB.Len() > 0 &&
		_node.PricesPerGB.IsValid() && _node.PricesPerGB.IsAllPositive() {

		n.PricesPerGB = _node.PricesPerGB
	}
	if !_node.InternetSpeed.IsNil() && _node.InternetSpeed.IsPositive() {
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
