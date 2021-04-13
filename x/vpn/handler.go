package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	var (
		providerServer     = provider.NewMsgServiceServer(k.Provider)
		nodeServer         = node.NewMsgServiceServer(k.Node)
		planServer         = plan.NewMsgServiceServer(k.Plan)
		subscriptionServer = subscription.NewMsgServiceServer(k.Subscription)
		sessionServer      = session.NewMsgServiceServer(k.Session)
	)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *provider.MsgRegisterRequest:
			res, err := providerServer.MsgRegister(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *provider.MsgUpdateRequest:
			res, err := providerServer.MsgUpdate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *node.MsgRegisterRequest:
			res, err := nodeServer.MsgRegister(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *node.MsgUpdateRequest:
			res, err := nodeServer.MsgUpdate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *node.MsgSetStatusRequest:
			res, err := nodeServer.MsgSetStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *plan.MsgAddRequest:
			res, err := planServer.MsgAdd(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plan.MsgSetStatusRequest:
			res, err := planServer.MsgSetStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plan.MsgAddNodeRequest:
			res, err := planServer.MsgAddNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plan.MsgRemoveNodeRequest:
			res, err := planServer.MsgRemoveNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *subscription.MsgSubscribeToPlanRequest:
			res, err := subscriptionServer.MsgSubscribeToPlan(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscription.MsgSubscribeToNodeRequest:
			res, err := subscriptionServer.MsgSubscribeToNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscription.MsgCancelRequest:
			res, err := subscriptionServer.MsgCancel(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscription.MsgAddQuotaRequest:
			res, err := subscriptionServer.MsgAddQuota(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscription.MsgUpdateQuotaRequest:
			res, err := subscriptionServer.MsgUpdateQuota(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *session.MsgUpsertRequest:
			res, err := sessionServer.MsgUpsert(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownMsgType, "%s", msg.Type())
		}
	}
}
