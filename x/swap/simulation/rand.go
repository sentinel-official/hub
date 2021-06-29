package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
)

func GetRandomSwapDenom(r *rand.Rand) string {
	// denom length should be between 3-128
	return simulation.RandStringOfLength(r, r.Intn(125)+3)
}

func GetRandomApproveBy(r *rand.Rand) string {
	bz := make([]byte, 20)
	_, err := r.Read(bz)
	if err != nil {
		panic(err)
	}

	return sdk.AccAddress(bz).String()
}

func GetRandomSwapEnabled(r *rand.Rand) bool {
	return r.Intn(2) == 1
}

