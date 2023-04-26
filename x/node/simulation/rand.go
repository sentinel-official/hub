package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	simulationhubtypes "github.com/sentinel-official/hub/types/simulation"
	"github.com/sentinel-official/hub/x/node/types"
)

const (
	MaxPricesAmount    = 1 << 18
	MaxRemoteURLLength = 48
)

func RandomNodes(r *rand.Rand, accounts []simulationtypes.Account) types.Nodes {
	var (
		duplicates = make(map[string]bool)
		items      = make(types.Nodes, 0, r.Intn(len(accounts)))
	)

	for len(items) < cap(items) {
		var (
			account, _ = simulationtypes.RandomAcc(r, accounts)
			address    = hubtypes.NodeAddress(account.Address).String()
		)

		if duplicates[address] {
			continue
		}

		var (
			gigabytePrices = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxPricesAmount,
					),
				),
			)
			hourlyPrices = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxPricesAmount,
					),
				),
			)
			remoteURL = fmt.Sprintf(
				"https://%s:8080",
				simulationtypes.RandStringOfLength(r, r.Intn(MaxRemoteURLLength)),
			)
			status   = hubtypes.StatusActive
			statusAt = time.Now()
		)

		if rand.Intn(2) == 0 {
			status = hubtypes.StatusInactive
		}

		duplicates[address] = true
		items = append(
			items,
			types.Node{
				Address:        address,
				GigabytePrices: gigabytePrices,
				HourlyPrices:   hourlyPrices,
				RemoteURL:      remoteURL,
				Status:         status,
				StatusAt:       statusAt,
			},
		)
	}

	return items
}
