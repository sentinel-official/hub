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
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	"github.com/sentinel-official/hub/x/plan/types"
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

func (q *queryServer) QueryPlan(c context.Context, req *types.QueryPlanRequest) (*types.QueryPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetPlan(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "plan does not exist for id %d", req.Id)
	}

	return &types.QueryPlanResponse{Plan: item}, nil
}

func (q *queryServer) QueryPlans(c context.Context, req *types.QueryPlansRequest) (res *types.QueryPlansResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items      types.Plans
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.Active) {
		store := prefix.NewStore(q.Store(ctx), types.ActivePlanKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.InactivePlanKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else {
		store := prefix.NewStore(q.Store(ctx), types.PlanKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Plan
			if err := q.cdc.UnmarshalBinaryBare(value, &item); err != nil {
				return false, err
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

	return &types.QueryPlansResponse{Plans: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPlansForProvider(c context.Context, req *types.QueryPlansForProviderRequest) (res *types.QueryPlansForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hubtypes.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items      types.Plans
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.Active) {
		store := prefix.NewStore(q.Store(ctx), types.GetActivePlanForProviderKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
			if !found {
				return false, nil
			}

			if accumulate {
				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactivePlanForProviderKeyPrefix(address))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
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

		store := prefix.NewStore(q.Store(ctx), types.PlanKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var item types.Plan
			if err := q.cdc.UnmarshalBinaryBare(value, &item); err != nil {
				return false, err
			}
			if !strings.EqualFold(item.Provider, req.Address) {
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

	return &types.QueryPlansForProviderResponse{Plans: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryNodesForPlan(c context.Context, req *types.QueryNodesForPlanRequest) (*types.QueryNodesForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items      nodetypes.Nodes
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	store := prefix.NewStore(q.Store(ctx), types.GetNodeForPlanKeyPrefix(req.Id))
	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		item, found := q.GetNode(ctx, key)
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

	return &types.QueryNodesForPlanResponse{Nodes: items, Pagination: pagination}, nil
}
