package ibc

import (
	"encoding/json"
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func NewVpnHandler(k vpn.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case hub.MsgCoinLocker:
				return handleSetNodeStatus(ctx, k, msg.IBCPacket)
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

func handleSetNodeStatus(ctx csdkTypes.Context, k vpn.Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {

	var Data sdkTypes.VpnDetails
	msg, _ := ibcPacket.Message.(hub.MsgCoinLocker)
	vpnId := msg.LockerId
	status := msg.Locked
	vpnData, err := k.GetVpnDetails(ctx, vpnId)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(vpnData, &Data)

	if err != nil {
		panic(err)
	}

	err = k.SetVpnStatus(ctx, vpnId, Data, status)

	if err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}
