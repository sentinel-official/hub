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
				}
				if strings.HasPrefix(newMsg.LockerID, k.SessionStoreKey.Name()+"/") {
					return handleSetSessionStatus(ctx, k, msg.IBCPacket)
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
	vpnID := msg.LockerID[len(k.VPNStoreKey.Name())+1:]

	if vpnDeatils := k.GetVPNDetails(ctx, vpnID); vpnDeatils == nil {
		// TODO: Replace with ErrVPNNotExists
		panic("vpndetails == nil")
	}

	switch msg.Status {
	case "LOCKED":
		k.SetVPNStatus(ctx, vpnID, "ACTIVE")
	case "RELEASED":
		k.SetVPNStatus(ctx, vpnID, "DEREGISTERED")
	default:
		// TODO: Replace with ErrInvalidLockStatus
		panic("invalid locker id status")
	}

	// TODO: Replace with SuccessSetNodeStatus
	return csdkTypes.Result{}
}

func handleSetSessionStatus(ctx csdkTypes.Context, k Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgLockerStatus)
	sessionID := msg.LockerID
	status := msg.Status == "LOCKED"

	sessionDetails := k.GetSessionDetails(ctx, sessionID)

	if sessionDetails == nil {
		panic("sessiondetails == nil")
	}

	k.SetSessionStatus(ctx, sessionID, status)

	return csdkTypes.Result{}
}
