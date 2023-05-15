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
		items     types.Plans
		keyPrefix []byte
		ctx       = sdk.UnwrapSDKContext(c)
	)

	switch req.Status {
	case hubtypes.StatusActive:
		keyPrefix = types.ActivePlanKeyPrefix
	case hubtypes.StatusInactive:
		keyPrefix = types.InactivePlanKeyPrefix
	default:
		keyPrefix = types.PlanKeyPrefix
	}

	store := prefix.NewStore(q.Store(ctx), keyPrefix)
	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item types.Plan
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPlansResponse{Plans: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPlansForProvider(c context.Context, req *types.QueryPlansForProviderRequest) (res *types.QueryPlansForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := hubtypes.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items types.Plans
		ctx   = sdk.UnwrapSDKContext(c)
	)

	store := prefix.NewStore(q.Store(ctx), types.GetPlanForProviderKeyPrefix(addr))
	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return false, fmt.Errorf("plan for key %X does not exist", key)
		}

		if req.Status.IsOneOf(item.Status, hubtypes.StatusUnspecified) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPlansForProviderResponse{Plans: items, Pagination: pagination}, nil
}
