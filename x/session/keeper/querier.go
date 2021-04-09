package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

type Querier struct {
	Keeper
}

func (q *Querier) QuerySession(c context.Context, req *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetSession(ctx, req.Id)
	if !found {
		return nil, nil
	}

	return &types.QuerySessionResponse{Session: item}, nil
}

func (q *Querier) QuerySessions(c context.Context, req *types.QuerySessionsRequest) (*types.QuerySessionsResponse, error) {
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

	return &types.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySessionsForSubscription(c context.Context, req *types.QuerySessionsForSubscriptionRequest) (*types.QuerySessionsForSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetSession(ctx, types.IDFromSessionForSubscriptionKey(key))
		if !found {
			return false, nil
		}

		if accumulate {
			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForSubscriptionResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySessionsForNode(c context.Context, req *types.QuerySessionsForNodeRequest) (*types.QuerySessionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hub.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForNodeKeyPrefix(address))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetSession(ctx, types.IDFromSessionForNodeKey(key))
		if !found {
			return false, nil
		}

		if accumulate {
			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForNodeResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySessionsForAddress(c context.Context, req *types.QuerySessionsForAddressRequest) (*types.QuerySessionsForAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSessionForAddressKeyPrefix(address))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetSession(ctx, types.IDFromSessionForAddressKey(key))
		if !found {
			return false, nil
		}

		if accumulate {
			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySessionsForAddressResponse{Sessions: items, Pagination: pagination}, nil
}

func (q *Querier) QueryActiveSession(c context.Context, req *types.QueryActiveSessionRequest) (*types.QueryActiveSessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	node, err := hub.NodeAddressFromBech32(req.Node)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid node address %s", req.Node)
	}

	ctx := sdk.UnwrapSDKContext(c)

	session, found := q.GetActiveSessionForAddress(ctx, address, req.Subscription, node)
	if !found {
		return nil, nil
	}

	return &types.QueryActiveSessionResponse{Session: session}, nil
}
