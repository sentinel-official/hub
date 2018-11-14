package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"crypto/md5"
	"strconv"
	"encoding/hex"
)

func NewHandler(k Keeper, ik ibc.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterNode:
			return handleRegisterNode(ctx, k, ik, msg)
		case MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, k, msg)
		case MsgPayVpnService:
			return handlePayVpnService(ctx, k,ik, msg)
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

func handlePayVpnService(ctx csdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgPayVpnService) csdkTypes.Result  {
	var err error
	hash := md5.New()

	sequence, err := k.Account.GetSequence(ctx, msg.From)

	if err != nil {
		panic(err)
	}

	addressbytes := []byte(msg.From.String() + "" + strconv.Itoa(int(sequence)))
	hash.Write(addressbytes)

	if err != nil {
		panic(err)
	}

	sessionKey := hex.EncodeToString(hash.Sum(nil))[:20]

	vpnpub, err := k.Account.GetPubKey(ctx, msg.Vpnaddr)

	if err != nil {
		panic(err)
	}

	time := ctx.BlockHeader().Time

	session := sdkTypes.GetNewSessionMap(msg.Coins, vpnpub, msg.Pubkey, msg.From, time)

	k.SetSessionDetails(ctx, session, sessionKey)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  "sentinel-vpn",
		DestChainId: "sentinel-hub",
		Message: hub.MsgLockCoins{
			LockerId: sessionKey,
			Address:  msg.From,
			Coins:    msg.Coins,
		},
	}

	if err := ik.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}

}