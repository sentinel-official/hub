package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/x/swap/types"
)

var (
	_ types.QueryServiceServer = &queryServer{}
)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(keeper Keeper) types.QueryServiceServer {
	return &queryServer{Keeper: keeper}
}

func (q *queryServer) QuerySwap(c context.Context, req *types.QuerySwapRequest) (*types.QuerySwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		ctx  = sdk.UnwrapSDKContext(c)
		hash = types.BytesToHash(req.TxHash)
	)

	item, found := q.GetSwap(ctx, hash)
	if !found {
		return nil, nil
	}

	return &types.QuerySwapResponse{Swap: item}, nil
}

func (q *queryServer) QuerySwaps(c context.Context, req *types.QuerySwapsRequest) (*types.QuerySwapsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Swaps
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SwapKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var item types.Swap
		if err := q.cdc.UnmarshalBinaryBare(value, &item); err != nil {
			return false, err
		}

		if accumulate {
			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySwapsResponse{Swaps: items, Pagination: pagination}, nil
}
