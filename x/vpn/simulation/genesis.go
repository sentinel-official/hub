package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	depositsim "github.com/sentinel-official/hub/x/deposit/simualtion"
	nodesim "github.com/sentinel-official/hub/x/node/simulation"
	plansim "github.com/sentinel-official/hub/x/plan/simulation"
	providersim "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsim "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsim "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func RandomizedGenesisState(simState *module.SimulationState) {
	deposits := depositsim.RandomizedGenesisState(simState)
	providers := providersim.RandomizedGenesisState(simState)
	nodes := nodesim.RandomizedGenState(simState)
	plans := plansim.RandomizedGenState(simState)
	subscriptions := subscriptionsim.RandomizedGenState(simState)
	sessions := sessionsim.RandomizedGenState(simState)

	vpnGenesis := types.NewGenesisState(deposits, providers, nodes, plans, subscriptions, sessions)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(vpnGenesis)
}
