package deposit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/deposit/types"
)

func InitGenesis(ctx sdk.Context, k Keeper, state types.GenesisState) {
	for _, deposit := range state {
		k.SetDeposit(ctx, deposit)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	return k.GetDeposits(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
