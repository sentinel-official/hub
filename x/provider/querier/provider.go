package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func queryProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	provider, found := k.GetProvider(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(provider)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryProviders(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryProvidersParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	providers := k.GetProviders(ctx, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(providers)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
