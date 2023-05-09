package plan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
		addr := item.Plan.GetAddress()
		k.SetPlan(ctx, item.Plan)
		k.SetPlanForProvider(ctx, addr, item.Plan.ID)

		for _, node := range item.Nodes {
			addr, err := hubtypes.NodeAddressFromBech32(node)
			if err != nil {
				panic(err)
			}

			k.SetNodeForPlan(ctx, item.Plan.ID, addr)
		}
	}

	count := uint64(0)
	for _, item := range state {
		if item.Plan.ID > count {
			count = item.Plan.ID
		}
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	var (
		plans = k.GetPlans(ctx)
		items = make(types.GenesisPlans, 0, len(plans))
	)

	for _, plan := range plans {
		item := types.GenesisPlan{
			Plan:  plan,
			Nodes: []string{},
		}

		nodes := k.GetNodesForPlan(ctx, plan.ID)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return types.NewGenesisState(items)
}
