package vpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	nodekeeper "github.com/sentinel-official/hub/x/node/keeper"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plankeeper "github.com/sentinel-official/hub/x/plan/keeper"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providerkeeper "github.com/sentinel-official/hub/x/provider/keeper"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessionkeeper "github.com/sentinel-official/hub/x/session/keeper"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptionkeeper "github.com/sentinel-official/hub/x/subscription/keeper"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	var (
		providerServer     = providerkeeper.NewMsgServiceServer(k.Provider)
		nodeServer         = nodekeeper.NewMsgServiceServer(k.Node)
		planServer         = plankeeper.NewMsgServiceServer(k.Plan)
		subscriptionServer = subscriptionkeeper.NewMsgServiceServer(k.Subscription)
		sessionServer      = sessionkeeper.NewMsgServiceServer(k.Session)
	)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *providertypes.MsgRegisterRequest:
			res, err := providerServer.MsgRegister(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *providertypes.MsgUpdateRequest:
			res, err := providerServer.MsgUpdate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *nodetypes.MsgRegisterRequest:
			res, err := nodeServer.MsgRegister(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *nodetypes.MsgUpdateRequest:
			res, err := nodeServer.MsgUpdate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *nodetypes.MsgSetStatusRequest:
			res, err := nodeServer.MsgSetStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *plantypes.MsgAddRequest:
			res, err := planServer.MsgAdd(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plantypes.MsgSetStatusRequest:
			res, err := planServer.MsgSetStatus(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plantypes.MsgAddNodeRequest:
			res, err := planServer.MsgAddNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *plantypes.MsgRemoveNodeRequest:
			res, err := planServer.MsgRemoveNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *subscriptiontypes.MsgSubscribeToPlanRequest:
			res, err := subscriptionServer.MsgSubscribeToPlan(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscriptiontypes.MsgSubscribeToNodeRequest:
			res, err := subscriptionServer.MsgSubscribeToNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscriptiontypes.MsgCancelRequest:
			res, err := subscriptionServer.MsgCancel(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscriptiontypes.MsgAddQuotaRequest:
			res, err := subscriptionServer.MsgAddQuota(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *subscriptiontypes.MsgUpdateQuotaRequest:
			res, err := subscriptionServer.MsgUpdateQuota(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *sessiontypes.MsgUpsertRequest:
			res, err := sessionServer.MsgUpsert(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownMsgType, "%s", msg.Type())
		}
	}
}
