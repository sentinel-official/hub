package types

import (
	"encoding/json"
	"fmt"
	"sort"
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

type SessionIDs []SessionID

func (s SessionIDs) Append(id ...SessionID) SessionIDs { return append(s, id...) }
func (s SessionIDs) Len() int                          { return len(s) }
func (s SessionIDs) Less(i, j int) bool                { return s[i].String() < s[j].String() }
func (s SessionIDs) Swap(i, j int)                     { s[i], s[j] = s[j], s[i] }

func (s SessionIDs) Sort() SessionIDs {
	sort.Sort(s)
	return s
}

func (s SessionIDs) Search(id SessionID) int {
	index := sort.Search(len(s), func(i int) bool {
		return s[i].String() >= id.String()
	})

	if index < s.Len() && s[index].String() != id.String() {
		return s.Len()
	}

	return index
}

func EmptySessionIDs() SessionIDs {
	return SessionIDs{}
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

func (s SessionDetails) Amount() csdkTypes.Coin {
	consumedBandwidth := s.Bandwidth.Consumed.Upload.Add(s.Bandwidth.Consumed.Download)
	amountInt := consumedBandwidth.Div(sdkTypes.GB.Add(sdkTypes.GB)).Mul(s.PricePerGB.Amount)

	amount := csdkTypes.NewCoin(s.PricePerGB.Denom, amountInt)
	if s.LockedAmount.IsLT(amount) || s.LockedAmount.IsEqual(amount) {
		return s.LockedAmount
	}

	return amount
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
