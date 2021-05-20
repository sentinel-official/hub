package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func getRandomCoins(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin("sent", sdk.NewInt(r.Int63n(10<<12))))
}

func getRandonDeposits(r *rand.Rand) types.Deposits {
	var deposits types.Deposits

	for _, acc := range sdksimulation.RandomAccounts(r, r.Intn(18)+2) {
		deposits = append(deposits, types.Deposit{
			Address: acc.Address.String(),
			Coins:   getRandomCoins(r),
		})
	}

	return deposits
}

func RandomizedGenesisState(simState *module.SimulationState) types.GenesisState {
	return types.NewGenesisState(getRandonDeposits(simState.Rand))
}
