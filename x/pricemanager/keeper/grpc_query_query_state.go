package keeper

import (
	"context"

	"github.com/sentinel-official/hub/v12/x/pricemanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) QueryState(goCtx context.Context, req *types.QueryQueryStateRequest) (*types.QueryQueryStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	qreq, err := k.GetQueryRequest(ctx, req.Sequence)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var anyQResp *cdctypes.Any
	qresp, err := k.GetQueryResponse(ctx, req.Sequence)
	if err == nil {
		anyQResp, err = cdctypes.NewAnyWithValue(&qresp)
		if err != nil {
			panic(err)
		}
	}

	anyQReq, err := cdctypes.NewAnyWithValue(&qreq)
	if err != nil {
		panic(err)
	}
	return &types.QueryQueryStateResponse{
		Request:  *anyQReq,
		Response: anyQResp,
	}, nil
}
