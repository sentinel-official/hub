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

	Moniker          string             `json:"moniker"`
	PricesPerGB      csdkTypes.Coins    `json:"prices_per_gb"`
	InternetSpeed    sdkTypes.Bandwidth `json:"internet_speed"`
	EncryptionMethod string             `json:"encryption_method"`
	Type             string             `json:"type"`
	Version          string             `json:"version"`

	Status                 string `json:"status"`
	StatusModifiedAtHeight int64  `json:"status_modified_at_height"`
}

func (n *Node) UpdateDetails(_node Node) {
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
	if len(_node.EncryptionMethod) != 0 {
		n.EncryptionMethod = _node.EncryptionMethod
	}
	if len(_node.Type) != 0 {
		n.Type = _node.Type
	}
	if len(_node.Version) != 0 {
		n.Version = _node.Version
	}
}

func (n Node) FindPricePerGB(denom string) csdkTypes.Coin {
	index := sort.Search(n.PricesPerGB.Len(), func(i int) bool {
		return n.PricesPerGB[i].Denom >= denom
	})

	if index == n.PricesPerGB.Len() ||
		(index < n.PricesPerGB.Len() && n.PricesPerGB[index].Denom != denom) {
		return csdkTypes.Coin{}
	}

	return n.PricesPerGB[index]
}

func (n Node) AmountToBandwidth(amount csdkTypes.Coin) (sdkTypes.Bandwidth, csdkTypes.Error) {
	pricePerGB := n.FindPricePerGB(amount.Denom)
	if len(pricePerGB.Denom) == 0 || pricePerGB.Amount.IsZero() {
		return sdkTypes.Bandwidth{}, ErrorInvalidPriceDenom()
	}

	upload := amount.Amount.Mul(sdkTypes.GB).Quo(pricePerGB.Amount)
	download := amount.Amount.Mul(sdkTypes.GB).Quo(pricePerGB.Amount)
	bandwidth := sdkTypes.NewBandwidth(upload, download)

	return bandwidth, nil
}
