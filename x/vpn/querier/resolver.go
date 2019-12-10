package querier

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func queryResolvers(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var address string
	var resolvers []types.Resolver

	if len(req.Data) != 0 {
		if err := types.ModuleCdc.UnmarshalJSON(req.Data, &address); err != nil {
			fmt.Println(err)
			return nil, types.ErrorUnmarshal()
		}

		_address, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return nil, types.ErrorUnmarshal()
		}

		resolver, found := k.GetResolver(ctx, _address)
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
