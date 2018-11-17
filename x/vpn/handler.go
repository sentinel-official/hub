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
		case MsgPayVPNService:
			return handlePayVPNService(ctx, k, ik, msg)
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
	vpnID := msg.From.String() + "" + strconv.Itoa(int(sequence))

	vpnDetails := k.GetVPNDetails(ctx, vpnID)

	if vpnDetails != nil {
		panic("Already registered")
	}

	k.SetVPNDetails(ctx, vpnID, msg.Details)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID: "vpn/" + vpnID,
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
	vpnDetails := k.GetVPNDetails(ctx, msg.VPNID)

	if vpnDetails == nil {
		panic("VPN not registered")
	}

	k.SetVPNStatus(ctx, msg.VPNID, msg.Status)

	return csdkTypes.Result{}
}

func handlePayVPNService(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVPNService) csdkTypes.Result {
	sequence, err := k.Account.GetSequence(ctx, msg.From)

	if err != nil {
		panic(err)
	}

	sessionKey := msg.From.String() + "" + strconv.Itoa(int(sequence))
	vpnDetails := k.GetVPNDetails(ctx, msg.VPNID)

	session := sdkTypes.GetNewSessionMap(msg.VPNID, msg.From, vpnDetails.PricePerGb, vpnDetails.PricePerGb,
		vpnDetails.NetSpeed.Upload, vpnDetails.NetSpeed.Download)

	k.SetSessionDetails(ctx, session, sessionKey)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID: "session/" + sessionKey,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}

}
