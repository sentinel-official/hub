package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func queryFreeNodesOfClient(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesOfFreeClientPrams

	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	freeClients := k.GetFreeNodesOfClient(ctx, params.Address)

	res, err := types.ModuleCdc.MarshalJSON(freeClients)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryFreeClientsOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryFreeClientsOfNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	freeClients := k.GetFreeClientsOfNode(ctx, params.ID)

	res, err := types.ModuleCdc.MarshalJSON(freeClients)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
