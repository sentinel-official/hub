package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	provider, found := k.GetNode(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(provider)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryNodes(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryNodesParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodes(ctx, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInactiveNodes(ctx, params.Skip, params.Limit)
	} else {
		nodes = k.GetNodes(ctx, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryNodesForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryNodesForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInactiveNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else {
		nodes = k.GetNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
