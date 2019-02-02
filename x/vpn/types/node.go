package types

import (
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
