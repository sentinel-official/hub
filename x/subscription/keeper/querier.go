package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

type Querier struct {
	Keeper
}

func (q *Querier) QuerySubscription(c context.Context, req *types.QuerySubscriptionRequest) (*types.QuerySubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetSubscription(ctx, req.Id)
	if !found {
		return nil, nil
	}

	return &types.QuerySubscriptionResponse{Subscription: item}, nil
}

func (q *Querier) QuerySubscriptions(c context.Context, req *types.QuerySubscriptionsRequest) (*types.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Subscriptions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var item types.Subscription
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

	return &types.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySubscriptionsForNode(c context.Context, req *types.QuerySubscriptionsForNodeRequest) (*types.QuerySubscriptionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hub.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Subscriptions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSubscriptionForNodeKeyPrefix(address))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetSubscription(ctx, types.IDFromSubscriptionForNodeKey(key))
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

	return &types.QuerySubscriptionsForNodeResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySubscriptionsForPlan(c context.Context, req *types.QuerySubscriptionsForPlanRequest) (*types.QuerySubscriptionsForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Subscriptions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.Id))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(key))
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

	return &types.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *Querier) QuerySubscriptionsForAddress(c context.Context, req *types.QuerySubscriptionsForAddressRequest) (*types.QuerySubscriptionsForAddressResponse, error) {
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

	if req.Status.Equal(hub.Active) {
		store := prefix.NewStore(q.Store(ctx), types.GetActiveSubscriptionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetSubscription(ctx, types.IDFromStatusSubscriptionForAddressKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hub.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactiveSubscriptionForAddressKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetSubscription(ctx, types.IDFromStatusSubscriptionForAddressKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else {
		// TODO: Use SubscriptionForAddressKeyPrefix?

		store := prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Subscription
			if err := q.cdc.UnmarshalBinaryBare(value, &item); err != nil {
				return false, err
			}
			if !strings.EqualFold(item.Owner, req.Address) {
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

	return &types.QuerySubscriptionsForAddressResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (q *Querier) QueryQuota(c context.Context, req *types.QueryQuotaRequest) (*types.QueryQuotaResponse, error) {
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
		return nil, nil
	}

	return &types.QueryQuotaResponse{Quota: item}, nil
}

func (q *Querier) QueryQuotas(c context.Context, req *types.QueryQuotasRequest) (*types.QueryQuotasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Quotas
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var item types.Quota
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

	return &types.QueryQuotasResponse{Quotas: items, Pagination: pagination}, nil
}
