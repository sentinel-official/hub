package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/plan"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/session"
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
		case plan.MsgAddPlan:
			return plan.HandleAddPlan(ctx, k.Plan, msg)
		case plan.MsgSetPlanStatus:
			return plan.HandleSetPlanStatus(ctx, k.Plan, msg)
		case plan.MsgAddNodeForPlan:
			return plan.HandleAddNodeForPlan(ctx, k.Plan, msg)
		case plan.MsgRemoveNodeForPlan:
			return plan.HandleRemoveNodeForPlan(ctx, k.Plan, msg)
		case subscription.MsgStartSubscription:
			return subscription.HandleStartSubscription(ctx, k.Subscription, msg)
		case subscription.MsgAddMemberForSubscription:
			return subscription.HandleAddMemberForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgRemoveMemberForSubscription:
			return subscription.HandleRemoveMemberForSubscription(ctx, k.Subscription, msg)
		case subscription.MsgEndSubscription:
			return subscription.HandleEndSubscription(ctx, k.Subscription, msg)
		case session.MsgUpdateSession:
			return session.HandleUpdateSession(ctx, k.Session, msg)
		default:
			return types.ErrorUnknownMsgType(msg.Type()).Result()
		}
	}
}
