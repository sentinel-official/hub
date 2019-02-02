package types

import (
	"encoding/json"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionBandwidth struct {
	ToProvide     sdkTypes.Bandwidth
	Consumed      sdkTypes.Bandwidth
	NodeOwnerSign []byte
	ClientSign    []byte
	UpdatedAt     time.Time
}

type SessionDetails struct {
	ID           string
	NodeID       string
	Client       csdkTypes.AccAddress
	LockedAmount csdkTypes.Coin
	PricePerGB   csdkTypes.Coin
	Bandwidth    SessionBandwidth
	Status       string
	StatusAt     time.Time
	StartedAt    time.Time
	EndedAt      time.Time
}

type BandwidthSignData struct {
	SessionID string
	Bandwidth sdkTypes.Bandwidth
	NodeOwner csdkTypes.AccAddress
	Client    csdkTypes.AccAddress
}

func (bsd BandwidthSignData) GetBytes() ([]byte, csdkTypes.Error) {
	bsdBytes, err := json.Marshal(bsd)
	if err != nil {
		return nil, ErrorMarshal()
	}

	return bsdBytes, nil
}

func NewBandwidthSignData(sessionID string, bandwidth sdkTypes.Bandwidth,
	nodeOwner, client csdkTypes.AccAddress) BandwidthSignData {

	return BandwidthSignData{
		SessionID: sessionID,
		Bandwidth: bandwidth,
		NodeOwner: nodeOwner,
		Client:    client,
	}
}
