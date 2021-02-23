package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/x/deposit/types"
)

type Querier struct {
	Keeper
}

func (q *Querier) QueryDeposit(c context.Context, req *types.QueryDepositRequest) (*types.QueryDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	deposit, found := q.GetDeposit(ctx, address)
	if !found {
		return nil, nil
	}

	return &types.QueryDepositResponse{Deposit: deposit}, nil
}

func (q *Querier) QueryDeposits(c context.Context, req *types.QueryDepositsRequest) (*types.QueryDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		deposits types.Deposits
		ctx      = sdk.UnwrapSDKContext(c)
		store    = prefix.NewStore(q.Store(ctx), types.DepositKeyPrefix)
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		var deposit types.Deposit
		if err := q.cdc.UnmarshalBinaryBare(value, &deposit); err != nil {
			return false, err
		}
		if accumulate {
			deposits = append(deposits, deposit)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDepositsResponse{Deposits: deposits, Pagination: pagination}, nil
}
