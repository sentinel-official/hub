package keeper

import (
	"context"

	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	return &types.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgShare(c context.Context, msg *types.MsgShareRequest) (*types.MsgShareResponse, error) {
	return &types.MsgShareResponse{}, nil
}

func (k *msgServer) MsgUpdateQuota(c context.Context, msg *types.MsgUpdateQuotaRequest) (*types.MsgUpdateQuotaResponse, error) {
	return &types.MsgUpdateQuotaResponse{}, nil
}
