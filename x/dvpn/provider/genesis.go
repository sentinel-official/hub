package provider

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return nil
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
