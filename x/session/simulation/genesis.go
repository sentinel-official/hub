package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sentinel-official/hub/x/session/types"
)

func RandomizedGenState(simState *module.SimulationState) *types.GenesisState {

	var (
		inactiveDuration          time.Duration
		proofVerificationEnabled bool
	)

	inactiveDurationSim := func(r *rand.Rand) {
		inactiveDuration = getRandomInactiveDuration(r)
	}

	proofVerificationEnabledSim := func(r *rand.Rand) {
		proofVerificationEnabled = getRandomProofVerificationEnabled(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyInactiveDuration), inactiveDuration, nil, inactiveDurationSim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyProofVerificationEnabled), proofVerificationEnabled, nil, proofVerificationEnabledSim)

	params := types.NewParams(inactiveDuration, proofVerificationEnabled)
	sessions := getRandomSessions(simState.Rand)

	state := types.NewGenesisState(sessions, params)
	bz := simState.Cdc.MustMarshalJSON(&state.Params)

	fmt.Printf("selected randomly generated sessions parameters: %s\n", bz)
	return state
}
