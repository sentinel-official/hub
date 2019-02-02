package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

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
