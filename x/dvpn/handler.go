package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
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
		case node.MsgUpdateNode:
			return node.HandleUpdateNode(ctx, k.Node, msg)
		case node.MsgSetNodeStatus:
			return node.HandleSetNodeStatus(ctx, k.Node, msg)
		case subscription.MsgAddPlan:
			return subscription.HandleAddPlan(ctx, k.Subscription, msg)
		case subscription.MsgSetPlanStatus:
			return subscription.HandleSetPlanStatus(ctx, k.Subscription, msg)
		case subscription.MsgAddNodeForPlan:
			return subscription.HandleAddNodeForPlan(ctx, k.Subscription, msg)
		case subscription.MsgRemoveNodeForPlan:
			return subscription.HandleRemoveNodeForPlan(ctx, k.Subscription, msg)
		case subscription.MsgStartSubscription:
			return subscription.HandleStartSubscription(ctx, k.Subscription, msg)
		case subscription.MsgAddAddressForSubscription:
			return subscription.HandleAddAddressForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgRemoveAddressForSubscription:
			return subscription.HandleRemoveAddressForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgEndSubscription:
			return subscription.HandleEndSubscription(ctx, k.Subscription, msg)
		default:
			return types.ErrorUnknownMsgType(msg.Type()).Result()
		}
	}
}
