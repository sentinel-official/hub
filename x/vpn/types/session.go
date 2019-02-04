package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionID string

func (s SessionID) Bytes() []byte  { return []byte(s) }
func (s SessionID) String() string { return string(s) }
func (s SessionID) Len() int       { return len(s) }
func (s SessionID) Valid() bool {
	splits := strings.Split(s.String(), "/")
	return len(splits) == 2
}

func NewSessionID(str string) SessionID {
	return SessionID(str)
}

func SessionIDFromOwnerCount(address csdkTypes.Address, count uint64) SessionID {
	id := fmt.Sprintf("%s/%d", address.String(), count)
	return NewSessionID(id)
}

type SessionBandwidth struct {
	ToProvide     sdkTypes.Bandwidth
	Consumed      sdkTypes.Bandwidth
	NodeOwnerSign []byte
	ClientSign    []byte
	UpdatedAt     time.Time
}
type SessionDetails struct {
	ID           SessionID
	NodeID       NodeID
	Client       csdkTypes.AccAddress
	LockedAmount csdkTypes.Coin
	PricePerGB   csdkTypes.Coin
	Bandwidth    SessionBandwidth
	Status       string
	StatusAt     time.Time
	StartedAt    time.Time
	EndedAt      time.Time
}

type BandwidthSign struct {
	SessionID SessionID
	Bandwidth sdkTypes.Bandwidth
	NodeOwner csdkTypes.AccAddress
	Client    csdkTypes.AccAddress
}

func (bsd BandwidthSign) GetBytes() ([]byte, csdkTypes.Error) {
	bsdBytes, err := json.Marshal(bsd)
	if err != nil {
		return nil, ErrorMarshal()
	}

	return bsdBytes, nil
}

func NewBandwidthSign(sessionID SessionID, bandwidth sdkTypes.Bandwidth,
	nodeOwner, client csdkTypes.AccAddress) *BandwidthSign {

	return &BandwidthSign{
		SessionID: sessionID,
		Bandwidth: bandwidth,
		NodeOwner: nodeOwner,
		Client:    client,
	}
}
