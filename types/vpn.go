package types

import (
	"fmt"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type VPNNodeDetails struct {
	Owner        csdkTypes.AccAddress
	LockedAmount csdkTypes.Coins
	APIPort      uint16
	NetSpeed     Bandwidth
	EncMethod    string
	PerGBAmount  csdkTypes.Coins
	Version      string
	Status       string
}

type VPNSessionDetails struct {
	VPNOwnerAddress    csdkTypes.AccAddress
	ClientAddress      csdkTypes.AccAddress
	PerGBAmount        csdkTypes.Coins
	BandwidthToProvide Bandwidth
	BandwidthConsumed  Bandwidth
	StartTime          *time.Time
	EndTime            *time.Time
	Status             string
}

func VPNNodesCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("vpn/nodes_count/%s", accAddress.String())
}

func VPNSessionsCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("vpn/sessions_count/%s", accAddress.String())
}

func VPNNodeKey(accAddress csdkTypes.AccAddress, count uint64) string {
	return fmt.Sprintf("%s/%d", accAddress.String(), count)
}

const (
	KeyActiveNodeIDs    = "ACTIVE_NODE_IDS"
	KeyActiveSessionIDs = "ACTIVE_SESSION_IDS"

	StoreKeyVPNSession = "vpn_session"
	StoreKeyVPNNode    = "vpn_node"

	StatusActive     = "ACTIVE"
	StatusDeregister = "DEREGISTERED"
	StatusEnd        = "ENDED"
	StatusInactive   = "INACTIVE"
	StatusRegister   = "REGISTERED"
	StatusStart      = "STARTED"
)
