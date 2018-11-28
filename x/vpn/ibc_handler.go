package vpn

import (
	"fmt"
	"reflect"
	"strings"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewIBCVPNHandler(k Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case ibc.MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case hub.MsgLockerStatus:
				newMsg, _ := msg.IBCPacket.Message.(hub.MsgLockerStatus)
				if strings.HasPrefix(newMsg.LockerID, k.VPNStoreKey.Name()+"/") {
					return handleUpdateNodeStatus(ctx, k, msg.IBCPacket)
				} else if strings.HasPrefix(newMsg.LockerID, k.SessionStoreKey.Name()+"/") {
					return handleUpdateSessionStatus(ctx, k, msg.IBCPacket)
				} else {
					errMsg := "Unrecognized locker id: " + newMsg.LockerID
					return csdkTypes.ErrUnknownRequest(errMsg).Result()
				}
			default:
				errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(ibcMsg).Name()
				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgLockerStatus)
	nodeID := msg.LockerID[len(k.VPNStoreKey.Name())+1:]

	if vpnDetails, err := k.GetVPNDetails(ctx, nodeID); true {
		if err != nil {
			return err.Result()
		}

		if vpnDetails == nil {
			return errorVPNNotExists().Result()
		}
	}

	switch msg.Status {
	case "LOCKED":
		if err := k.SetVPNStatus(ctx, nodeID, "ACTIVE"); err != nil {
			return err.Result()
		}
		if err := k.AddActiveNodeID(ctx, nodeID); err != nil {
			return err.Result()
		}
	case "RELEASED":
		if err := k.SetVPNStatus(ctx, nodeID, "DEREGISTERED"); err != nil {
			return err.Result()
		}
	default:
		return errorInvalidLockStatus().Result()
	}

	return csdkTypes.Result{}
}

func handleUpdateSessionStatus(ctx csdkTypes.Context, k Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgLockerStatus)
	sessionID := msg.LockerID[len(k.SessionStoreKey.Name())+1:]

	if sessionDetails, err := k.GetSessionDetails(ctx, sessionID); true {
		if err != nil {
			return err.Result()
		}

		if sessionDetails == nil {
			return errorSessionNotExists().Result()
		}
	}

	switch msg.Status {
	case "LOCKED":
		if err := k.SetSessionStatus(ctx, sessionID, "ACTIVE"); err != nil {
			return err.Result()
		}
		if err := k.AddActiveSessionID(ctx, sessionID); err != nil {
			return err.Result()
		}
	case "RELEASED":
		if err := k.SetVPNStatus(ctx, sessionID, "ENDED"); err != nil {
			return err.Result()
		}
	default:
		return errorInvalidLockStatus().Result()
	}

	return csdkTypes.Result{}
}
