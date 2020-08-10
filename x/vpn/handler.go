package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	ctx, write := ctx.CacheContext()
	defer write()

	node.BeginBlock(ctx, k.Node)
}

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case provider.MsgRegister:
			return provider.HandleRegister(ctx, k.Provider, msg)
		case provider.MsgUpdate:
			return provider.HandleUpdate(ctx, k.Provider, msg)

		case node.MsgRegister:
			return node.HandleRegister(ctx, k.Node, msg)
		case node.MsgUpdate:
			return node.HandleUpdate(ctx, k.Node, msg)
		case node.MsgSetStatus:
			return node.HandleSetStatus(ctx, k.Node, msg)

		case plan.MsgAdd:
			return plan.HandleAdd(ctx, k.Plan, msg)
		case plan.MsgSetStatus:
			return plan.HandleSetStatus(ctx, k.Plan, msg)
		case plan.MsgAddNode:
			return plan.HandleAddNode(ctx, k.Plan, msg)
		case plan.MsgRemoveNode:
			return plan.HandleRemoveNode(ctx, k.Plan, msg)

		case subscription.MsgSubscribeToPlan:
			return subscription.HandleSubscribeToPlan(ctx, k.Subscription, msg)
		case subscription.MsgSubscribeToNode:
			return subscription.HandleSubscribeToNode(ctx, k.Subscription, msg)
		case subscription.MsgEnd:
			return subscription.HandleEnd(ctx, k.Subscription, msg)
		case subscription.MsgAddQuota:
			return subscription.HandleAddQuota(ctx, k.Subscription, msg)
		case subscription.MsgUpdateQuota:
			return subscription.HandleUpdateQuota(ctx, k.Subscription, msg)

		case session.MsgUpdateSession:
			return session.HandleUpdateSession(ctx, k.Session, msg)
		default:
			return types.ErrorUnknownMsgType(msg.Type()).Result()
		}
	}
}
