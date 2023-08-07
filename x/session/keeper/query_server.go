package keeper

import (
	"context"
	"fmt"

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

	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item types.Session
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySessionsForAccount(c context.Context, req *types.QuerySessionsForAccountRequest) (*types.QuerySessionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForAccountKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForAccountResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySessionsForNode(c context.Context, req *types.QuerySessionsForNodeRequest) (*types.QuerySessionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := hubtypes.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForNodeKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForNodeResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySessionsForSubscription(c context.Context, req *types.QuerySessionsForSubscriptionRequest) (*types.QuerySessionsForSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForSubscriptionResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySessionsForAllocation(c context.Context, req *types.QuerySessionsForAllocationRequest) (*types.QuerySessionsForAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForAllocationKeyPrefix(req.Id, addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForAllocationResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
