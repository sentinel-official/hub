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

	item, found := q.GetSubscription(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription does not exist for id %d", req.Id)
	}

	return &types.QuerySubscriptionResponse{Subscription: item}, nil
}

func (q *queryServer) QuerySubscriptions(c context.Context, req *types.QuerySubscriptionsRequest) (*types.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Subscriptions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			var item types.Subscription
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

	return &types.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *queryServer) QuerySubscriptionsForAddress(c context.Context, req *types.QuerySubscriptionsForAddressRequest) (*types.QuerySubscriptionsForAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items      types.Subscriptions
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.Active) {
		store := prefix.NewStore(q.Store(ctx), types.GetActiveSubscriptionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetSubscription(ctx, sdk.BigEndianToUint64(key))
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactiveSubscriptionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetSubscription(ctx, sdk.BigEndianToUint64(key))
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else {
		// NOTE: Do not use this; less efficient; consider using active + inactive

		store := prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			if accumulate {
				var item types.Subscription
				if err := q.cdc.Unmarshal(value, &item); err != nil {
					return false, err
				}
				if !strings.EqualFold(item.Owner, req.Address) {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsForAddressResponse{Subscriptions: items, Pagination: pagination}, nil
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
