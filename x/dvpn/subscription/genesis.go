package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state.Plans {
		k.SetPlan(ctx, item.Plan)
		k.SetPlanIDForProvider(ctx, item.Plan.Provider, item.Plan.ID)

		for _, address := range item.Nodes {
			k.SetNodeAddressForPlan(ctx, item.Plan.ID, address)
		}

		k.SetPlansCount(ctx, k.GetPlansCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	state := types.NewGenesisState(nil)
	plans := k.GetPlans(ctx)

	for _, item := range plans {
		plan := GenesisPlan{
			Plan:  item,
			Nodes: nil,
		}

		nodes := k.GetNodesForPlan(ctx, item.ID)
		for _, node := range nodes {
			plan.Nodes = append(plan.Nodes, node.Address)
		}

		state.Plans = append(state.Plans, plan)
	}

	return state
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
