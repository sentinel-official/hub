package types

import (
	"sort"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type NodeDetails struct {
	ID           string
	Owner        csdkTypes.AccAddress
	LockedAmount csdkTypes.Coin
	APIPort      uint16
	NetSpeed     sdkTypes.Bandwidth
	EncMethod    string
	PricesPerGB  csdkTypes.Coins
	NodeType     string
	Version      string
	Status       string
	StatusAt     time.Time
	DetailsAt    time.Time
}

func (nd NodeDetails) FindPricePerGB(denom string) csdkTypes.Coin {
	index := sort.Search(nd.PricesPerGB.Len(), func(i int) bool {
		return nd.PricesPerGB[i].Denom >= denom
	})

	if index == nd.PricesPerGB.Len() {
		return csdkTypes.Coin{}
	}

	return nd.PricesPerGB[index]
}

func (nd NodeDetails) CalculateBandwidth(amount csdkTypes.Coin) (sdkTypes.Bandwidth, csdkTypes.Error) {
	pricePerGB := nd.FindPricePerGB(amount.Denom)
	if len(pricePerGB.Denom) == 0 || pricePerGB.Amount.IsZero() {
		return sdkTypes.Bandwidth{}, ErrorInvalidPriceDenom()
	}

	upload := amount.Amount.Div(pricePerGB.Amount).Mul(GB)
	download := amount.Amount.Div(pricePerGB.Amount).Mul(GB)
	bandwidth := sdkTypes.NewBandwidth(upload, download)

	return bandwidth, nil
}
