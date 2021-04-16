package v0_6

import (
	"github.com/sentinel-official/hub/x/plan/types"
	legacy "github.com/sentinel-official/hub/x/plan/types/legacy/v0.5"
)

func MigrateGenesisPlan(item legacy.GenesisPlan) types.GenesisPlan {
	var nodes []string
	for _, item := range item.Nodes {
		nodes = append(nodes, item.String())
	}

	return types.GenesisPlan{
		Plan:  MigratePlan(item.Plan),
		Nodes: nodes,
	}
}

func MigrateGenesisPlans(items legacy.GenesisPlans) types.GenesisPlans {
	var plans types.GenesisPlans
	for _, item := range items {
		plans = append(plans, MigrateGenesisPlan(item))
	}

	return plans
}

func MigrateGenesisState(state legacy.GenesisState) types.GenesisState {
	return types.NewGenesisState(
		MigrateGenesisPlans(legacy.GenesisPlans(state)),
	)
}
