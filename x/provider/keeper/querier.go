package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

type Querier struct {
	Keeper
}

func (q *Querier) QueryProvider(c context.Context, req *types.QueryProviderRequest) (*types.QueryProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := hub.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	provider, found := q.GetProvider(ctx, address)
	if !found {
		return nil, nil
	}

	return &types.QueryProviderResponse{Provider: provider}, nil
}

func (q *Querier) QueryProviders(c context.Context, req *types.QueryProvidersRequest) (*types.QueryProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		providers types.Providers
		ctx       = sdk.UnwrapSDKContext(c)
		store     = prefix.NewStore(q.Store(ctx), types.ProviderKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var provider types.Provider
		if err := q.cdc.UnmarshalBinaryBare(value, &provider); err != nil {
			return false, err
		}
		if accumulate {
			providers = append(providers, provider)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProvidersResponse{Providers: providers, Pagination: pagination}, nil
}
