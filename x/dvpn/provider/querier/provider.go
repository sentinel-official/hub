package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func queryProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	provider, found := k.GetProvider(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(provider)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryProviders(ctx sdk.Context, _ abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	res, err := types.ModuleCdc.MarshalJSON(k.GetProviders(ctx))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
