package plan

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, item := range state {
		k.SetPlan(ctx, item.Plan)

		if item.Plan.Status.Equal(hubtypes.StatusActive) {
			k.SetActivePlan(ctx, item.Plan.Id)
			k.SetActivePlanForProvider(ctx, item.Plan.GetProvider(), item.Plan.Id)
		} else {
			k.SetInactivePlan(ctx, item.Plan.Id)
			k.SetInactivePlanForProvider(ctx, item.Plan.GetProvider(), item.Plan.Id)
		}

		for _, node := range item.Nodes {
			address, err := hubtypes.NodeAddressFromBech32(node)
			if err != nil {
				panic(err)
			}

			k.SetNodeForPlan(ctx, item.Plan.Id, address)
		}
	}

	k.SetCount(ctx, uint64(len(state)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	plans := k.GetPlans(ctx, 0, 0)

	items := make(types.GenesisPlans, 0, len(plans))
	for _, plan := range plans {
		item := types.GenesisPlan{
			Plan:  plan,
			Nodes: nil,
		}

		nodes := k.GetNodesForPlan(ctx, plan.Id, 0, 0)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return types.NewGenesisState(items)
}

func ValidateGenesis(state types.GenesisState) error {
	for _, item := range state {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	plans := make(map[uint64]bool)
	for _, item := range state {
		id := item.Plan.Id
		if plans[id] {
			return fmt.Errorf("duplicate plan id %d", id)
		}

		plans[id] = true
	}

	for _, item := range state {
		nodes := make(map[string]bool)
		for _, address := range item.Nodes {
			if nodes[address] {
				return fmt.Errorf("duplicate node for plan %d", item.Plan.Id)
			}

			nodes[address] = true
		}
	}

	return nil
}
