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
	"github.com/sentinel-official/hub/x/node/types"
)

type Querier struct {
	Keeper
}

func (q *Querier) QueryNode(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hub.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	node, found := q.GetNode(ctx, address)
	if !found {
		return nil, nil
	}

	return &types.QueryNodeResponse{Node: node}, nil
}

func (q *Querier) QueryNodes(c context.Context, req *types.QueryNodesRequest) (res *types.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		nodes      types.Nodes
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hub.Active) {
		store := prefix.NewStore(q.Store(ctx), types.ActiveNodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			node, found := q.GetNode(ctx, types.AddressFromStatusNodeKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	} else if req.Status.Equal(hub.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.InactiveNodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			node, found := q.GetNode(ctx, types.AddressFromStatusNodeKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	} else {
		store := prefix.NewStore(q.Store(ctx), types.NodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var node types.Node
			if err := q.cdc.UnmarshalBinaryBare(value, &node); err != nil {
				return false, err
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNodesResponse{Nodes: nodes, Pagination: pagination}, nil
}

func (q *Querier) QueryNodesForProvider(c context.Context, req *types.QueryNodesForProviderRequest) (res *types.QueryNodesForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	provider, err := hub.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		nodes      types.Nodes
		pagination *query.PageResponse
		ctx        = sdk.UnwrapSDKContext(c)
	)

	if req.Status.Equal(hub.Active) {
		store := prefix.NewStore(q.Store(ctx), types.GetActiveNodeForProviderKeyPrefix(provider))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			node, found := q.GetNode(ctx, types.AddressFromStatusNodeForProviderKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	} else if req.Status.Equal(hub.Inactive) {
		store := prefix.NewStore(q.Store(ctx), types.GetInactiveNodeForProviderKeyPrefix(provider))
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
			node, found := q.GetNode(ctx, types.AddressFromStatusNodeForProviderKey(key))
			if !found {
				return false, nil
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	} else {
		// TODO: Use NodeForProviderKeyPrefix?

		store := prefix.NewStore(q.Store(ctx), types.NodeKeyPrefix)
		pagination, err = query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
			var node types.Node
			if err := q.cdc.UnmarshalBinaryBare(value, &node); err != nil {
				return false, err
			}
			if !strings.EqualFold(node.Provider, req.Address) {
				return false, nil
			}

			if accumulate {
				nodes = append(nodes, node)
			}

			return true, nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNodesForProviderResponse{Nodes: nodes, Pagination: pagination}, nil
}
