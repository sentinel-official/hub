package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/vpn/node/keeper"
	"github.com/sentinel-official/hub/x/vpn/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, node := range state.Nodes {
		k.SetNode(ctx, node)
		k.SetNodeForProvider(ctx, node.Provider, node.Address)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(
		k.GetNodes(ctx),
		k.GetParams(ctx),
	)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
