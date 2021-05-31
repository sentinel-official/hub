package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
)

func getRandomDeposit(r *rand.Rand) int64 {
	return int64(r.Intn(100) + 1)
}

func getRandomInactiveDuration(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60, 60<<13))
}

func getNodeAddress() hubtypes.NodeAddress {
	bz := make([]byte, 20)
	_, err := rand.Read(bz)
	if err != nil {
		panic(err)
	}

	return hubtypes.NodeAddress(bz)
}

