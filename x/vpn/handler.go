package vpn

import (
	"reflect"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
	"github.com/ironman0x7b2/sentinel-hub/x/ibc"
	"encoding/json"
)

func NewHandler(k Keeper, im ibc.Keeper) sdkTypes.Handler {

	return func(ctx sdkTypes.Context, msg sdkTypes.Msg) sdkTypes.Result {

		switch msg := msg.(type) {

		case MsgRegisterVpn:
			return handleRegisterVpn(ctx, k, im, msg)
		case MsgAliveNode:
			return handleAliveNode(ctx, k, msg)
		default:
			errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(msg).Name()
			return sdkTypes.ErrUnknownRequest(errMsg).Result()

		}

	}

}

func handleRegisterVpn(ctx sdkTypes.Context, k Keeper, ik ibc.Keeper, msg MsgRegisterVpn) sdkTypes.Result {

	vpnId := msg.From
	cdc := codec.New()
	vpnData, err := k.GetVpnDetails(ctx, vpnId)

	if err != nil {
		panic(err)
	}

	if vpnData != nil {
		panic("Already registered")
	}

	err = k.SetVpnDetails(ctx, vpnId, msg.Details)

	if err != nil {
		panic(err)
	}

	ibcPacket := hubTypes.IBCPacket{
		SrcChainId:  "vpn",
		DestChainId: "Sentinel-hub",
		Message: hubTypes.IBCMsgRegisterVpn{
			VpnId:   vpnId,
			Address: msg.From,
			Coins:   msg.Coins,
		},
	}

	err = ik.PostIBCPacket(ctx, ibcPacket)

	if err != nil {
		panic(err)
	}

	tags := sdkTypes.NewTags("Registered Vpn address:", []byte(msg.From.String()))
	data, _ := cdc.MarshalJSON(msg)

	return sdkTypes.Result{
		Tags: tags,
		Data: data,
	}
}

func handleAliveNode(ctx sdkTypes.Context, k Keeper, msg MsgAliveNode) sdkTypes.Result {
	var Data hubTypes.VpnDetails

	vpnId := msg.From
	vpnData, err := k.GetVpnDetails(ctx, vpnId)

	if err != nil {
		panic(err)
	}

	if vpnData != nil {
		panic("Already registered")
	}

	err = json.Unmarshal(vpnData, &Data)

	if err != nil {
		panic(err)
	}

	err = k.SetAliveNode(ctx, vpnId, Data)

	if err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}
