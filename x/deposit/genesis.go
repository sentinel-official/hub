package deposit

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, deposit := range state {
		k.SetDeposit(ctx, deposit)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetDeposits(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, deposit := range state {
		if err := deposit.Validate(); err != nil {
			return err
		}
	}

	deposits := make(map[string]bool)
	for _, deposit := range state {
		address := deposit.Address.String()
		if deposits[address] {
			return fmt.Errorf("found duplicate deposit address %s", address)
		}

		deposits[address] = true
	}

	return nil
}
