package simulation

import (
	"math/rand"

	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func RandomProvider(r *rand.Rand, providers types.Providers) types.Provider {
	if len(providers) == 0 {
		return types.Provider{
			Address: []byte("address"),
		}
	}

	return providers[r.Intn(
		len(providers),
	)]
}
