package keeper

import (
	"context"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	hubtypes "github.com/sentinel-official/hub/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/x/subscription/types"
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

	return &types.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryQuota(c context.Context, req *types.QueryQuotaRequest) (*types.QueryQuotaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetQuota(ctx, req.Id, address)
	if !found {
		return nil, status.Errorf(codes.NotFound, "quota does not exist for id %d, and address %s", req.Id, req.Address)
	}

	return &types.QueryQuotaResponse{Quota: item}, nil
}

func (q *queryServer) QueryQuotas(c context.Context, req *types.QueryQuotasRequest) (*types.QueryQuotasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Quotas
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetQuotaKeyPrefix(req.Id))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			var item types.Quota
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryQuotasResponse{Quotas: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
