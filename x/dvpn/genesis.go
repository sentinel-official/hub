package dvpn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	provider.InitGenesis(ctx, k.Provider, state.Providers)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Providers: provider.ExportGenesis(ctx, k.Provider),
	}
}

func ValidateGenesis(state types.GenesisState) error {
	if err := provider.ValidateGenesis(state.Providers); err != nil {
		return err
	}

	return nil
}
