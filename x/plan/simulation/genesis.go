package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sentinel-official/hub/x/plan/types"
)

const (
	dayDurationInSeconds   = 86400
	weekDurationInSeconds  = dayDurationInSeconds * 7
	monthDurationInSeconds = dayDurationInSeconds * 30
	gigabytes              = 1024 << 20
	terabytes              = gigabytes * 1024
)

func RandomizedGenState(simState *module.SimulationState) types.GenesisState {
	var plans types.GenesisPlans

	for _, plan := range getRandomPlans(simState.Rand) {
		plans = append(plans, types.GenesisPlan{
			Plan:  plan,
			Nodes: getRandomNodes(simState.Rand),
		})
	}

	return types.NewGenesisState(plans)
}
