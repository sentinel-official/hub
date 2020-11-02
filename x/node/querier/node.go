package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	provider, found := k.GetNode(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(provider)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodes(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodes(ctx, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInActiveNodes(ctx, params.Skip, params.Limit)
	} else {
		nodes = k.GetNodes(ctx, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodesForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInActiveNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	} else {
		nodes = k.GetNodesForProvider(ctx, params.Address, params.Skip, params.Limit)
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
