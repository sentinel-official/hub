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
	MaxPriceAmount     = 1 << 18
	MaxRemoteURLLength = 48
)

func RandomNode(r *rand.Rand, items types.Nodes) types.Node {
	if len(items) == 0 {
		return types.Node{}
	}

	return items[r.Intn(len(items))]
}

func RandomNodes(r *rand.Rand, accounts []simulationtypes.Account) types.Nodes {
	var (
		duplicates = make(map[string]bool)
		items      = make(types.Nodes, 0, r.Intn(len(accounts)))
	)

	for ; len(items) < cap(items); {
		var (
			account, _ = simulationtypes.RandomAcc(r, accounts)
			address    = hubtypes.NodeAddress(account.Address).String()
		)

		if duplicates[address] {
			continue
		}

		var (
			price = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxPriceAmount,
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
				Address:   address,
				Provider:  "",
				Price:     price,
				RemoteURL: remoteURL,
				Status:    status,
				StatusAt:  statusAt,
			},
		)
	}

	return items
}
