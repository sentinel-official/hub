package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	ModuleName       = "vpn"
	StoreKeySession  = "vpn_session"
	StoreKeyNode     = "vpn_node"
	QuerierRoute     = ModuleName
	RouterKey        = ModuleName
	StatusRegister   = "REGISTER"
	StatusActive     = "ACTIVE"
	StatusInactive   = "INACTIVE"
	StatusDeregister = "DEREGISTER"
	StatusInit       = "INIT"
	StatusEnd        = "END"
)

var (
	NodeKeyPrefix               = []byte{0x01}
	NodesCountKeyPrefix         = []byte{0x02}
	ActiveNodeIDsAtHeightPrefix = []byte{0x03}

	SessionKeyPrefix               = []byte{0x01}
	SessionsCountKeyPrefix         = []byte{0x02}
	ActiveSessionIDsAtHeightPrefix = []byte{0x03}
)

func NodeKey(id sdkTypes.ID) []byte {
	return append(NodeKeyPrefix, id.Bytes()...)
}

func NodesCountKey(address csdkTypes.AccAddress) []byte {
	return append(NodesCountKeyPrefix, address.Bytes()...)
}

func ActiveNodeIDsAtHeightKey(height int64) []byte {
	return append(ActiveNodeIDsAtHeightPrefix, []byte(fmt.Sprintf("%d", height))...)
}

func SessionKey(id sdkTypes.ID) []byte {
	return append(SessionKeyPrefix, id.Bytes()...)
}

func SessionsCountKey(address csdkTypes.AccAddress) []byte {
	return append(SessionsCountKeyPrefix, address.Bytes()...)
}

func ActiveSessionIDsAtHeightKey(height int64) []byte {
	return append(ActiveSessionIDsAtHeightPrefix, []byte(fmt.Sprintf("%d", height))...)
}
