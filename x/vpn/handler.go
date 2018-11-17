package vpn

import (
	"reflect"

	"strconv"

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
		case MsgPayVPNService:
			return handlePayVPNService(ctx, k, ik, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).Name()
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	sequence, err := k.AccountKeeper.GetSequence(ctx, msg.From)

	if err != nil {
		// TODO: Replace with ErrGetSequence
		return csdkTypes.Result{}
	}

	vpnID := msg.From.String() + "/" + strconv.Itoa(int(sequence))

	if vpnDetails := k.GetVPNDetails(ctx, vpnID); vpnDetails != nil {
		// TODO: Replace with ErrVPNAlreadyExists
		return csdkTypes.Result{}
	}

	k.SetVPNDetails(ctx, vpnID, &msg.Details)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID: k.VPNStoreKey.String() + "/" + vpnID,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		return csdkTypes.Result{}
	}

	// TODO: Replace with SuccessRegisterNode
	return csdkTypes.Result{}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k Keeper, msg MsgUpdateNodeStatus) csdkTypes.Result {
	if vpnDetails := k.GetVPNDetails(ctx, msg.VPNID); vpnDetails == nil {
		// TODO: Replace with ErrVPNNotExists
		return csdkTypes.Result{}
	}

	k.SetVPNStatus(ctx, msg.VPNID, msg.Status)

	return csdkTypes.Result{}
}

func handlePayVPNService(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVPNService) csdkTypes.Result {
	sequence, err := k.AccountKeeper.GetSequence(ctx, msg.From)

	if err != nil {
		// TODO: Replace with ErrGetSequence
		return csdkTypes.Result{}
	}

	sessionID := msg.From.String() + "/" + strconv.Itoa(int(sequence))

	if session := k.GetSessionDetails(ctx, sessionID); session != nil {
		// TODO: Replace wtih ErrSessionAlreadyExists
		return csdkTypes.Result{}
	}

	vpnDetails := k.GetVPNDetails(ctx, msg.VPNID)

	if vpnDetails == nil {
		// TODO: Replace with ErrVPNNotExists
		return csdkTypes.Result{}
	}

	session := sdkTypes.GetNewSession(msg.VPNID, msg.From, vpnDetails.PricePerGb, vpnDetails.PricePerGb)
	k.SetSessionDetails(ctx, sessionID, &session)
	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID: k.SessionStoreKey.String() + "/" + sessionID,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		return csdkTypes.Result{}
	}

	// TODO: Replace with SuccessRegisterNode
	return csdkTypes.Result{}
}
