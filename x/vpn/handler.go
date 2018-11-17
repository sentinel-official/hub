package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"strconv"
)

func NewHandler(k Keeper, ik ibc.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterNode:
			return handleRegisterNode(ctx, k, ik, msg)
		case MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, k, msg)
		case MsgPayVpnService:
			return handlePayVpnService(ctx, k, ik, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).Name()

			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	sequence, err := k.Account.GetSequence(ctx, msg.From)
	if err != nil {
		panic(err)
	}
	vpnId := msg.From.String() + "" + strconv.Itoa(int(sequence))

	vpnDetails := k.GetVPNDetails(ctx, vpnId)

	if vpnDetails != nil {
		panic("Already registered")
	}

	k.SetVPNDetails(ctx, vpnId, msg.Details)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  "sentinel-vpn",
		DestChainId: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerId: "vpn/" + vpnId,
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

func handlePayVpnService(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVpnService) csdkTypes.Result {
	sequence, err := k.Account.GetSequence(ctx, msg.From)

	if err != nil {
		panic(err)
	}

	sessionKey := msg.From.String() + "" + strconv.Itoa(int(sequence))
	vpnDetails := k.GetVPNDetails(ctx, msg.VpnId)

	session := sdkTypes.GetNewSessionMap(msg.VpnId, msg.From, vpnDetails.PricePerGb, vpnDetails.PricePerGb,
		vpnDetails.NetSpeed.Upload, vpnDetails.NetSpeed.Download)

	k.SetSessionDetails(ctx, session, sessionKey)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  "sentinel-vpn",
		DestChainId: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerId: "session/" + sessionKey,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}

}
