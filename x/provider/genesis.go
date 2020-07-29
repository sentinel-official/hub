package provider

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, provider := range state {
		k.SetProvider(ctx, provider)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetProviders(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, provider := range state {
		if err := provider.Validate(); err != nil {
			return err
		}
	}

	providers := make(map[string]bool)
	for _, provider := range state {
		address := provider.Address.String()
		if providers[address] {
			return fmt.Errorf("found duplicate provider address '%s'", provider.Address)
		}

		providers[address] = true
	}

	return nil
}
