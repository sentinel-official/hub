package simulation

import (
	"math/rand"

	csim "github.com/cosmos/cosmos-sdk/x/simulation"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	status = []string{types.StatusActive, types.StatusInactive, types.StatusDeRegistered}
	denom  = []string{"stake", "steak"}
)

func getRandomID(r *rand.Rand) hub.ID {
	id := r.Intn(1000)
	_id := hub.NewIDFromUInt64(uint64(id))
	
	return _id
}

func getRandomCoin(r *rand.Rand) sdk.Coin {
	amount := r.Intn(1000)
	coin := sdk.NewCoin(getRandomDenom(r), sdk.NewInt(int64(amount)))

	return coin
}

func getRandomCoins(r *rand.Rand) sdk.Coins {
	index := r.Intn(2)
	var coins sdk.Coins
	for i := 0; i < index; i++ {
		amount := r.Intn(1000)
		coins = append(coins, sdk.NewCoin(getRandomDenom(r), sdk.NewInt(int64(amount))))
	}

	return coins
}

func getRandomBandwidth(r *rand.Rand) hub.Bandwidth {
	upload := r.Intn(2000000000) - 1000000000
	download := r.Intn(2000000000) - 1000000000

	return hub.NewBandwidthFromInt64(int64(upload), int64(download))
}

func getRandomSignData(r *rand.Rand, accs []csim.Account) auth.StdSignature {
	randSignBandwidthData := types.NewBandwidthSignatureData(getRandomID(r), getRandomIndex(r), getRandomBandwidth(r))
	randAccount := csim.RandomAcc(r, accs)
	signData, _ := randAccount.PrivKey.Sign(randSignBandwidthData.Bytes())

	return auth.StdSignature{
		PubKey:    randAccount.PubKey,
		Signature: signData,
	}
}

func getRandomDenom(r *rand.Rand) string { return denom[r.Intn(len(denom))] }

func getRandomIndex(r *rand.Rand) uint64 { return uint64(r.Intn(1000)) }

func getRandomEncryption(r *rand.Rand) string { return csim.RandStringOfLength(r, 10) }

func getRandomType(r *rand.Rand) string { return csim.RandStringOfLength(r, 10) }

func getRandomVersion(r *rand.Rand) string { return csim.RandStringOfLength(r, 10) }

func getRandomMoniker(r *rand.Rand) string { return csim.RandStringOfLength(r, 10) }

func getRandomStatus(r *rand.Rand) string { return status[r.Intn(len(status))] }
