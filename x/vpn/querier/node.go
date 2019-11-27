package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func queryNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodeParams

	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	node, found := k.GetNode(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(node)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodesOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesOfAddressPrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	nodes := k.GetNodesOfAddress(ctx, params.Address)

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllNodes(ctx sdk.Context, k keeper.Keeper) ([]byte, sdk.Error) {
	nodes := k.GetAllNodes(ctx)

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
