package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewHandler(k Keeper, ik ibc.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterNode:
			return handleRegisterNode(ctx, k, ik, msg)
		case MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, k, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).Name()

			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	vpnId := msg.From.String()
	vpnDetails := k.GetVPNDetails(ctx, vpnId)

	if vpnDetails != nil {
		panic("Already registered")
	}

	k.SetVPNDetails(ctx, vpnId, msg.Details)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  "sentinel-vpn",
		DestChainId: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerId: vpnId,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k Keeper, msg MsgUpdateNodeStatus) csdkTypes.Result {
	vpnDetails := k.GetVPNDetails(ctx, msg.VPNId)

	if vpnDetails == nil {
		panic("VPN not registered")
	}

	k.SetVPNStatus(ctx, msg.VPNId, msg.Status)

	return csdkTypes.Result{}
}
