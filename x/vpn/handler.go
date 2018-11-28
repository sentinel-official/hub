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
		case MsgPayVPNService:
			return handleSessionPayment(ctx, k, ik, msg)
		case MsgDeregisterNode:
			return handleDeregisterNode(ctx, k, ik, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).String()
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	vpnsCount, err := k.GetVPNsCount(ctx)

	if err != nil {
		return err.Result()
	}

	vpnID := msg.From.String() + "/" + strconv.Itoa(int(vpnsCount))

	if lockerID := k.VPNStoreKey.Name() + "/" + vpnID; msg.LockerID != lockerID {
		return errorLockerIDMismatch().Result()
	}

	if vpnDetails, err := k.GetVPNDetails(ctx, vpnID); true {
		if err != nil {
			return err.Result()
		}

		if vpnDetails != nil {
			return errorVPNAlreadyExists().Result()
		}
	}

	vpnDetails := sdkTypes.VPNDetails{
		Address:    msg.From,
		APIPort:    msg.APIPort,
		Location:   msg.Location,
		NetSpeed:   msg.NetSpeed,
		EncMethod:  msg.EncMethod,
		PricePerGB: msg.PricePerGB,
		Version:    msg.Version,
		Status:     "REGISTERED",
		LockerID:   msg.LockerID,
	}

	if err := k.AddVPN(ctx, vpnID, &vpnDetails); err != nil {
		return err.Result()
	}

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
		return err.Result()
	}

	return csdkTypes.Result{}
}

func handleSessionPayment(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVPNService) csdkTypes.Result {
	sessionsCount, err := k.GetSessionsCount(ctx)

	if err != nil {
		return err.Result()
	}

	sessionID := msg.From.String() + "/" + strconv.Itoa(int(sessionsCount))

	if lockerID := k.SessionStoreKey.Name() + "/" + sessionID; msg.LockerID != lockerID {
		return errorLockerIDMismatch().Result()
	}

	if sessionDetails, err := k.GetSessionDetails(ctx, sessionID); true {
		if err != nil {
			return err.Result()
		}

		if sessionDetails != nil {
			return errorSessionAlreadyExists().Result()
		}
	}

	vpnDetails, err := k.GetVPNDetails(ctx, msg.VPNID)

	if err != nil {
		return err.Result()
	}

	if vpnDetails == nil {
		return errorVPNNotExists().Result()
	}

	sessionDetails := sdkTypes.SessionDetails{
		VPNID:         msg.VPNID,
		ClientAddress: msg.From,
		GBToProvide:   0,
		PricePerGB:    vpnDetails.PricePerGB,
		Status:        "STARTED",
	}

	if err := k.AddSession(ctx, sessionID, &sessionDetails); err != nil {
		return err.Result()
	}

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
		return err.Result()
	}

	return csdkTypes.Result{}
}

func handleDeregisterNode(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgDeregisterNode) csdkTypes.Result {
	vpnDetails, err := k.GetVPNDetails(ctx, msg.VPNID)

	if err != nil {
		return err.Result()
	}

	if vpnDetails == nil {
		return errorVPNNotExists().Result()
	}

	if !msg.From.Equals(vpnDetails.Address) {
		return errorInvalidNodeOwnerAddress().Result()
	}

	if msg.LockerID != vpnDetails.LockerID {
		return errorLockerIDMismatch().Result()
	}

	if err := k.SetVPNStatus(ctx, msg.VPNID, "INACTIVE"); err != nil {
		return err.Result()
	}

	if err := k.RemoveActiveNodeID(ctx, msg.VPNID); err != nil {
		return err.Result()
	}

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgReleaseCoins{
			LockerID:  msg.LockerID,
			PubKey:    msg.PubKey,
			Signature: msg.Signature,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		return err.Result()
	}
	
	return csdkTypes.Result{}
}
