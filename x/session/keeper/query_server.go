package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

var (
	_ types.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(keeper Keeper) types.QueryServiceServer {
	return &queryServer{Keeper: keeper}
}

func (q *queryServer) QuerySession(c context.Context, req *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetSession(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "session does not exist for id %d", req.Id)
	}

	return &types.QuerySessionResponse{Session: item}, nil
}

func (q *queryServer) QuerySessions(c context.Context, req *types.QuerySessionsRequest) (*types.QuerySessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SessionKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var item types.Session
		if err := q.cdc.Unmarshal(value, &item); err != nil {
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

	return &types.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySessionsForAddress(c context.Context, req *types.QuerySessionsForAddressRequest) (*types.QuerySessionsForAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items      types.Sessions
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.StatusActive) {
		store := prefix.NewStore(q.Store(ctx), types.GetActiveSessionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key []byte, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.StatusInactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactiveSessionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key []byte, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else {
		// NOTE: Do not use this; less efficient; consider using active + inactive

		store := prefix.NewStore(q.Store(ctx), types.SessionKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Session
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}
			if !strings.EqualFold(item.Address, req.Address) {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForAddressResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
