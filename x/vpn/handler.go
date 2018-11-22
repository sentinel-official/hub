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
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).String()
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	sequence, err := k.AccountKeeper.GetSequence(ctx, msg.From)

	if err != nil {
		// TODO: Replace with ErrGetSequence
		panic(err)
	}

	vpnID := msg.From.String() + "/" + strconv.Itoa(int(sequence))

	if lockerID := k.VPNStoreKey.Name() + "/" + vpnID; msg.LockerID != lockerID {
		// TODO: Replace with ErrLockerIDMismatch
		panic("msg.lockerid != lockerid")
	}

	if vpnDetails := k.GetVPNDetails(ctx, vpnID); vpnDetails != nil {
		// TODO: Replace with ErrVPNAlreadyExists
		panic("vpndetails != nil")
	}

	vpnDetails := sdkTypes.VPNDetails{
		Address:    msg.From,
		APIPort:    msg.APIPort,
		Location:   msg.Location,
		NetSpeed:   msg.NetSpeed,
		EncMethod:  msg.EncMethod,
		PricePerGB: msg.PricePerGB,
		Version:    msg.Version,
		LockerID:   msg.LockerID,
	}

	k.SetVPNDetails(ctx, vpnID, &vpnDetails)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID:  msg.LockerID,
			Coins:     msg.Coins,
			PubKey:    msg.PubKey,
			Signature: msg.Signature,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		panic(err)
	}

	// TODO: Replace with SuccessRegisterNode
	return csdkTypes.Result{}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k Keeper, msg MsgUpdateNodeStatus) csdkTypes.Result {
	if vpnDetails := k.GetVPNDetails(ctx, msg.VPNID); vpnDetails == nil {
		// TODO: Replace with ErrVPNNotExists
		panic("vpndetails == nil")
	}

	k.SetVPNStatus(ctx, msg.VPNID, msg.Status)

	// TODO: Replace with SuccessUpdateNodeStatus
	return csdkTypes.Result{}
}

func handlePayVPNService(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVPNService) csdkTypes.Result {
	sequence, err := k.AccountKeeper.GetSequence(ctx, msg.From)

	if err != nil {
		// TODO: Replace with ErrGetSequence
		panic(err)
	}

	sessionID := msg.From.String() + "/" + strconv.Itoa(int(sequence))

	if lockerID := k.SessionStoreKey.Name() + "/" + sessionID; msg.LockerID != lockerID {
		// TODO: Replace with ErrLockerIDMismatch
		panic("msg.lockerid != lockerid")
	}

	if sessionDetails := k.GetSessionDetails(ctx, sessionID); sessionDetails != nil {
		// TODO: Replace wtih ErrSessionAlreadyExists
		panic("sessiondetails != nil")
	}

	vpnDetails := k.GetVPNDetails(ctx, msg.VPNID)

	if vpnDetails == nil {
		// TODO: Replace with ErrVPNNotExists
		panic("vpndetails == nil")
	}

	sessionDetails := sdkTypes.SessionDetails{
		VPNID:         msg.VPNID,
		ClientAddress: msg.From,
		GBToProvide:   0,
		PricePerGB:    vpnDetails.PricePerGB,
	}

	k.SetSessionDetails(ctx, sessionID, &sessionDetails)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerID:  msg.LockerID,
			Coins:     msg.Coins,
			PubKey:    msg.PubKey,
			Signature: msg.Signature,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		panic(err)
	}

	// TODO: Replace with SuccessPayVPNService
	return csdkTypes.Result{}
}
