package keeper

import (
	"context"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
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

func (q *queryServer) QuerySubscription(c context.Context, req *types.QuerySubscriptionRequest) (*types.QuerySubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	v, found := q.GetSubscription(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription does not exist for id %d", req.Id)
	}

	item, err := codectypes.NewAnyWithValue(v)
	if err != nil {
		return nil, err
	}

	return &types.QuerySubscriptionResponse{Subscription: item}, nil
}

func (q *queryServer) QuerySubscriptions(c context.Context, req *types.QuerySubscriptionsRequest) (*types.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var v types.Subscription
		if err := q.cdc.UnmarshalInterface(value, &v); err != nil {
			return err
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySubscriptionsForAccount(c context.Context, req *types.QuerySubscriptionsForAccountRequest) (*types.QuerySubscriptionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := q.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsForAccountResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySubscriptionsForNode(c context.Context, req *types.QuerySubscriptionsForNodeRequest) (*types.QuerySubscriptionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := hubtypes.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSubscriptionForNodeKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := q.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsForNodeResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySubscriptionsForPlan(c context.Context, req *types.QuerySubscriptionsForPlanRequest) (*types.QuerySubscriptionsForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []*codectypes.Any
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.Id))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		v, found := q.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		item, err := codectypes.NewAnyWithValue(v)
		if err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryAllocation(c context.Context, req *types.QueryAllocationRequest) (*types.QueryAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetAllocation(ctx, req.Id, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allocation %d/%s does not exist", req.Id, req.Address)
	}

	return &types.QueryAllocationResponse{Allocation: item}, nil
}

func (q *queryServer) QueryAllocations(c context.Context, req *types.QueryAllocationsRequest) (*types.QueryAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Allocations
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetAllocationForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item types.Allocation
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllocationsResponse{Allocations: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPayout(c context.Context, req *types.QueryPayoutRequest) (*types.QueryPayoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetPayout(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "payout does not exist for id %d", req.Id)
	}

	return &types.QueryPayoutResponse{Payout: item}, nil
}

func (q *queryServer) QueryPayouts(c context.Context, req *types.QueryPayoutsRequest) (res *types.QueryPayoutsResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.PayoutKeyPrefix)
	)

	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item types.Payout
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPayoutsResponse{Payouts: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPayoutsForAccount(c context.Context, req *types.QueryPayoutsForAccountRequest) (res *types.QueryPayoutsForAccountResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	var (
		items types.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetPayoutForAccountKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetPayout(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("payout for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPayoutsForAccountResponse{Payouts: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPayoutsForNode(c context.Context, req *types.QueryPayoutsForNodeRequest) (res *types.QueryPayoutsForNodeResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := hubtypes.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	var (
		items types.Payouts
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetPayoutForNodeKeyPrefix(addr))
	)

	pagination, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := q.GetPayout(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("payout for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPayoutsForNodeResponse{Payouts: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
