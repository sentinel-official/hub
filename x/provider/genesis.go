package provider

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, provider := range state.Providers {
		k.SetProvider(ctx, provider)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(k.GetProviders(ctx, 0, 0), k.GetParams(ctx))
}

func ValidateGenesis(state types.GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, provider := range state.Providers {
		if err := provider.Validate(); err != nil {
			return err
		}
	}

	providers := make(map[string]bool)
	for _, provider := range state.Providers {
		address := provider.Address.String()
		if providers[address] {
			return fmt.Errorf("found duplicate provider address %s", address)
		}

		providers[address] = true
	}

	return nil
}
