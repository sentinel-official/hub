package swap

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, item := range state.Swaps {
		k.SetSwap(ctx, item)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(
		k.GetSwaps(ctx, 0, 0),
		k.GetParams(ctx),
	)
}

func ValidateGenesis(state types.GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, item := range state.Swaps {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	swaps := make(map[string]bool)
	for _, item := range state.Swaps {
		txHash := item.GetTxHash().String()
		if swaps[txHash] {
			return fmt.Errorf("duplicate swap for tx_hash %s", txHash)
		}

		swaps[txHash] = true
	}

	return nil
}
