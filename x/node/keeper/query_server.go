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
	"github.com/sentinel-official/hub/x/node/types"
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

func (q *queryServer) QueryNode(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hubtypes.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetNode(ctx, address)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node does not exist for address %s", req.Address)
	}

	return &types.QueryNodeResponse{Node: item}, nil
}

func (q *queryServer) QueryNodes(c context.Context, req *types.QueryNodesRequest) (res *types.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items      types.Nodes
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.Active) {
		store := prefix.NewStore(q.Store(ctx), types.ActiveNodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetNode(ctx, key[1:])
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.InactiveNodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetNode(ctx, key[1:])
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else {
		store := prefix.NewStore(q.Store(ctx), types.NodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			if accumulate {
				var item types.Node
				if err := q.cdc.Unmarshal(value, &item); err != nil {
					return false, err
				}

				items = append(items, item)
			}

			return true, nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNodesResponse{Nodes: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryNodesForProvider(c context.Context, req *types.QueryNodesForProviderRequest) (res *types.QueryNodesForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	provider, err := hubtypes.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items      types.Nodes
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hubtypes.Active) {
		store := prefix.NewStore(q.Store(ctx), types.GetActiveNodeForProviderKeyPrefix(provider))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetNode(ctx, key[1:])
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else if req.Status.Equal(hubtypes.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactiveNodeForProviderKeyPrefix(provider))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			if accumulate {
				item, found := q.GetNode(ctx, key[1:])
				if !found {
					return false, nil
				}

				items = append(items, item)
			}

			return true, nil
		})
	} else {
		// NOTE: Do not use this; less efficient; consider using active + inactive

		store := prefix.NewStore(q.Store(ctx), types.NodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			if accumulate {
				var item types.Node
				if err := q.cdc.Unmarshal(value, &item); err != nil {
					return false, err
				}
				if !strings.EqualFold(item.Provider, req.Address) {
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

	return &types.QueryNodesForProviderResponse{Nodes: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
