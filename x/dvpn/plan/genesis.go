package plan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/plan/keeper"
	"github.com/sentinel-official/hub/x/dvpn/plan/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
		k.SetPlan(ctx, item.Plan)
		k.SetPlanForProvider(ctx, item.Plan.Provider, item.Plan.ID)

		for _, node := range item.Nodes {
			k.SetNodeForPlan(ctx, item.Plan.ID, node)
			k.SetPlanForNode(ctx, node, item.Plan.ID)
		}

		k.SetPlansCount(ctx, k.GetPlansCount(ctx)+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	_plans := k.GetPlans(ctx)

	plans := make(types.GenesisPlans, 0, len(_plans))
	for _, item := range _plans {
		plan := types.GenesisPlan{
			Plan:  item,
			Nodes: nil,
		}

		nodes := k.GetNodesForPlan(ctx, item.ID)
		for _, node := range nodes {
			plan.Nodes = append(plan.Nodes, node.Address)
		}

		plans = append(plans, plan)
	}

	return types.NewGenesisState(plans)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
