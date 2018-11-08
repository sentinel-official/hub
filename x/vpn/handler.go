package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	vpnTypes "github.com/ironman0x7b2/sentinel-hub/types"
	"github.com/ironman0x7b2/sentinel-hub/x"
)

func handleRegisterVpn(ctx sdkTypes.Context, k Keeper, im ibc.Mapper, msg MsgRegisterVpn) sdkTypes.Result {
	vpnId := msg.Register.Ip + msg.Register.Port
	vpnStore := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	cdc := codec.New()
	vpnData := vpnStore.Get(vpnIdBytes)

	if vpnData != nil {
		panic("Already registered")
	}
	id, err := k.SetVpnDetails(ctx, msg.Register, vpnId)
	if err != nil {
		panic(err)
	}

	Packet := vpnTypes.VpnIBCPacket{
		VpnId:     vpnId,
		Address:   msg.From,
		Coin:      msg.Coin,
		DestChain: "Hub",
	}

	err = x.PostIBCPacket(ctx, k, im, Packet)
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
