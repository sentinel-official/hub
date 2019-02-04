package types

import (
	"fmt"
	"sort"
	"strings"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type NodeID string

func (n NodeID) Bytes() []byte  { return []byte(n) }
func (n NodeID) String() string { return string(n) }
func (n NodeID) Len() int       { return len(n) }
func (n NodeID) Valid() bool {
	splits := strings.Split(n.String(), "/")
	return len(splits) == 2
}

func NewNodeID(str string) NodeID {
	return NodeID(str)
}

func NodeIDFromOwnerCount(address csdkTypes.Address, count uint64) NodeID {
	id := fmt.Sprintf("%s/%d", address.String(), count)
	return NewNodeID(id)
}

type APIPort uint32

func (a APIPort) IsZero() bool { return a == 0 }
func (a APIPort) Valid() bool  { return a > 0 && a <= 65535 }

func NewAPIPort(num uint32) APIPort {
	return APIPort(num)
}

type NodeDetails struct {
	ID           NodeID
	Owner        csdkTypes.AccAddress
	LockedAmount csdkTypes.Coin
	PricesPerGB  csdkTypes.Coins
	NetSpeed     sdkTypes.Bandwidth
	APIPort      APIPort
	EncMethod    string
	NodeType     string
	Version      string
	Status       string
	StatusAt     time.Time
	DetailsAt    time.Time
}

func (nd *NodeDetails) UpdateDetails(details NodeDetails) {
	if details.ID.Len() != 0 {
		nd.ID = details.ID
	}
	if details.Owner != nil && !details.Owner.Empty() {
		nd.Owner = details.Owner
	}
	if len(details.LockedAmount.Denom) != 0 && !details.LockedAmount.IsZero() {
		nd.LockedAmount = details.LockedAmount
	}
	if details.PricesPerGB != nil && details.PricesPerGB.Len() > 0 {
		nd.PricesPerGB = details.PricesPerGB
	}
	if details.NetSpeed.IsPositive() {
		nd.NetSpeed = details.NetSpeed
	}
	if details.APIPort.IsZero() {
		nd.APIPort = details.APIPort
	}
	if len(details.EncMethod) != 0 {
		nd.EncMethod = details.EncMethod
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
