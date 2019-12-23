package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
	
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func queryResolvers(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var id hub.ResolverID
	var resolvers []types.Resolver
	
	if len(req.Data) != 0 {
		if err := types.ModuleCdc.UnmarshalJSON(req.Data, &id); err != nil {
			return nil, types.ErrorUnmarshal()
		}
		
		resolver, found := k.GetResolver(ctx, id)
		if !found {
			return nil, types.ErrorResolverDoesNotExist()
		}
		
		resolvers = append(resolvers, resolver)
	} else {
		resolvers = k.GetAllResolvers(ctx)
	}
	
	res, err := types.ModuleCdc.MarshalJSON(resolvers)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	return res, nil
}

func queryResolversOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var address sdk.AccAddress
	
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &address); err != nil {
		return nil, types.ErrorUnmarshal()
	}
	
	resolvers := k.GetResolversOfAddress(ctx, address)
	res, err := types.ModuleCdc.MarshalJSON(resolvers)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	return res, nil
}

func queryNodesOfResolver(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesOfResolverPrams
	
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}
	
	nodes := k.GetNodesOfResolver(ctx, params.ID)
	
	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	
	return res, nil
}

func queryResolversOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryResolversOfNodeParams
	
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}
	
	resolvers := k.GetResolversOfNode(ctx, params.ID)
	
	res, err := types.ModuleCdc.MarshalJSON(resolvers)
	if err != nil {
		return nil, types.ErrorMarshal()
	}
	
	return res, nil
}
