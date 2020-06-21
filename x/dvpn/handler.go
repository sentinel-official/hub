package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case provider.MsgRegisterProvider:
			return provider.HandleRegisterProvider(ctx, k.Provider, msg)
		case provider.MsgUpdateProvider:
			return provider.HandleUpdateProvider(ctx, k.Provider, msg)
		case node.MsgRegisterNode:
			return node.HandleRegisterNode(ctx, k.Node, msg)
		default:
			return types.ErrorUnknownMsgType(msg.Type()).Result()
		}
	}
}
