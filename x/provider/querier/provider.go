package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
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

func queryProviders(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryProvidersParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	providers := k.GetProviders(ctx)

	start, end := client.Paginate(len(providers), params.Page, params.Limit, len(providers))
	if start < 0 || end < 0 {
		providers = types.Providers{}
	} else {
		providers = providers[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(providers)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
