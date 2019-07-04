package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	status = []string{types.StatusActive, types.StatusInactive}
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

func getRandomValidCoin(r *rand.Rand) sdk.Coin {
	amount := r.Intn(1000)+5
	coin := sdk.NewCoin("stake", sdk.NewInt(int64(amount)))

	return coin
}

func getRandomCoins(r *rand.Rand) sdk.Coins {
	index := r.Intn(2)
	var coins sdk.Coins
	for i := 0; i < index; i++ {
		coins = append(coins, getRandomCoin(r))
	}

	return coins
}

func getRandomValidCoins(r *rand.Rand) sdk.Coins {

	return sdk.Coins{getRandomValidCoin(r)}
}

func getRandomBandwidth(r *rand.Rand) hub.Bandwidth {
	upload := r.Intn(2000000000) - 1000000000
	download := r.Intn(2000000000) - 1000000000

	return hub.NewBandwidthFromInt64(int64(upload), int64(download))
}

func getRandomValidBandwidth(r *rand.Rand) hub.Bandwidth {
	upload := r.Intn(1000000000)
	download := r.Intn(1000000000)

	return hub.NewBandwidthFromInt64(int64(upload), int64(download))
}

func getRandomSignData(r *rand.Rand, accs []simulation.Account) auth.StdSignature {
	randSignBandwidthData := types.NewBandwidthSignatureData(getRandomID(r), getRandomIndex(r), getRandomBandwidth(r))
	randAccount := simulation.RandomAcc(r, accs)
	signData, _ := randAccount.PrivKey.Sign(randSignBandwidthData.Bytes())

	return auth.StdSignature{
		PubKey:    randAccount.PubKey,
		Signature: signData,
	}
}

func getRandomDeposit(r *rand.Rand, accs []simulation.Account) deposit.Deposit {

	return deposit.Deposit{
		Address: simulation.RandomAcc(r, accs).Address,
		Coins:   getRandomValidCoins(r),
	}
}

func GetRandomDeposits(r *rand.Rand, acc []simulation.Account) []deposit.Deposit {
	var deposits []deposit.Deposit
	for i := 0; i < r.Intn(1); i++ {
		deposits = append(deposits, getRandomDeposit(r, acc))
	}

	return deposits
}

func getRandomNode(r *rand.Rand, accs []simulation.Account) types.Node {

	return types.Node{
		ID:               getRandomID(r),
		Owner:            simulation.RandomAcc(r, accs).Address,
		Deposit:          getRandomValidCoin(r),
		Type:             getRandomType(r),
		Version:          getRandomVersion(r),
		Moniker:          getRandomMoniker(r),
		PricesPerGB:      getRandomValidCoins(r),
		InternetSpeed:    getRandomValidBandwidth(r),
		Encryption:       getRandomEncryption(r),
		Status:           getRandomStatus(r),
		StatusModifiedAt: int64(r.Intn(100)),
	}
}

func GetRandomNodes(r *rand.Rand, accs []simulation.Account) []types.Node {
	var nodes []types.Node
	for i := 0; i < r.Intn(100); i++ {
		nodes = append(nodes, getRandomNode(r, accs))
	}

	return nodes
}

func getRandomSubscription(r *rand.Rand, accs []simulation.Account) types.Subscription {

	return types.Subscription{
		ID:                 getRandomID(r),
		NodeID:             getRandomID(r),
		Client:             simulation.RandomAcc(r, accs).Address,
		PricePerGB:         getRandomValidCoin(r),
		TotalDeposit:       getRandomValidCoin(r),
		RemainingDeposit:   sdk.Coin{"stake",sdk.NewInt(1)},
		RemainingBandwidth: hub.Bandwidth{},
		Status:             getRandomStatus(r),
		StatusModifiedAt:   int64(r.Intn(100)),
	}
}

func GetRandomSubscriptions(r *rand.Rand, accs []simulation.Account) []types.Subscription {
	var subscriptions []types.Subscription
	for i := 0; i < r.Intn(100); i++ {
		subscriptions = append(subscriptions, getRandomSubscription(r, accs))
	}

	return subscriptions
}

func getRandomSession(r *rand.Rand, accs []simulation.Account) types.Session {

	return types.Session{
		ID:               getRandomID(r),
		SubscriptionID:   getRandomID(r),
		Bandwidth:        getRandomValidBandwidth(r),
		Status:           getRandomStatus(r),
		StatusModifiedAt: int64(r.Intn(100)),
	}
}

func GetRandomSessions(r *rand.Rand, accs []simulation.Account) []types.Session {
	var session []types.Session
	for i := 0; i < r.Intn(100); i++ {
		session = append(session, getRandomSession(r, accs))
	}

	return session
}

func getRandomDenom(r *rand.Rand) string { return denom[r.Intn(len(denom))] }

func getRandomIndex(r *rand.Rand) uint64 { return uint64(r.Intn(1000)) }

func getRandomEncryption(r *rand.Rand) string { return simulation.RandStringOfLength(r, 10) }

func getRandomType(r *rand.Rand) string { return simulation.RandStringOfLength(r, 10) }

func getRandomVersion(r *rand.Rand) string { return simulation.RandStringOfLength(r, 10) }

func getRandomMoniker(r *rand.Rand) string { return simulation.RandStringOfLength(r, 10) }

func getRandomStatus(r *rand.Rand) string { return status[r.Intn(len(status))] }
