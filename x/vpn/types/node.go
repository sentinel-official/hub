package types

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type NodeDetails struct {
	ID           sdkTypes.ID
	Owner        csdkTypes.AccAddress
	PubKey       crypto.PubKey
	LockedAmount csdkTypes.Coin

	PricesPerGB csdkTypes.Coins
	NetSpeed    sdkTypes.Bandwidth
	APIPort     uint16
	Encryption  string
	NodeType    string
	Version     string

	Status          string
	StatusAtHeight  int64
	DetailsAtHeight int64
}

func (nd *NodeDetails) UpdateDetails(details NodeDetails) {
	if details.PricesPerGB != nil && details.PricesPerGB.Len() > 0 &&
		details.PricesPerGB.IsValid() && details.PricesPerGB.IsAllPositive() {
		nd.PricesPerGB = details.PricesPerGB
	}
	if !details.NetSpeed.IsNil() && details.NetSpeed.IsPositive() {
		nd.NetSpeed = details.NetSpeed
	}
	if details.APIPort != 0 {
		nd.APIPort = details.APIPort
	}
	if len(details.Encryption) != 0 {
		nd.Encryption = details.Encryption
	}
	if len(details.NodeType) != 0 {
		nd.NodeType = details.NodeType
	}
	if len(details.Version) != 0 {
		nd.Version = details.Version
	}
}

func (nd NodeDetails) FindPricePerGB(denom string) csdkTypes.Coin {
	index := sort.Search(nd.PricesPerGB.Len(), func(i int) bool {
		return nd.PricesPerGB[i].Denom >= denom
	})

	if index == nd.PricesPerGB.Len() ||
		(index < nd.PricesPerGB.Len() && nd.PricesPerGB[index].Denom != denom) {
		return csdkTypes.Coin{}
	}

	return nd.PricesPerGB[index]
}

func (nd NodeDetails) CalculateBandwidth(amount csdkTypes.Coin) (sdkTypes.Bandwidth, csdkTypes.Error) {
	pricePerGB := nd.FindPricePerGB(amount.Denom)
	if len(pricePerGB.Denom) == 0 || pricePerGB.Amount.IsZero() {
		return sdkTypes.Bandwidth{}, ErrorInvalidPriceDenom()
	}

	upload := amount.Amount.Quo(pricePerGB.Amount).Mul(sdkTypes.GB)
	download := amount.Amount.Quo(pricePerGB.Amount).Mul(sdkTypes.GB)
	bandwidth := sdkTypes.NewBandwidth(upload, download)

	return bandwidth, nil
}
