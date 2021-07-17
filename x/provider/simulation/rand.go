package simulation

import (
	"math/rand"

	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

const (
	MaxNameLength        = 56
	MaxIdentityLength    = 64
	MaxWebsiteLength     = 64
	MaxDescriptionLength = 256
)

func RandomProvider(r *rand.Rand, items types.Providers) types.Provider {
	if len(items) == 0 {
		return types.Provider{}
	}

	return items[r.Intn(len(items))]
}

func RandomProviders(r *rand.Rand, accounts []simulationtypes.Account) types.Providers {
	var (
		duplicates = make(map[string]bool)
		items      = make(types.Providers, 0, r.Intn(len(accounts)))
	)

	for ; len(items) < cap(items); {
		var (
			account, _ = simulationtypes.RandomAcc(r, accounts)
			address    = hubtypes.ProvAddress(account.Address).String()
		)

		if duplicates[address] {
			continue
		}

		var (
			name        = simulationtypes.RandStringOfLength(r, r.Intn(MaxNameLength)+8)
			identity    = simulationtypes.RandStringOfLength(r, r.Intn(MaxIdentityLength))
			website     = simulationtypes.RandStringOfLength(r, r.Intn(MaxWebsiteLength))
			description = simulationtypes.RandStringOfLength(r, r.Intn(MaxDescriptionLength))
		)

		duplicates[address] = true
		items = append(
			items,
			types.Provider{
				Address:     address,
				Name:        name,
				Identity:    identity,
				Website:     website,
				Description: description,
			},
		)
	}

	return items
}
