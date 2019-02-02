package vpn

import (
	"fmt"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/types"
)

type NodeDetails struct {
	ID           string
	Owner        csdkTypes.AccAddress
	LockedAmount csdkTypes.Coin
	APIPort      uint16
	NetSpeed     types.Bandwidth
	EncMethod    string
	PricesPerGB  csdkTypes.Coins
	NodeType     string
	Version      string
	Status       string
	StatusAt     time.Time
	DetailsAt    time.Time
}

type SessionDetails struct {
	ID            string
	NodeID        string
	ClientAddress csdkTypes.AccAddress
	LockedAmount  csdkTypes.Coin
	PricePerGB    csdkTypes.Coin
	Bandwidth     struct {
		ToProvide     types.Bandwidth
		Consumed      types.Bandwidth
		NodeOwnerSign []byte
		ClientSign    []byte
		UpdatedAt     time.Time
	}
	Status    string
	StatusAt  time.Time
	StartedAt time.Time
	EndedAt   time.Time
}

var (
	NodesCountKeyPrefix    = []byte("NODES_COUNT")
	SessionsCountKeyPrefix = []byte("SESSION_COUNT")
)

func NodesCountKey(accAddress csdkTypes.AccAddress) []byte {
	return append(NodesCountKeyPrefix, accAddress.Bytes()...)
}

func SessionsCountKey(accAddress csdkTypes.AccAddress) []byte {
	return append(SessionsCountKeyPrefix, accAddress.Bytes()...)
}

func NodeKey(accAddress csdkTypes.AccAddress, count uint64) string {
	return fmt.Sprintf("%s/%d", accAddress.String(), count)
}

func SessionKey(accAddress csdkTypes.AccAddress, count uint64) string {
	return fmt.Sprintf("%s/%d", accAddress.String(), count)
}

const (
	KeyActiveNodeIDs    = "ACTIVE_NODE_IDS"
	KeyActiveSessionIDs = "ACTIVE_SESSION_IDS"

	StoreKeySession = "vpn_session"
	StoreKeyNode    = "vpn_node"

	RouterKey    = "vpn"
	QuerierRoute = "vpn"

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeregistered = "DEREGISTERED"
	StatusStart        = "STARTED"
	StatusEnd          = "ENDED"
)
