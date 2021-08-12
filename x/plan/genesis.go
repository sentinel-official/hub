package plan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
		var (
			provider = item.Plan.GetProvider()
		)

		k.SetPlan(ctx, item.Plan)
		if item.Plan.Status.Equal(hubtypes.StatusActive) {
			k.SetActivePlan(ctx, item.Plan.Id)
			k.SetActivePlanForProvider(ctx, provider, item.Plan.Id)
		} else {
			k.SetInactivePlan(ctx, item.Plan.Id)
			k.SetInactivePlanForProvider(ctx, provider, item.Plan.Id)
		}

		for _, node := range item.Nodes {
			address, err := hubtypes.NodeAddressFromBech32(node)
			if err != nil {
				panic(err)
			}

			k.SetNodeForPlan(ctx, item.Plan.Id, address)
			k.IncreaseCountForNodeByProvider(ctx, provider, address)
		}
	}

	count := uint64(0)
	for _, item := range state {
		if item.Plan.Id > count {
			count = item.Plan.Id
		}
	}

	k.SetCount(ctx, count)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	var (
		plans = k.GetPlans(ctx, 0, 0)
		items = make(types.GenesisPlans, 0, len(plans))
	)

	for _, plan := range plans {
		item := types.GenesisPlan{
			Plan:  plan,
			Nodes: []string{},
		}

		nodes := k.GetNodesForPlan(ctx, plan.Id, 0, 0)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return types.NewGenesisState(items)
}
