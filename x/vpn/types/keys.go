package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	ModuleName           = "vpn"
	StoreKeySession      = "vpnSession"
	StoreKeyNode         = "vpnNode"
	StoreKeySubscription = "vpnSubscription"
	QuerierRoute         = ModuleName
	RouterKey            = ModuleName

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeRegistered = "DE-REGISTERED"
	StatusStarted      = "STARTED"
	StatusEnded        = "ENDED"
)

var (
	NodesCountKeyPrefix   = []byte{0x00}
	NodeKeyPrefix         = []byte{0x01}
	SubscriptionKeyPrefix = []byte{0x01}
	SessionKeyPrefix      = []byte{0x01}
)

func NodesCountKey(address csdkTypes.AccAddress) []byte {
	return append(NodesCountKeyPrefix, address.Bytes()...)
}

func NodeID(address csdkTypes.Address, count uint64) []byte {
	return append(address.Bytes(), []byte(fmt.Sprintf("$%d", count))...)
}

func NodeKey(id sdkTypes.ID) []byte {
	return append(NodeKeyPrefix, id.Bytes()...)
}

func SubscriptionID(nodeID sdkTypes.ID, count uint64) []byte {
	return append(nodeID.Bytes(), []byte(fmt.Sprintf("$%d", count))...)
}

func SubscriptionKey(id sdkTypes.ID) []byte {
	return append(SubscriptionKeyPrefix, id.Bytes()...)
}

func SessionID(subscriptionID sdkTypes.ID, count uint64) []byte {
	return append(subscriptionID.Bytes(), []byte(fmt.Sprintf("$%d", count))...)
}

func SessionKey(id sdkTypes.ID) []byte {
	return append(SessionKeyPrefix, id.Bytes()...)
}

func ActiveNodeIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}

func ActiveSessionIDsKey(height int64) []byte {
	return csdkTypes.Uint64ToBigEndian(uint64(height))
}
