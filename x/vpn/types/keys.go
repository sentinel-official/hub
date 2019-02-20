package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	NodeKeyPrefix               = []byte{0x01}
	NodesCountKeyPrefix         = []byte{0x02}
	ActiveNodeIDsAtHeightPrefix = []byte{0x03}

	SessionKeyPrefix               = []byte{0x01}
	SessionsCountKeyPrefix         = []byte{0x02}
	ActiveSessionIDsAtHeightPrefix = []byte{0x03}

	KeyActiveNodeIDs = []byte("ACTIVE_NODE_IDS")
)

func NodeKey(id sdkTypes.ID) []byte {
	return append(NodeKeyPrefix, id.Bytes()...)
}

func SessionKey(id sdkTypes.ID) []byte {
	return append(SessionKeyPrefix, id.Bytes()...)
}

func NodesCountKey(address csdkTypes.Address) []byte {
	return append(NodesCountKeyPrefix, address.Bytes()...)
}

func SessionsCountKey(address csdkTypes.Address) []byte {
	return append(SessionsCountKeyPrefix, address.Bytes()...)
}

func ActiveNodeIDsAtHeightKey(height int64) []byte {
	return append(ActiveNodeIDsAtHeightPrefix, []byte(fmt.Sprintf("%d", height))...)
}

func ActiveSessionIDsAtHeightKey(height int64) []byte {
	return append(ActiveSessionIDsAtHeightPrefix, []byte(fmt.Sprintf("%d", height))...)
}

const (
	StoreKeySession = "vpn_session"
	StoreKeyNode    = "vpn_node"

	RouterKey    = "vpn"
	QuerierRoute = "vpn"

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeregistered = "DEREGISTERED"
	StatusInit         = "INIT"
	StatusEnd          = "ENDED"
)
