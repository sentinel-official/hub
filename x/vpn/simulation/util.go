package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	denoms   = []string{"stake", "xxx", "yyy", "zzz"}
	statuses = []string{types.StatusActive, types.StatusInactive}
)

func getRandomDenom(r *rand.Rand) string {
	index := r.Intn(len(denoms))
	return denoms[index]
}

func getRandomStatus(r *rand.Rand) string {
	index := r.Intn(len(statuses))
	return statuses[index]
}

func getRandomID(r *rand.Rand) hub.ID {
	i := uint64(r.Int63n(10))

	return hub.NewNodeID(i)
}

func getRandomEncryption(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomType(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomVersion(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomMoniker(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomCoin(r *rand.Rand) sdk.Coin {
	denom := getRandomDenom(r)
	amount := simulation.RandIntBetween(r, 1, 1000)

	return sdk.NewCoin(denom, sdk.NewInt(int64(amount)))
}

func getRandomCoins(r *rand.Rand) (coins sdk.Coins) {
	coins = append(coins, getRandomCoin(r))

	size := r.Intn(2)
	for i := 0; i < size; i++ {
		coin := getRandomCoin(r)
		if coins == nil || coins.AmountOf(coin.Denom).IsZero() {
			coins = append(coins, coin)
		}
	}

	return coins.Sort()
}

func getRandomBandwidth(r *rand.Rand) hub.Bandwidth {
	upload := r.Int63n(hub.GB.Int64())
	download := r.Int63n(hub.GB.Int64())

	return hub.NewBandwidthFromInt64(upload, download)
}

func getRandomBandwidthSignature(r *rand.Rand, accounts []simulation.Account) auth.StdSignature {
	account := simulation.RandomAcc(r, accounts)
	bandwidthSignData := types.NewBandwidthSignatureData(getRandomID(r), uint64(r.Int63n(10)), getRandomBandwidth(r))
	signData, _ := account.PrivKey.Sign(bandwidthSignData.Bytes())

	return auth.StdSignature{
		PubKey:    account.PubKey,
		Signature: signData,
	}
}
