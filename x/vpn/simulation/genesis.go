package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	depositsimulation "github.com/sentinel-official/hub/x/deposit/simulation"
	nodesimulation "github.com/sentinel-official/hub/x/node/simulation"
	plansimulation "github.com/sentinel-official/hub/x/plan/simulation"
	providersimulation "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsimulation "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsimulation "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func RandomizedGenesisState(state *module.SimulationState) {
	var (
		depositGenesisState      = depositsimulation.RandomizedGenesisState(state)
		providerGenesisState     = providersimulation.RandomizedGenesisState(state)
		nodeGenesisState         = nodesimulation.RandomizedGenesisState(state)
		planGenesisState         = plansimulation.RandomizedGenesisState(state)
		subscriptionGenesisState = subscriptionsimulation.RandomizedGenesisState(state)
		sessionGenesisState      = sessionsimulation.RandomizedGenesisState(state)
	)

	for i := 0; i < len(planGenesisState); i++ {
		if state.Rand.Intn(2) == 0 {
			planGenesisState[i].Plan.Address = providersimulation.RandomProvider(
				state.Rand,
				providerGenesisState.Providers,
			).Address
		}
	}

	state.GenState[types.ModuleName] = state.Cdc.MustMarshalJSON(
		types.NewGenesisState(
			depositGenesisState,
			providerGenesisState,
			nodeGenesisState,
			planGenesisState,
			subscriptionGenesisState,
			sessionGenesisState,
		),
	)
}
