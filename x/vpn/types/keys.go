package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	NodeKeyPrefix       = []byte{0x01}
	NodesCountKeyPrefix = []byte{0x02}

	SessionKeyPrefix               = []byte{0x01}
	SessionsCountKeyPrefix         = []byte{0x02}
	NodesActiveSessionIDsKeyPrefix = []byte{0x03}

	KeyActiveNodeIDs = []byte("ACTIVE_NODE_IDS")
)

func NodeKey(id NodeID) []byte {
	return append(NodeKeyPrefix, id.Bytes()...)
}

func SessionKey(id SessionID) []byte {
	return append(SessionKeyPrefix, id.Bytes()...)
}

func NodesCountKey(address csdkTypes.Address) []byte {
	return append(NodesCountKeyPrefix, address.Bytes()...)
}

func SessionsCountKey(address csdkTypes.Address) []byte {
	return append(SessionsCountKeyPrefix, address.Bytes()...)
}

func NodesActiveSessionIDsKey(id NodeID) []byte {
	return append(NodesActiveSessionIDsKeyPrefix, id.Bytes()...)
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
